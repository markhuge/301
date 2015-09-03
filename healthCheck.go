package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

// Will add more stats pending feedback
type HealthState struct {
	Hostname string
	PID      int
}

func handleHealthCheck(res http.ResponseWriter, req *http.Request) {

	ip, _, _ := net.SplitHostPort(req.RemoteAddr)
	log.Println("Health Check From", ip)

	// Having to do this makes me think I'm bad at Go
	hostname, _ := os.Hostname()

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

func healthCheckServer(port int) {
	healthstring := fmt.Sprintf(":%d", port)

	fmt.Println("Listening for health checks on", port)
	healthCheck := http.NewServeMux()
	healthCheck.HandleFunc("/", handleHealthCheck)

	go func() { http.ListenAndServe(healthstring, healthCheck) }()
}
