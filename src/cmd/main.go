package main

import (
	"KVS/src/internal/server"
	"flag"
	"fmt"
)

func main() {
	// Parse command line flags
	var (
		port = flag.String("port", "8080", "Port to listen on")
		host = flag.String("host", "0.0.0.0", "Host address to bind to")
	)
	flag.Parse()

	// Format the address string
	address := fmt.Sprintf("%s:%s", *host, *port)
	
	fmt.Printf("Starting Key-Value Store TCP server on %s\n", address)
	
	// Start the server
	server.StartServer(address)
}