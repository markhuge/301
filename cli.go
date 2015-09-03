package main

import (
	"fmt"
	"github.com/mgutz/minimist"
	"os"
)

var port, health int
var domain, proto string
var version = "v2.0.0"
var usage = `
A super simple HTTP redirection server

-h, --help    print this message
-v, --version print 301 version
-p, --port    port 301 will listen on (default 8080)
-d, --domain  domain requests will redirect to (default 127.0.0.1)
    --proto   protocol HTTP or HTTPS (default HTTP)
    --health  optional port on which to listen for a health check
              Handy for load balancers that will only accept a "200"
              response to keep 301 instance(s) in load.
`

func handleCLI() (proto, domain string, port, health int) {
	args := minimist.Parse()

	if args.MayBool(false, "help", "h", "?") {
		fmt.Println(usage)
		os.Exit(0)
	}

	if args.MayBool(false, "v", "version") {
		fmt.Println(version)
		os.Exit(0)
	}

	port = args.MayInt(8080, "port", "p")
	health = args.AsInt("health")
	domain = args.MayString("127.0.0.1", "domain", "d")
	proto = args.MayString("http", "proto")

	return proto, domain, port, health
}
