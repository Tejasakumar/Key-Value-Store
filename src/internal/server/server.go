package server

import (
	"KVS/src/internal/QueryInterface"
	"bufio"
	"log"
	"net"
	"strings"
)

// StartServer initializes and runs the TCP server
func StartServer(address string) {
	// Initialize the master engine
	masterEngine := QueryInterface.GetMasterEngine()

	// Start TCP server
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Error starting TCP server: %v", err)
	}
	defer listener.Close()

	log.Printf("Server started, listening on %s", address)

	// Accept client connections
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %v", err)
			continue
		}

		// Handle each client connection in a separate goroutine
		go handleClient(conn, masterEngine)
	}
}

// handleClient processes commands from a client
func handleClient(conn net.Conn, masterEngine *QueryInterface.MasterEngine) {
	defer conn.Close()
	clientAddr := conn.RemoteAddr().String()
	log.Printf("New client connected: %s", clientAddr)

	// Create a new QueryExecutionEngine for this client
	engine := QueryInterface.NewQueryExecutionEngine(masterEngine)

	// Send welcome message
	conn.Write([]byte("Welcome to Key-Value Store Server\n"))
	conn.Write([]byte("Type 'exit' to disconnect\n"))

	// Create a scanner to read input from the client
	scanner := bufio.NewScanner(conn)

	// Process client commands
	for scanner.Scan() {
		userInput := scanner.Text()

		// Check if client wants to exit
		if strings.ToLower(userInput) == "exit" {
			conn.Write([]byte("Goodbye!\n"))
			log.Printf("Client disconnected: %s", clientAddr)
			break
		}

		// Parse and execute the input
		query, err := QueryInterface.ParseQuery(userInput)
		if err != nil {
			conn.Write([]byte(err.Error() + "\n"))
			continue
		}

		// Execute the query
		result := engine.Execute(query)
		conn.Write([]byte(result + "\n"))
	}

	// Check for scanner errors
	if err := scanner.Err(); err != nil {
		log.Printf("Error reading from client %s: %v", clientAddr, err)
	}
}
