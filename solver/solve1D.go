package solver

import (
	"fmt"

	"github.com/fem-library/analisys"
	"github.com/fem-library/force"
	"gonum.org/v1/gonum/mat"
)

// Funzione per l'assemblaggio delle matrici di rigidezza locali in una globale
func AssembleGlobalStiffnessMatrix1D(analisys *analisys.Analisys) {
    globalSize := 1 * countDistinctNodes(analisys.Analisys1D) // Assicurati che countDistinctNodes restituisca il numero corretto di nodi distinti
	globalMatrix := mat.NewDense(globalSize, globalSize, nil) // Inizializza correttamente la matrice di rigidezza globale
	globalForces := mat.NewVecDense(globalSize, nil)
	analisys.Force = mat.NewVecDense(globalSize, nil)

    // Itera attraverso le diverse strutture (Bar1D, Spring1D, Beam1D) e riempi la matrice globale
	if len(analisys.Analisys1D.Bar1D) > 0 {
		for _, element := range analisys.Analisys1D.Bar1D {
			element.StiffnessMatrix()
			element.GlobalStiffnessMatrix()
			fillGlobalStiffnessFromElement(element.KGlobal, globalMatrix, element.Node1.ID, element.Node2.ID)
		}
	}
	if len(analisys.Analisys1D.Spring1D) > 0 {
		for _, element := range analisys.Analisys1D.Spring1D {
			element.StiffnessMatrix()
			element.GlobalStiffnessMatrix()
			fillGlobalStiffnessFromElement(element.KGlobal, globalMatrix, element.Node1.ID, element.Node2.ID)
			fillGlobalForcesFromElement( &element.Force, globalForces, element.Node1.ID, element.Node2.ID)
			analisys.Force = globalForces
		}
	}
	if len(analisys.Analisys1D.Beam1D) > 0 {
		for _, element := range analisys.Analisys1D.Beam1D {
			element.StiffnessMatrix()
			element.GlobalStiffnessMatrix()
			fillGlobalStiffnessFromElement(element.KGlobal, globalMatrix, element.Node1.ID, element.Node2.ID)
		}
	}
    
	// Assegna direttamente globalMatrix ad analisys.GlobalStiffnessMatrix senza dereferenziare
    analisys.GlobalStiffnessMatrix = *globalMatrix
	
	// Risolvi il sistema Ku = F
	var u mat.VecDense
	err := u.SolveVec(globalMatrix, globalForces)
	if err != nil {
		fmt.Printf("Errore durante la risoluzione del sistema: %v\n", err)
		return
	}
}

func fillGlobalStiffnessFromElement(localMatrix *mat.Dense, globalMatrix *mat.Dense, nodeIndex1 int, nodeIndex2 int) {
	globalMatrix.Set(nodeIndex1-1, nodeIndex1-1, globalMatrix.At(nodeIndex1-1, nodeIndex1-1)+localMatrix.At(0, 0))
	globalMatrix.Set(nodeIndex1-1, nodeIndex2-1, globalMatrix.At(nodeIndex1-1, nodeIndex2-1)+localMatrix.At(0, 1))
	globalMatrix.Set(nodeIndex2-1, nodeIndex1-1, globalMatrix.At(nodeIndex2-1, nodeIndex1)+localMatrix.At(1, 0))
	globalMatrix.Set(nodeIndex2-1, nodeIndex2-1, globalMatrix.At(nodeIndex2-1, nodeIndex2-1)+localMatrix.At(1, 1))
}

func fillGlobalForcesFromElement(elementForces *force.Force1D, globalForces *mat.VecDense,  nodeIndex1 int, nodeIndex2 int){
	fmt.Println("Nodo1:", nodeIndex1, "Nodo2:", nodeIndex2)
	globalForces.SetVec(nodeIndex1-1,globalForces.At(nodeIndex1-1,0)+elementForces.X1)
	globalForces.SetVec(nodeIndex2-1,globalForces.At(nodeIndex2-1,0)+elementForces.X2)
	fmt.Println(mat.Formatted(globalForces))
}

func countDistinctNodes(analisys analisys.Analisys1D) int {
	nodeMap := make(map[int]bool)
	springData := analisys.Spring1D
	beamData:= analisys.Beam1D
	barData:= analisys.Bar1D
	for _, spring := range springData {
		nodeMap[spring.Node1.ID] = true
		nodeMap[spring.Node2.ID] = true
	}
	for _, beam := range beamData {
		nodeMap[beam.Node1.ID] = true
		nodeMap[beam.Node2.ID] = true
	}
	for _, bar := range barData {
		nodeMap[bar.Node1.ID] = true
		nodeMap[bar.Node2.ID] = true
	}
	return len(nodeMap)
}
