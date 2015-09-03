package main

import (
	"encoding/json"
	"fmt"
	"github.com/mgutz/minimist"
	"log"
	"net"
	"net/http"
	"os"
)

var port, health int
var domain, proto string
var usage = `
A super simple HTTP redirection server

-h, --help    print this message
-p, --port    port 301 will listen on (default 8080)
-d, --domain  domain requests will redirect to (default 127.0.0.1)
    --proto   protocol HTTP or HTTPS (default HTTP)
    --health  optional port on which to listen for a health check
              Handy for load balancers that will only accept a "200"
              response to keep 301 instance(s) in load.
`

// Will add more stats pending feedback
type HealthState struct {
	Hostname string
	PID      int
}

func handleRedirect(res http.ResponseWriter, req *http.Request) {
	log.Println("Got request for", req.URL.Path)
	http.Redirect(res, req, proto+"://"+domain+req.URL.Path[1:], 301)
}

func handleHealthCheck(res http.ResponseWriter, req *http.Request) {

	// Having to do this makes me think I'm bad at Go
	hostname, _ := os.Hostname()
	ip, _, _ := net.SplitHostPort(req.RemoteAddr)
	log.Println("Health Check From", ip)
	healthstate := HealthState{
		hostname,
		os.Getpid(),
	}
	payload, err := json.Marshal(healthstate)

	res.Header().Set("Server", "301 Redirect Server")

	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	res.Header().Set("Content-Type", "application/json")
	res.Write(payload)

}

func handleCLI() {
	args := minimist.Parse()

	if args.MayBool(false, "help", "h", "?") {
		fmt.Println(usage)
		os.Exit(0)
	}

	port = args.MayInt(8080, "port", "p")
	health = args.AsInt("health")
	domain = args.MayString("127.0.0.1", "domain", "d")
	proto = args.MayString("http", "proto")
}

func redirServer() {

	// Go, WHY U NO INTERPOLATE??
	portstring := fmt.Sprintf(":%d", port)
	fmt.Println("Listening on port", port, "Redirecting to", domain)

	redirServer := http.NewServeMux()
	redirServer.HandleFunc("/", handleRedirect)

	http.ListenAndServe(portstring, redirServer)
}

func healthCheckServer() {
	healthstring := fmt.Sprintf(":%d", health)

	fmt.Println("Listening for health checks on", health)
	healthCheck := http.NewServeMux()
	healthCheck.HandleFunc("/", handleHealthCheck)

	go func() { http.ListenAndServe(healthstring, healthCheck) }()
}

func main() {

	handleCLI()

	if health != 0 {
		healthCheckServer()
	}

	redirServer()

}
