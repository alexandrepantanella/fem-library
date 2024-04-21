package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/fem-library/analisys"
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
	var analisys analisys.Analisys1D
	err = json.Unmarshal(jsonData, &analisys)
	if err != nil {
		fmt.Println("Errore nella decodifica JSON:", err)
		return
	}

	// Stampa la struttura popolata
	fmt.Printf("%+v\n", analisys)

	// Itera su ogni molla e stampa il risultato del metodo StiffnessMatrix
	for _, spring := range analisys.Spring1D {
		fmt.Printf("Spring ID: %d\n", spring.ID)
		fmt.Println(spring.StiffnessMatrix())
	}
}
