package element

import (
	"github.com/fem-library/node"
	"github.com/gonum/matrix/mat64"
)

type Spring2D struct {
	ID    int64       // ID della molla
	Node1 node.Node1D // Nodo iniziale
	Node2 node.Node1D // Nodo finale
	K     float64     // Costante elastica
}

// Metodo per restituire il suo numero
func (r *Spring2D) ElementNumber() int64 {
	return r.ID
}

// StiffnessMatrix calcola la matrice di rigidità della molla
func (r *Spring2D) StiffnessMatrix() *mat64.Dense {

	// Calcola i coefficienti per la matrice di rigidità
	k := r.K

	// Crea la matrice di rigidità 12x12
	stiffnessMatrix := mat64.NewDense(6, 6, nil)

	stiffnessMatrix.Set(0, 0, k)
	stiffnessMatrix.Set(0, 3, -k)
	stiffnessMatrix.Set(3, 0, -k)
	stiffnessMatrix.Set(3, 3, k)

	return stiffnessMatrix
}

// GlobalStiffnessMatrix calcola la matrice di rigidità globale per una molla in 1D
func (r *Spring2D) GlobalStiffnessMatrix() *mat64.Dense {
	// Calcola la matrice di rigidità locale
	return r.StiffnessMatrix()
}
