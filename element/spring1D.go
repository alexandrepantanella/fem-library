package element

import (
	"github.com/fem-library/force"
	"github.com/fem-library/node"
	"gonum.org/v1/gonum/mat"
)

type Spring1D struct {
	ID    int64               // ID della molla
	Node1 node.Node1D         // Nodo iniziale
	Node2 node.Node1D         // Nodo finale
	K     float64             // Costante elastica
	KLocal   *mat.Dense       // Matrice di rigidezza locale
	KGlobal  *mat.Dense       // Matrice di rigidezza globale
	Mass     *mat.Dense       // Matrice delle masse
	Force    force.Force1D    // Vettore delle forze applicate sull'elemento
}

// Metodo per restituire il suo numero
func (r *Spring1D) ElementNumber() int64 {
	return r.ID
}

// StiffnessMatrix calcola la matrice di rigidità della molla
func (r *Spring1D) StiffnessMatrix() {

	k := r.K
	stiffnessMatrix := mat.NewDense(2, 2, nil)

	// Assegna i valori della matrice di rigidezza
	stiffnessMatrix.Set(0, 0, k)
	stiffnessMatrix.Set(0, 1, -k)
	stiffnessMatrix.Set(1, 0, -k)
	stiffnessMatrix.Set(1, 1, k)

	r.KLocal =  stiffnessMatrix
}

// GlobalStiffnessMatrix calcola la matrice di rigidità globale per una molla in 1D
func (r *Spring1D) GlobalStiffnessMatrix() {
	r.KGlobal = r.KLocal
}

