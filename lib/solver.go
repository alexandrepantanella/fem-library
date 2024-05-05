package lib

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

func Solve(a Analysis){
	//Initial check
	err := CheckStructure(a)
	if err != nil {
		fmt.Println("Error on checking json:", err)
		return
	}

	var globalSize int
	var globalMatrix *mat.Dense
	
	//Count elements and nodes
	a.CountDistinctNodes()
	a.CountElements()

	//Set sizeof global matrix
	globalSize = a.InputData.DoF * a.CalcData.NumNode
	globalMatrix = mat.NewDense(globalSize, globalSize, nil)

	//Init CalcData
	calcData := &a.CalcData
	calcData.Init()
	
	//Set constrains
	a.SetConstraints()

	//Fill Local and Global Stiffness for each elements
	if len(a.InputData.Element1D) > 0 {
		for i := range a.InputData.Element1D {
			e := &a.InputData.Element1D[i]
			e.Init(a.InputData.DoF)
			e.SetLength(a.InputData.Node, a.InputData.DoF)
			e.SetTheta(a.InputData.Node, a.InputData.DoF)
			e.SetKL(a.InputData.DoF)
			e.SetKG(a.InputData.DoF)
			//Save in Calculated data structure
			a.CalcData.LocalStiffness[e.Id] = e.KL
			a.CalcData.GlobalStiffness[e.Id] = e.KG
			// Add local stiffness to global matrix
			assembleGlobalStiffnessMatrix(globalMatrix, e, a.InputData.DoF)
			a.CalcData.Global = globalMatrix
		}
	}

	PrintSolution(a)
	
}

func assembleGlobalStiffnessMatrix(globalMatrix *mat.Dense, e *Element1D, DoF int) {
	if e.N1-1 >= 0 && e.N2-1 >= 0 {
		for i := 0; i < DoF; i++ {
			for j := 0; j < DoF; j++ {
				globalMatrix.Set(e.N1-1+i, e.N1-1+j, globalMatrix.At(e.N1-1+i, e.N1-1+j)+e.KG.At(i, j))
				globalMatrix.Set(e.N1-1+i, e.N2-1+j, globalMatrix.At(e.N1-1+i, e.N2-1+j)+e.KG.At(i, j+DoF))
				globalMatrix.Set(e.N2-1+i, e.N1-1+j, globalMatrix.At(e.N2-1+i, e.N1-1+j)+e.KG.At(i+DoF, j))
				globalMatrix.Set(e.N2-1+i, e.N2-1+j, globalMatrix.At(e.N2-1+i, e.N2-1+j)+e.KG.At(i+DoF, j+DoF))
			}
		}
	}
}

// Function to remove rows and columns from a matrix
func removeRowsAndColsFromMatrix(matrix *mat.Dense, rowIndices []int, colIndices []int) *mat.Dense {
	rows, cols := matrix.Dims()
	rowsToRemove := len(rowIndices)
	colsToRemove := len(colIndices)

	// Create a new data slice excluding the specified rows and columns
	newData := make([]float64, 0, (rows-rowsToRemove)*(cols-colsToRemove))
	for i := 0; i < rows; i++ {
		if !containsIndex(rowIndices, i) {
			for j := 0; j < cols; j++ {
				if !containsIndex(colIndices, j) {
					newData = append(newData, matrix.At(i, j))
				}
			}
		}
	}

	// Create a new matrix with the excluded data
	newRows := rows - rowsToRemove
	newCols := cols - colsToRemove
	newMatrix := mat.NewDense(newRows, newCols, newData)

	return newMatrix
}

// Function to remove elements from a vector
func removeElementsFromVector(vec *mat.VecDense, indices []int) *mat.VecDense {
	size := vec.Len()
	indicesMap := make(map[int]bool)

	// Create a map of indices to remove for faster access
	for _, index := range indices {
		indicesMap[index] = true
	}

	// Create a new data slice excluding the specified indices
	newData := make([]float64, 0, size)
	for i := 0; i < size; i++ {
		if !indicesMap[i] {
			newData = append(newData, vec.AtVec(i))
		}
	}

	// Create a new mat.VecDense with the excluded data
	newVec := mat.NewVecDense(len(newData), newData)

	return newVec
}

// Function to check if an index exists in a slice of indices
func containsIndex(indices []int, index int) bool {
	for _, idx := range indices {
		if idx == index {
			return true
		}
	}
	return false
}


// Print solution
func PrintSolution(a Analysis){
	fmt.Println("Analisys name:")
	fmt.Println(a.InputData.Name)
	fmt.Println(a.InputData.Description)
	fmt.Println(a.InputData.Reference)
	fmt.Println("")
	fmt.Println("Nodes:")
	fmt.Println(a.CalcData.NumNode)
	fmt.Println("")
	fmt.Println("Number of elements:")
	fmt.Println(a.CalcData.NumElement)
	fmt.Println("")
	fmt.Println("Map Constrains:")
	fmt.Println(a.CalcData.Constraints)
	fmt.Println("")
	fmt.Println("Elements Local Stiffness Matrix:")
	for id, e := range a.CalcData.LocalStiffness {
		fmt.Println("Element:", id)
		fmt.Println(mat.Formatted(e))
	}
	fmt.Println("")
	fmt.Println("Elements Global Stiffness Matrix:")
	for id, e := range a.CalcData.GlobalStiffness {
		fmt.Println("Element:", id)
		fmt.Println(mat.Formatted(e))
	}
	fmt.Println("")
	fmt.Println("Analisys Global Stiffness Matrix:")
	fmt.Println(mat.Formatted(a.CalcData.Global))
	// fmt.Println("")
	// fmt.Println("Analisys Forces:")
	// fmt.Println(VecToString(&analysis.Force))
	// fmt.Println("")
	// fmt.Println("Displacements:")
	// fmt.Println(VecToString(&analysis.Output.Displacement))
	// fmt.Println("")
	// fmt.Println("Reactions:")
	// fmt.Println(VecToString(&analysis.Output.Reaction))
}