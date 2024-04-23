package D1

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

/// Function for assembling local stiffness matrices into a global one
func AssembleKG(analisys Analysis) {
	analisys.CountDistinctNodes()
	analisys.CountElements()
	analisys.SetConstraints()

	globalSize := 1 * analisys.NumNode
	
	globalForces := mat.NewVecDense(globalSize, nil)

	// Iterate through the different structures (Spring1D, Bar1D, Beam1D) and fill the global matrix
	globalMatrix := mat.NewDense(globalSize, globalSize, nil)
	if len(analisys.Spring) > 0 {
		for _, element := range analisys.Spring {
			element.SetKL()
			element.SetKG()
			fillGlobalStiffnessFromElement(globalMatrix, element.Kg, element.N1, element.N2)
		}
	}
	if len(analisys.Bar) > 0 {
		for _, element := range analisys.Bar {
			element.SetL(analisys.Node)
			element.SetKL()
			element.SetKG()
			fillGlobalStiffnessFromElement(globalMatrix, element.Kg, element.N1, element.N2)
		}
	}
	if len(analisys.Beam) > 0 {
		for _, element := range analisys.Beam {
			element.SetL(analisys.Node)
			element.SetKL()
			element.SetKG()
			fillGlobalStiffnessFromElement(globalMatrix, element.Kg, element.N1, element.N2)
		}
	}
	analisys.Kg = *globalMatrix

	// Iterate through the structure and fill the global forces
	for _, force := range analisys.NodalForce {
		globalForces.SetVec(force.Id -1, force.F)
	}
	analisys.Force = *globalForces

	// Remove rows and cols
	mapCostrain := analisys.Constraints
	constrainedIDs := make([]int, 0)
	for id, isConstrained := range mapCostrain {
		if isConstrained {
			constrainedIDs = append(constrainedIDs, id-1)
		}
	}
	globalMatrixReduced := removeRowsAndColsFromMatrix(globalMatrix, constrainedIDs, constrainedIDs)
	globalForcesReduced := removeElementsFromVector(globalForces, constrainedIDs)
	analisys.KgRed = *globalMatrixReduced
	analisys.ForceRed = *globalForcesReduced
	
	// Solve the system Ku = F
	var u mat.VecDense
	err := u.SolveVec(globalMatrixReduced, globalForcesReduced)
	if err != nil {
		fmt.Printf("Error while solving the system: %v\n", err)
		return
	}

	// Fill the displacements vector
	analisys.CalculateDisplacements(u)

	// Fill the reactions vector
	analisys.CalculateReactions()

	PrintSolution(analisys)
}

// Print solution
func PrintSolution(analysis Analysis){
	fmt.Println("Analisys name:")
	fmt.Println(analysis.Name)
	fmt.Println(analysis.Description)
	fmt.Println("")
	fmt.Println("Nodes:")
	fmt.Println(analysis.NumNode)
	fmt.Println("")
	fmt.Println("Elements:")
	fmt.Println(analysis.NumElement)
	fmt.Println("")
	fmt.Println("Constrains:")
	fmt.Println(analysis.Constraints)
	fmt.Println("")
	fmt.Println("Analisys Kg:")
	fmt.Println(mat.Formatted(&analysis.Kg))
	fmt.Println("")
	fmt.Println("Analisys Forces:")
	fmt.Println(VecToString(&analysis.Force))
	fmt.Println("")
	fmt.Println("Displacements:")
	fmt.Println(VecToString(&analysis.Output.Displacement))
	fmt.Println("")
	fmt.Println("Reactions:")
	fmt.Println(VecToString(&analysis.Output.Reaction))
}

/* Utility functions */
/*********************/

// Function to fill the global stiffness matrix from an element's local stiffness matrix
func fillGlobalStiffnessFromElement(globalMatrix *mat.Dense, localMatrix *mat.Dense, nodeIndex1 int, nodeIndex2 int) {
	globalMatrix.Set(nodeIndex1-1, nodeIndex1-1, globalMatrix.At(nodeIndex1-1, nodeIndex1-1)+localMatrix.At(0, 0))
	globalMatrix.Set(nodeIndex1-1, nodeIndex2-1, globalMatrix.At(nodeIndex1-1, nodeIndex2-1)+localMatrix.At(0, 1))
	globalMatrix.Set(nodeIndex2-1, nodeIndex1-1, globalMatrix.At(nodeIndex2-1, nodeIndex1-1)+localMatrix.At(1, 0))
	globalMatrix.Set(nodeIndex2-1, nodeIndex2-1, globalMatrix.At(nodeIndex2-1, nodeIndex2-1)+localMatrix.At(1, 1))
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

