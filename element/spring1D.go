package element

import (
	"github.com/fem-library/node"
	"github.com/gonum/matrix/mat64"
)

type Spring1D struct {
	ID    int64       // ID della molla
	Node1 node.Node1D // Nodo iniziale
	Node2 node.Node1D // Nodo finale
	K     float64     // Costante elastica
}

// Metodo per restituire il suo numero
func (r *Spring1D) ElementNumber() int64 {
	return r.ID
}

// StiffnessMatrix calcola la matrice di rigidità della molla
func (r *Spring1D) StiffnessMatrix() *mat64.Dense {

	k := r.K
	stiffnessMatrix := mat64.NewDense(2, 2, nil)

	// Assegna i valori della matrice di rigidezza
	stiffnessMatrix.Set(0, 0, k)
	stiffnessMatrix.Set(0, 1, -k)
	stiffnessMatrix.Set(1, 0, k)
	stiffnessMatrix.Set(1, 1, -k)

	return stiffnessMatrix
}

// GlobalStiffnessMatrix calcola la matrice di rigidità globale per una molla in 1D
func (r *Spring1D) GlobalStiffnessMatrix() *mat64.Dense {
	// Calcola la matrice di rigidità locale
	return r.StiffnessMatrix()
}
