package main

import (
	"fmt"
	"log"
	"net/http"
)

func handleRedirect(res http.ResponseWriter, req *http.Request, proto, domain string, port int) {
	log.Println("Got request for", req.URL.Path)
	http.Redirect(res, req, proto+"://"+domain+"/"+req.URL.Path[1:], 301)
}

func redirServer(proto, domain string, port int) {

	// Go, WHY U NO INTERPOLATE??
	portstring := fmt.Sprintf(":%d", port)
	fmt.Println("Listening on port", port, "Redirecting to", domain)

	redirServer := http.NewServeMux()

	// This feels super dirty
	redirServer.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		handleRedirect(res, req, proto, domain, port)
	})

	http.ListenAndServe(portstring, redirServer)
}
