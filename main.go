package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/fem-library/analisys"
	"github.com/fem-library/solver"
	"gonum.org/v1/gonum/mat"
)

func main() {
	// Controllo se è stato fornito il nome del file come argomento
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <filename.json>")
		return
	}

	// Leggo il nome del file JSON dall'argomento della riga di comando
	fileName := os.Args[1]

	// Leggere il contenuto del file JSON
	jsonData, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Errore nella lettura del file JSON:", err)
		return
	}

	// Decodifica il JSON nella struttura Analisys1D
	var analisys analisys.Analisys
	err = json.Unmarshal(jsonData, &analisys)
	if err != nil {
		fmt.Println("Errore nella decodifica JSON:", err)
		return
	}
	

	switch 
	{
	case analisys.Dim == 1: 
		Run1D(&analisys)
	case analisys.Dim == 2: 
		Run2D(&analisys)
	case analisys.Dim == 3: 
		Run3D(&analisys)
	default: 
		fmt.Println(" Tipo analisi non definita")
	}
}

func Run1D(analisys *analisys.Analisys){
	if analisys == nil {
        fmt.Println("Errore: il puntatore analisys è nil")
        return
    }
	solver.AssembleGlobalStiffnessMatrix1D(analisys)
	fmt.Println("Matrice Globale:")
	fmt.Println(mat.Formatted(&analisys.GlobalStiffnessMatrix))
}

func Run2D(analisys *analisys.Analisys){
}

func Run3D(analisys *analisys.Analisys){
}