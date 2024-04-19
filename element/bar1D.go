package element

import (
	"math"

	"github.com/fem-library/material"
	"github.com/fem-library/node"
	"github.com/fem-library/section"
	"gonum.org/v1/gonum/mat"
)

type Bar1D struct {
	ID       int64                       // ID dell'asta
	Material material.Material           // Materiale dell'asta
	Node1    node.Node1D                 // Nodo iniziale
	Node2    node.Node1D                 // Nodo finale
	Section  section.GeometricProperties // Proprietà geometriche della sezione
	KLocal   *mat.Dense                  //Matrice di rigidezza locale
	KGlobal  *mat.Dense                  //Matrice di rigidezza globale
	Mass     *mat.Dense                  //Matrice delle masse
}

// Metodo per calcolare la sua lunghezza
func (r *Bar1D) Length() float64 {
	return math.Abs(r.Node2.Coordinate - r.Node1.Coordinate)
}

// Metodo per restituire il suo numero
func (r *Bar1D) ElementNumber() int64 {
	return r.ID
}

// Metodo per restituire la matrice di rigidezza locale
func (r *Bar1D) GetKLocal() *mat.Dense {
	if r.KLocal == nil {
		r.KLocal = r.StiffnessMatrix()
	}
	return r.KLocal
}

// Metodo per restituire la matrice di rigidezza globale
func (r *Bar1D) GetKGlobal() *mat.Dense {
	if r.KGlobal == nil {
		r.KGlobal = r.GlobalStiffnessMatrix()
	}
	return r.KGlobal
}

// Metodo per restituire la matrice delle masse
func (r *Bar1D) GetMass() *mat.Dense {
	if r.Mass == nil {
		r.Mass = r.MassMatrix()
	}
	return r.Mass
}

// StiffnessMatrix calcola la matrice di rigidità dell'asta
func (r *Bar1D) StiffnessMatrix() *mat.Dense {
	length := r.Length()
	area := r.Section.Area
	modulus := r.Material.YoungModulus

	k := (modulus * area) / length
	stiffnessMatrix := mat.NewDense(2, 2, nil)
	stiffnessMatrix.Set(0, 0, k)
	stiffnessMatrix.Set(0, 1, -k)
	stiffnessMatrix.Set(1, 0, -k)
	stiffnessMatrix.Set(1, 1, k)

	return stiffnessMatrix
}

// MassMatrix calcola la matrice delle masse dell'asta
func (r *Bar1D) MassMatrix() *mat.Dense {
	length := r.Length()
	density := r.Material.Density
	area := r.Section.Area
	mass := density * length * area
	massMatrix := mat.NewDense(2, 2, nil)
	massMatrix.Set(0, 0, mass/3)
	massMatrix.Set(0, 1, mass/6)
	massMatrix.Set(1, 0, mass/6)
	massMatrix.Set(1, 1, mass/3)

	return massMatrix
}

// GlobalStiffnessMatrix calcola la matrice di rigidità globale per un'asta in 1D
func (r *Bar1D) GlobalStiffnessMatrix() *mat.Dense {
	// Calcola la matrice di rigidità locale
	return r.StiffnessMatrix()
}
