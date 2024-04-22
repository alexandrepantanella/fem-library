package element

import (
	"math"

	"github.com/fem-library/material"
	"github.com/fem-library/node"
	"github.com/fem-library/section"
	"gonum.org/v1/gonum/mat"
)

type Beam1D struct {
	ID       int64                       //ID del beam
	Material material.Material           // Materiale del beam
	Node1    node.Node1D                 // Nodo 1
	Node2    node.Node1D                 // Nodo 2
	Section  section.GeometricProperties //Proprietà geometriche
	KLocal   *mat.Dense                  //Matrice di rigidezza locale
	KGlobal  *mat.Dense                  //Matrice di rigidezza globale
	Mass     *mat.Dense                  //Matrice delle masse
}

// Metodo per calcolare la sua lunghezza
func (b *Beam1D) Length() float64 {
	return math.Abs(b.Node2.Coordinate - b.Node1.Coordinate)
}

// Metodo per restituire i gradi di libertà
func (b *Beam1D) DoF() int {
	return 1
}

// Metodo per restituire il suo numero
func (b *Beam1D) ElementNumber() int64 {
	return b.ID
}

// StiffnessMatrix calcola la matrice di rigidità del beam
func (b *Beam1D) StiffnessMatrix() *mat.Dense {
	length := b.Length()
	area := b.Section.Area
	modulus := b.Material.YoungModulus
	k := (modulus * area) / length
	stiffnessMatrix := mat.NewDense(2, 2, nil)
	stiffnessMatrix.Set(0, 0, k)
	stiffnessMatrix.Set(0, 1, -k)
	stiffnessMatrix.Set(1, 0, -k)
	stiffnessMatrix.Set(1, 1, k)
	return stiffnessMatrix
}

// MassMatrix calcola la matrice delle masse del beam
func (b *Beam1D) MassMatrix() *mat.Dense {
	length := b.Length()
	density := b.Material.Density
	mass := density * length
	massMatrix := mat.NewDense(2, 2, nil)
	massMatrix.Set(0, 0, mass/3)
	massMatrix.Set(0, 1, mass/6)
	massMatrix.Set(1, 0, mass/6)
	massMatrix.Set(1, 1, mass/3)
	return massMatrix
}

// GlobalStiffnessMatrix calcola la matrice di rigidità globale per una trave monodimensionale
func (b *Beam1D) GlobalStiffnessMatrix() *mat.Dense {
	// Calcola la matrice di rigidità locale
	return b.StiffnessMatrix()
}
