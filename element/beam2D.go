package element

import (
	"math"

	"gonum.org/v1/gonum/mat"

	"alexcloud.it/dsm/material"
	"alexcloud.it/dsm/node"
	"alexcloud.it/dsm/section"
)

type Beam2D struct {
	ID          int64                       // ID del beam
	Material    material.Material           // Materiale del beam
	Section     section.GeometricProperties //Proprietà geometriche
	Node1       node.Node2D                 // Nodo 1
	Node2       node.Node2D                 // Nodo 2
	Rotation    float64                     // Rotazione del beam
	Moment      float64                     // Momento applicato al beam
	Constraints Constraint                  // Vincoli sui nodi del beam
	Forces      Force                       // Forze applicate ai nodi del beam
}

// Constraint rappresenta i vincoli applicati a un nodo
type Constraint struct {
	X, Y     bool // Vincoli sulle direzioni x, y
	Rotation bool // Vincolo di rotazione
}

// Force rappresenta le forze applicate a un nodo
type Force struct {
	X, Y   float64 // Forze nelle direzioni x, y
	Moment float64 // Momento applicato al nodo
}

// Metodo per calcolare la sua lunghezza
func (b *Beam2D) Length() float64 {
	return math.Sqrt(math.Pow(b.Node2.Coordinates[0]-b.Node1.Coordinates[0], 2) + math.Pow(b.Node2.Coordinates[1]-b.Node1.Coordinates[1], 2))
}

// Metodo per restituire i gradi di libertà
func (b *Beam2D) DoF() int {
	return 2
}

// Metodo per restituire il suo numero
func (b *Beam2D) ElementNumber() int64 {
	return b.ID
}

// StiffnessMatrix calcola la matrice di rigidità del beam
func (b *Beam2D) StiffnessMatrix() *mat.Dense {
	length := b.Length()
	area := b.Section.Area
	modulus := b.Material.YoungModulus
	// Calcola i coefficienti per la matrice di rigidità
	EA := modulus * area
	L := length
	k := EA / L
	k2 := 12 * EA / (L * L * L)
	k3 := 6 * EA / (L * L)
	k4 := 4 * EA / L

	// Crea la matrice di rigidità 6x6
	stiffnessMatrix := mat.NewDense(6, 6, nil)

	// Imposta i valori della matrice
	stiffnessMatrix.Set(0, 0, k)
	stiffnessMatrix.Set(0, 3, -k)
	stiffnessMatrix.Set(1, 1, k2)
	stiffnessMatrix.Set(1, 2, k3)
	stiffnessMatrix.Set(1, 4, -k2)
	stiffnessMatrix.Set(1, 5, k3)
	stiffnessMatrix.Set(2, 1, k3)
	stiffnessMatrix.Set(2, 2, 2*k4)
	stiffnessMatrix.Set(2, 4, -k3)
	stiffnessMatrix.Set(2, 5, k4)
	stiffnessMatrix.Set(3, 0, -k)
	stiffnessMatrix.Set(3, 3, k)
	stiffnessMatrix.Set(4, 1, -k2)
	stiffnessMatrix.Set(4, 2, -k3)
	stiffnessMatrix.Set(4, 4, k2)
	stiffnessMatrix.Set(4, 5, -k3)
	stiffnessMatrix.Set(5, 1, k3)
	stiffnessMatrix.Set(5, 2, k4)
	stiffnessMatrix.Set(5, 4, -k3)
	stiffnessMatrix.Set(5, 5, 2*k4)

	return stiffnessMatrix
}

// MassMatrix calcola la matrice delle masse del beam
func (b *Beam2D) MassMatrix() *mat.Dense {
	// Estrae la densità del materiale
	density := b.Material.Density

	// Calcola la lunghezza del beam
	length := b.Length()

	// Calcola l'area trasversale
	area := b.Section.Area

	// Calcola la massa per unità di lunghezza
	massPerLength := density * area * length

	// Crea la matrice delle masse 6x6
	massMatrix := mat.NewDense(6, 6, nil)

	// Imposta i valori della matrice
	for i := 0; i < 6; i++ {
		massMatrix.Set(i, i, massPerLength/3)
		massMatrix.Set(i+6, i+6, massPerLength/3)
	}
	for i := 0; i < 6; i++ {
		massMatrix.Set(i, i+6, massPerLength/6)
		massMatrix.Set(i+6, i, massPerLength/6)
	}

	return massMatrix
}

// Metodo per applicare i vincoli ai nodi del beam
func (b *Beam2D) ApplyConstraints(constraints Constraint) {
	b.Constraints = constraints
}

// Metodo per applicare le forze ai nodi del beam
func (b *Beam2D) ApplyForces(forces Force) {
	b.Forces = forces
}
