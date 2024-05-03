package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Checking if the file name has been provided as an argument.
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run <analisysDim> main.go <filename.json>")
		return
	}

	// Reading dimension analisys
	dimStr := os.Args[1]
	dim, err := strconv.Atoi(dimStr)
	if err != nil || dim >3 {
		fmt.Println("Invalid dimension:", dimStr)
		return
	}

	// Reading the name of the JSON file from the command-line argument.
	fileName := os.Args[2]

	// Reading the content of the JSON file.
	jsonData, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error while reading the JSON file:", err)
		return
	}

	
	switch 
	{
	case dim == 1: 
		Run1D(jsonData)
	// case dim == 2: 
	// 	Run2D(jsonData)
	// case dim == 3: 
	// 	Run3D(jsonData)
	default: 
		fmt.Println("Analisys not defined")
	}
}

func Run1D(jsonData []byte){
	
	analisys := Analysis{}

	//  Decode the JSON into the Analisys structure
	err := json.Unmarshal(jsonData, &analisys)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
	if analisys.InputData.DoF == 0 {
        fmt.Println("Error: JSON decoding was unsuccessful")
        return
    }
	checkStructure(analisys)
	Solve(analisys)

	// fmt.Println(mat.Formatted(&analisys.GlobalStiffnessMatrix))
}

// func Run2D(analisys *D1.Analisys){
// }

// func Run3D(analisys *D1.Analisys){
// }