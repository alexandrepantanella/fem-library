package utility

import "math"

func Round(x float64, decimals int) float64 {
	factor := math.Pow(10, float64(decimals))
	return math.Round(x*factor) / factor
}

func RoundMatrix(matrix [][]float64, decimals int) [][]float64 {
	roundedMatrix := make([][]float64, len(matrix))
	for i, row := range matrix {
		roundedRow := make([]float64, len(row))
		for j, val := range row {
			roundedRow[j] = Round(val, decimals)
		}
		roundedMatrix[i] = roundedRow
	}
	return roundedMatrix
}
