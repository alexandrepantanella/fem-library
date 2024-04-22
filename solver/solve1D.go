package solver

import (
	"github.com/fem-library/analisys"
	"gonum.org/v1/gonum/mat"
)

// Funzione per l'assemblaggio delle matrici di rigidezza locali in una globale
func AssembleGlobalStiffnessMatrix1D(analisys analisys.Analisys1D) {
	globalSize := 2 * countDistinctNodes(analisys) // La dimensione della matrice globale sarà il doppio del numero di nodi
	globalMatrix := mat.NewDense(globalSize, globalSize, nil)

	// Iterazione su tutte le matrici di rigidezza locali e l'assemblaggio nella matrice globale
	for i, localMatrix := range []analisys.Beam1D {
		nodeIndex := i * 2
		// Aggiungi la matrice di rigidezza locale alla posizione corretta nella matrice globale
		for j := 0; j < 2; j++ {
			for k := 0; k < 2; k++ {
				globalMatrix.Set(nodeIndex+j, nodeIndex+k, globalMatrix.At(nodeIndex+j, nodeIndex+k)+localMatrix.K.At(j, k))
			}
		}
	}
	for i, localMatrix = range []analisys.Bar1D {
		nodeIndex = i * 2
		// Aggiungi la matrice di rigidezza locale alla posizione corretta nella matrice globale
		for j := 0; j < 2; j++ {
			for k := 0; k < 2; k++ {
				globalMatrix.Set(nodeIndex+j, nodeIndex+k, globalMatrix.At(nodeIndex+j, nodeIndex+k)+localMatrix.K.At(j, k))
			}
		}
	}
	for i, localMatrix = range []analisys.Spring1D {
		nodeIndex = i * 2
		// Aggiungi la matrice di rigidezza locale alla posizione corretta nella matrice globale
		for j := 0; j < 2; j++ {
			for k := 0; k < 2; k++ {
				globalMatrix.Set(nodeIndex+j, nodeIndex+k, globalMatrix.At(nodeIndex+j, nodeIndex+k)+localMatrix.K.At(j, k))
			}
		}
	}


	analisys.GlobalStiffnessMatrix := globalMatrix
}

func countDistinctNodes(analisys *analysis.Analisys1D) int {
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