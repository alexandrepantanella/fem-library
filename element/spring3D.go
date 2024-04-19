package element

import (
	"alexcloud.it/dsm/node"
	"github.com/gonum/matrix/mat64"
)

// Spring3D rappresenta una molla tridimensionale
type Spring3D struct {
	ID    int64       // ID della molla
	Node1 node.Node3D // Nodo iniziale
	Node2 node.Node3D // Nodo finale
	K     float64     // Costante elastica
}

// Metodo per restituire il suo numero
func (r *Spring3D) ElementNumber() int64 {
	return r.ID
}

func (r *Spring3D) StiffnessMatrix() *mat64.Dense {
	// Calcola la costante elastica
	k := r.K

	// Crea la matrice di rigidezza 12x12
	stiffnessMatrix := mat64.NewDense(12, 12, nil)

	// Assegna i valori della matrice di rigidezza
	stiffnessMatrix.Set(0, 0, k)
	stiffnessMatrix.Set(0, 6, -k)
	stiffnessMatrix.Set(6, 0, -k)
	stiffnessMatrix.Set(6, 6, k)

	return stiffnessMatrix
}

// GlobalStiffnessMatrix calcola la matrice di rigidezza globale per una molla tridimensionale
func (r *Spring3D) GlobalStiffnessMatrix() *mat64.Dense {
	// La matrice di rigidezza globale è uguale alla matrice di rigidezza locale
	return r.StiffnessMatrix()
}
