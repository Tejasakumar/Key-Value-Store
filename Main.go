package main

import (
	"KVS/QueryInterface"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	masterEngine := QueryInterface.GetMasterEngine()
	scanner := bufio.NewScanner(os.Stdin)
	engine := QueryInterface.NewQueryExecutionEngine(masterEngine)

	fmt.Println("Interactive Input Processor")
	fmt.Println("Type 'exit' to quit")

	for {
		fmt.Print("> ")
		if scanner.Scan() {
			userInput := scanner.Text()

			if strings.ToLower(userInput) == "exit" {
				fmt.Println("Goodbye!")
				break
			}
			query, err := QueryInterface.ParseQuery(userInput)
			if err != nil {
				fmt.Println(err)
				continue
			}
			result := engine.Execute(query)
			fmt.Println(result)
		} else {
			if err := scanner.Err(); err != nil {
				fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
			}
			break
		}
	}
}
