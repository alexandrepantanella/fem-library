package solver

import (
	"fmt"

	"github.com/fem-library/analisys"
	"gonum.org/v1/gonum/mat"
)

// Funzione per l'assemblaggio delle matrici di rigidezza locali in una globale
func AssembleGlobalStiffnessMatrix1D(analisys *analisys.Analisys) {
    globalSize := 1 * countDistinctNodes(analisys.Analisys1D) // Assicurati che countDistinctNodes restituisca il numero corretto di nodi distinti
	globalMatrix := mat.NewDense(globalSize, globalSize, nil) // Inizializza correttamente la matrice globale
    // Itera attraverso le diverse strutture (Bar1D, Spring1D, Beam1D) e riempi la matrice globale
	if len(analisys.Analisys1D.Bar1D) > 0 {
		for _, localMatrix := range analisys.Analisys1D.Bar1D {
			localMatrix.StiffnessMatrix()
			localMatrix.GlobalStiffnessMatrix()
			fillGlobalFromElement(localMatrix.KGlobal, globalMatrix, localMatrix.Node1.ID, localMatrix.Node2.ID)
		}
	}
	if len(analisys.Analisys1D.Spring1D) > 0 {
		for _, localMatrix := range analisys.Analisys1D.Spring1D {
			localMatrix.StiffnessMatrix()
			localMatrix.GlobalStiffnessMatrix()
			fmt.Println("Matrice Locale:")
			fmt.Println(mat.Formatted(localMatrix.KGlobal))
			fillGlobalFromElement(localMatrix.KGlobal, globalMatrix, localMatrix.Node1.ID, localMatrix.Node2.ID)
			fmt.Println(mat.Formatted(globalMatrix))
		}
	}
	if len(analisys.Analisys1D.Beam1D) > 0 {
		for _, localMatrix := range analisys.Analisys1D.Beam1D {
			localMatrix.StiffnessMatrix()
			localMatrix.GlobalStiffnessMatrix()
			fillGlobalFromElement(localMatrix.KGlobal, globalMatrix, localMatrix.Node1.ID, localMatrix.Node2.ID)
		}
	}
    // Assegna direttamente globalMatrix ad analisys.GlobalStiffnessMatrix senza dereferenziare
    analisys.GlobalStiffnessMatrix = *globalMatrix
}

func fillGlobalFromElement(localMatrix *mat.Dense, globalMatrix *mat.Dense, nodeIndex1 int, nodeIndex2 int) {
	fmt.Println("Nodo1:", nodeIndex1)
	fmt.Println("Nodo2:", nodeIndex2)
	globalMatrix.Set(nodeIndex1-1, nodeIndex1-1, globalMatrix.At(nodeIndex1-1, nodeIndex1-1)+localMatrix.At(0, 0))
	globalMatrix.Set(nodeIndex1-1, nodeIndex2-1, globalMatrix.At(nodeIndex1-1, nodeIndex2-1)+localMatrix.At(0, 1))
	globalMatrix.Set(nodeIndex2-1, nodeIndex1-1, globalMatrix.At(nodeIndex2-1, nodeIndex1)+localMatrix.At(1, 0))
	globalMatrix.Set(nodeIndex2-1, nodeIndex2-1, globalMatrix.At(nodeIndex2-1, nodeIndex2-1)+localMatrix.At(1, 1))
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