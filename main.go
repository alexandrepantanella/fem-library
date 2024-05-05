package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/fem-library/lib"
)

func main() {
	// Checking if the file name has been provided as an argument.
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <filename.json>")
		return
	}

	// Reading the name of the JSON file from the command-line argument.
	fileName := os.Args[1]

	// Reading the content of the JSON file.
	jsonData, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error while reading the JSON file:", err)
		return
	}

	a := lib.Analysis{}

	//  Decode the JSON into the Analisys structure
	err = json.Unmarshal(jsonData, &a)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
	if a.InputData.DoF == 0 {
		fmt.Println("Error: JSON decoding was unsuccessful")
		return
		}
	lib.Solve(a)

	
}
