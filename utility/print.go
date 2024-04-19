package utility

import (
	"fmt"
	"net/http"
)

// PrintMatrix stampa la matrice in forma matriciale in console
func PrintMatrix(matrix [][]float64) {
	for _, row := range matrix {
		for _, element := range row {
			fmt.Printf("%.2f ", element)
		}
		fmt.Println()
	}
}

// PrintMatrixHttp stampa la matrice in forma matriciale in http
func PrintMatrixHttp(w http.ResponseWriter, matrix [][]float64) {
	for _, row := range matrix {
		for _, element := range row {
			fmt.Fprintf(w, "%.2f ", element)
		}
		fmt.Fprintln(w)
	}
}
