package main

func main() {
	proto, domain, port, health := handleCLI()

	if health != 0 {
		healthCheckServer(health)
	}

	redirServer(proto, domain, port)
}
