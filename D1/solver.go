package D1

import (
	"gonum.org/v1/gonum/mat"
)

//Init all elements
func InitElements(analisys Analysis) {
	for _, element := range analisys.Spring {
		element.SetKL()
		element.SetKG()	
	}
	for _, element := range analisys.Bar {
		element.SetKL()
		element.SetKG()	
	}
	for _, element := range analisys.Beam {
		element.SetKL()
		element.SetKG()	
	}
}


// Funzione per l'assemblaggio delle matrici di rigidezza locali in una globale
func AssembleKG(analisys Analisys) {
    globalSize := 1 * countDistinctNodes(analisys)
	globalMatrix := mat.NewDense(globalSize, globalSize, nil)
	globalForces := mat.NewVecDense(globalSize, nil)
	analisys.Force = mat.NewVecDense(globalSize, nil)

    // Itera attraverso le diverse strutture (Bar1D, Spring1D, Beam1D) e riempi la matrice globale
	if len(analisys.Bar) > 0 {
		for _, element := range analisys.Bar {
			fillGlobalStiffnessFromElement(element.Kg, globalMatrix, element.N1, element.N2)
		}
	}
	if len(analisys.Analisys1D.Spring) > 0 {
		for _, element := range analisys.Spring {
			fillGlobalStiffnessFromElement(element.Kg, globalMatrix, element.N1, element.N2)
			//fillGlobalForcesFromElement( &element.Force, globalForces, element.Node1.ID, element.Node2.ID)
			analisys.Force = globalForces
		}
	}
	if len(analisys.Analisys1D.Beam1D) > 0 {
		for _, element := range analisys.Analisys1D.Beam1D {
			fillGlobalStiffnessFromElement(element.Kg, globalMatrix, element.N1, element.N2)
		}
	}
    
	// Assegna direttamente globalMatrix ad analisys.GlobalStiffnessMatrix senza dereferenziare
    analisys.GlobalStiffnessMatrix = *globalMatrix
	
	// Risolvi il sistema Ku = F
	// var u mat.VecDense
	// err := u.SolveVec(globalMatrix, globalForces)
	// if err != nil {
	// 	fmt.Printf("Errore durante la risoluzione del sistema: %v\n", err)
	// 	return
	// }
}

func fillGlobalStiffnessFromElement(globalMatrix *mat.Dense, localMatrix *mat.Dense, nodeIndex1 int, nodeIndex2 int) {
	globalMatrix.Set(nodeIndex1-1, nodeIndex1-1, globalMatrix.At(nodeIndex1-1, nodeIndex1-1)+localMatrix.At(0, 0))
	globalMatrix.Set(nodeIndex1-1, nodeIndex2-1, globalMatrix.At(nodeIndex1-1, nodeIndex2-1)+localMatrix.At(0, 1))
	globalMatrix.Set(nodeIndex2-1, nodeIndex1-1, globalMatrix.At(nodeIndex2-1, nodeIndex1)+localMatrix.At(1, 0))
	globalMatrix.Set(nodeIndex2-1, nodeIndex2-1, globalMatrix.At(nodeIndex2-1, nodeIndex2-1)+localMatrix.At(1, 1))
}

// func fillGlobalForcesFromElement(elementForces *force.Force1D, globalForces *mat.VecDense,  nodeIndex1 int, nodeIndex2 int){
// 	fmt.Println("Nodo1:", nodeIndex1, "Nodo2:", nodeIndex2)
// 	globalForces.SetVec(nodeIndex1-1,globalForces.At(nodeIndex1-1,0)+elementForces.X1)
// 	globalForces.SetVec(nodeIndex2-1,globalForces.At(nodeIndex2-1,0)+elementForces.X2)
// 	fmt.Println(mat.Formatted(globalForces))
// }

func countDistinctNodes(analisys Analisys) int {
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
