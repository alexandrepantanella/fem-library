package D1

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

// Function for assembling local stiffness matrices into a global one
func AssembleKG(analisys Analysis) {
	globalSize := 1 * countDistinctNodes(analisys)
	globalMatrix := mat.NewDense(globalSize, globalSize, nil)
	globalForces := mat.NewVecDense(globalSize, nil)

    // Iterate through the different structures (Spring1D, Bar1D, Beam1D) and fill the global matrix
	if len(analisys.Spring) > 0 {
		for _, element := range analisys.Spring {
			element.SetKL()
			element.SetKG()	
			fmt.Println("Spring Kg:")
			fmt.Println(mat.Formatted(element.Kg))
			fillGlobalStiffnessFromElement(globalMatrix, element.Kg, element.N1, element.N2)
			fillGlobalForcesFromElement( globalForces, element.F1, element.F2, element.N1, element.N2)
		}
	}
	if len(analisys.Bar) > 0 {
		for _, element := range analisys.Bar {
			element.SetKL()
			element.SetKG()	
			fmt.Println("Bar Kg:")
			fmt.Println(mat.Formatted(element.Kg))
			fillGlobalStiffnessFromElement(globalMatrix, element.Kg, element.N1, element.N2)
			fillGlobalForcesFromElement( globalForces, element.F1, element.F2, element.N1, element.N2)
		}
	}
	if len(analisys.Beam) > 0 {
		for _, element := range analisys.Beam {
			element.SetKL()
			element.SetKG()	
			fmt.Println("Beam Kg:")
			fmt.Println(mat.Formatted(element.Kg))
			fillGlobalStiffnessFromElement(globalMatrix, element.Kg, element.N1, element.N2)
			fillGlobalForcesFromElement( globalForces, element.F1, element.F2, element.N1, element.N2)
		}
	}
    
    analisys.Kg = *globalMatrix
	analisys.Force = *globalForces

	fmt.Println("Analisys Kg:")
	fmt.Println(mat.Formatted(globalMatrix))
	fmt.Println("Analisys Forces:")
	fmt.Println(mat.Formatted(globalForces))

	//Reducing the global stiffness matrix with constrains
	reducedMatrix, reducedForces, _ := AssembleReducedKG(analisys)

	fmt.Println("Reduced analisys Kg:")
	fmt.Println(mat.Formatted(reducedMatrix))
	fmt.Println("Reduced analisys Forces:")
	fmt.Println(mat.Formatted(reducedForces))
	
	// Solve the system Ku = F
	var u mat.VecDense
	err := u.SolveVec(reducedMatrix, reducedForces)
	if err != nil {
		fmt.Printf("Error while solving the system: %v\n", err)
		return
	}
}

func fillGlobalStiffnessFromElement(globalMatrix *mat.Dense, localMatrix *mat.Dense, nodeIndex1 int, nodeIndex2 int) {
	fmt.Println("Node1:", nodeIndex1, "Node2:", nodeIndex2)
	globalMatrix.Set(nodeIndex1-1, nodeIndex1-1, globalMatrix.At(nodeIndex1-1, nodeIndex1-1)+localMatrix.At(0, 0))
	globalMatrix.Set(nodeIndex1-1, nodeIndex2-1, globalMatrix.At(nodeIndex1-1, nodeIndex2-1)+localMatrix.At(0, 1))
	globalMatrix.Set(nodeIndex2-1, nodeIndex1-1, globalMatrix.At(nodeIndex2-1, nodeIndex1)+localMatrix.At(1, 0))
	globalMatrix.Set(nodeIndex2-1, nodeIndex2-1, globalMatrix.At(nodeIndex2-1, nodeIndex2-1)+localMatrix.At(1, 1))
}

func fillGlobalForcesFromElement( globalForces *mat.VecDense, F1 float64, F2 float64,  nodeIndex1 int, nodeIndex2 int){
	fmt.Println("Node1:", nodeIndex1, "Node2:", nodeIndex2)
	globalForces.SetVec(nodeIndex1-1,globalForces.At(nodeIndex1-1,0)+F1)
	globalForces.SetVec(nodeIndex2-1,globalForces.At(nodeIndex2-1,0)+F2)
	fmt.Println(mat.Formatted(globalForces))
}

func AssembleReducedKG(analisys Analysis) (reducedMatrix *mat.Dense, reducedForces *mat.VecDense, err error) {
	// Ottieni i vincoli dalla struttura
	constraints := getConstraints(analisys)
	fmt.Println(constraints)

	// Calcola le dimensioni della matrice di rigidità globale ridotta
	reducedSize := countConstrainedNodes(constraints)
	fmt.Println(reducedSize)

	// Inizializza la matrice di rigidità globale ridotta e il vettore delle forze ridotto
	reducedMatrix = mat.NewDense(reducedSize, reducedSize, nil)
	reducedForces = mat.NewVecDense(reducedSize, nil)

	// Itera attraverso gli elementi strutturali e riempi la matrice di rigidità globale ridotta e il vettore delle forze ridotto
	for _, element := range analisys.Spring {
		if !isConstrainedNode(constraints, element.N1) && !isConstrainedNode(constraints, element.N2) {
			// Se entrambi i nodi non sono vincolati, aggiungi le righe e le colonne corrispondenti alla matrice ridotta
			fillGlobalStiffnessFromElement(reducedMatrix, element.Kg, element.N1, element.N2)
			fillGlobalForcesFromElement(reducedForces, element.F1, element.F2, element.N1, element.N2)
		}
	}

	return reducedMatrix, reducedForces, nil
}

func countDistinctNodes(analisys Analysis) int {
	nodeMap := make(map[int]bool)
	springData := analisys.Spring
	beamData:= analisys.Beam
	barData:= analisys.Bar
	for _, spring := range springData {
		nodeMap[spring.N1] = true
		nodeMap[spring.N2] = true
	}
	for _, beam := range beamData {
		nodeMap[beam.N1] = true
		nodeMap[beam.N2] = true
	}
	for _, bar := range barData {
		nodeMap[bar.N1] = true
		nodeMap[bar.N2] = true
	}
	return len(nodeMap)
}

func getConstraints(analisys Analysis) []int {
	constraints := make([]int, 0)
	for _, node := range analisys.Node {
		if !node.C  {
			constraints = append(constraints, node.Id)
		}
	}
	return constraints
}

func countConstrainedNodes(constraints []int) int {
	return len(constraints)
}

func isConstrainedNode(constraints []int, nodeID int) bool {
	for _, c := range constraints {
		if c == nodeID {
			return true
		}
	}
	return false
}