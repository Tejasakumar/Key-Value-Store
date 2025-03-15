package main

import (
	"KVS/Storage"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
var currentDb string = ""
func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Interactive Input Processor")
	fmt.Println("Type 'exit' to quit")
	dbs := make(map[string]*Storage.Db)

	for {
		fmt.Print("> ")
		if scanner.Scan() {
			userInput := scanner.Text()

			// Check if user wants to exit
			if strings.ToLower(userInput) == "exit" {
				fmt.Println("Goodbye!")
				break
			}

			// Process the input
			processInput(userInput, dbs)
		} else {
			// Handle potential scanner errors
			if err := scanner.Err(); err != nil {
				fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
			}
			break
		}
	}
}

// processInput handles the user input
func processInput(input string, dbs map[string]*Storage.Db) {
	// This is where you can implement your logic to process the input

	input = strings.TrimSpace(input)
	input = strings.ToLower(input)

	
	if strings.HasPrefix(input, "use") {
		parts := strings.Split(input, " ")
		if len(parts) != 2 {
			fmt.Println("Invalid input. Usage: use <db>")
			return
		}

		currentDb = parts[1]
		_, ok := dbs[currentDb]
		if !ok {
			fmt.Println("Database not found creating database" + currentDb)
			dbs[currentDb] = Storage.GetDb(currentDb)
		}
		
	}else if strings.HasPrefix(input, "showdbs") {

		for key := range dbs {
			fmt.Println(key)
		}

	}else if strings.HasPrefix(input, "curdb") {

		if currentDb != "" {
			fmt.Println(currentDb)
		}else{
			fmt.Println("No database selected use database using 'use <db>'")
		}
		return

	}
	if len(dbs) == 0 {

		fmt.Println("No databases found first create a database using 'use <db>'")
		return

	}

	db := dbs[currentDb]
	if strings.HasPrefix(input, "keys") {
		db.Keys()
		return
	}else if strings.HasPrefix(input, "dropdb") {

		parts := strings.Split(input, " ")
		if len(parts) != 2 {
			fmt.Println("Invalid input. Usage: use <db>")
			return
		}
		dbtmp := dbs[parts[1]]
		dbtmp.DropDb()
		delete(dbs, parts[1])
		if currentDb == parts[1] {
			currentDb = ""
		}

	}else if strings.HasPrefix(input, "setttl") {

		parts := strings.Split(input, " ")
		if len(parts) != 3 {
			fmt.Println("Invalid input. Usage: setttl <key> <ttl>")
			return
		}

		key := parts[1]
		ttl, err := strconv.Atoi(parts[2])
		if err != nil {
			fmt.Println("Invalid TTL. Usage: setttl <key> <ttl>")
			return
		}

		err = db.Setttl(key, ttl)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Printf(" TTL for '%s': %d seconds\n", key, ttl)

	}else if strings.HasPrefix(input, "getttl") {
		parts := strings.Split(input, " ")
		if len(parts) != 2 {
			fmt.Println("Invalid input. Usage: getttl <key>")
			return
		}

		key := parts[1]
		ttl, err := db.Getttl(key)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Printf("\tTTL for '%s': %d seconds\n", key, ttl)
	}

	if strings.HasPrefix(input, "put") {
		parts := strings.Split(input, " ")
		if len(parts) != 3 {
			fmt.Println("Invalid input. Usage: put <key> <value>")
			return
		}

		key := parts[1]
		value := parts[2]
		db.Put(key, value)	

	}else if strings.HasPrefix(input, "get") {
		parts := strings.Split(input, " ")
		if len(parts) != 2 {
			fmt.Println("Invalid input. Usage: get <key>")
			return
		}

		key := parts[1]
		value, err := db.Get(key)
		if err != nil {
			fmt.Println("Error:", err)	
		}else{
			fmt.Printf(" Value for '%s': %v\n", key, value.Data)
		}

	}else if strings.HasPrefix(input, "delete") {
		parts := strings.Split(input, " ")
		if len(parts) != 2 {
			fmt.Println("Invalid input. Usage: delete <key>")
			return
		}

		key := parts[1]
		err := db.Delete(key)
		if err != nil {
			fmt.Println("Error:", err)
			return	
		}
		fmt.Printf(" Deleted key: %s\n", key)
	}else if strings.HasPrefix(input, "list") {

		db.List()

	}else if strings.HasPrefix(input, "rmttl") {
		parts := strings.Split(input, " ")
		if len(parts) != 2 {
			fmt.Println("Invalid input. Usage: rmttl <key>")
			return
		}
		key := parts[1]
		db.RemoveTTL(key)
	}else if strings.HasPrefix(input, "upttl") {
		parts := strings.Split(input, " ")
		if len(parts) != 3 {
			fmt.Println("Invalid input. Usage: upttl <key> <ttl>")
			return
		}
		key := parts[1]
		ttl, err := strconv.Atoi(parts[2])
		if err != nil {
			fmt.Println("Invalid TTL. Usage: upttl <key> <ttl>")
			return
		}
		db.Updatettldb(key, ttl)
	}
}