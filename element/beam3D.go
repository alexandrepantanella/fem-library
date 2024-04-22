package element

import (
	"math"

	"gonum.org/v1/gonum/mat"

	"github.com/fem-library/material"
	"github.com/fem-library/node"
	"github.com/fem-library/section"
	"github.com/fem-library/transform"
)

type Beam3D struct {
	ID          int64                       // ID del beam
	Material    material.Material           // Materiale del beam
	Section     section.GeometricProperties //Proprietà geometriche
	Node1       node.Node3D                 // Nodo 1
	Node2       node.Node3D                 // Nodo 2
	ThetaX      float64                     // Angolo con asse X
	ThetaY      float64                     // Angolo con asse Y
	ThetaZ      float64                     // Angolo con asse z
	Rotation    float64                     // Rotazione del beam
	Moment      float64                     // Momento applicato al beam
	Constraints Constraint3D                // Vincoli sui nodi del beam
	Forces      Force3D                     // Forze applicate ai nodi del beam
}

// Constraint rappresenta i vincoli applicati a un nodo
type Constraint3D struct {
	X, Y, Z int // Vincoli sulle direzioni x, y
	//alpha, beta, gamma int // Vincolo di rotazione
}

// Force rappresenta le forze applicate a un nodo
type Force3D struct {
	Fx, Fy, Fz float64 // Forze nelle direzioni x, y,z
	Mx, My, Mz float64 // Momento applicato al nodo
}

// Metodo per calcolare la sua lunghezza
func (b *Beam3D) Length() float64 {
	return math.Sqrt(
		math.Pow(b.Node2.Coordinates[0]-b.Node1.Coordinates[0], 2) +
			math.Pow(b.Node2.Coordinates[1]-b.Node1.Coordinates[1], 2) +
			math.Pow(b.Node2.Coordinates[2]-b.Node1.Coordinates[2], 2))
}

func (b *Beam3D) SetAxisAngle() {
	// Calcola la differenza vettoriale tra i punti
	dx := b.Node2.Coordinates[0] - b.Node1.Coordinates[0]
	dy := b.Node2.Coordinates[1] - b.Node1.Coordinates[1]
	dz := b.Node2.Coordinates[2] - b.Node1.Coordinates[2]

	if dx != 0 || dz != 0 {
		b.ThetaX = math.Atan2(dz, dy)
	} else {
		// Se dx e dz sono entrambi zero, imposta l'angolo a 90° o -90° a seconda del segno di dy
		if dy >= 0 {
			b.ThetaX = math.Pi / 2 // 90°
		} else {
			b.ThetaX = -math.Pi / 2 // -90°
		}
	}

	if dx != 0 {
		b.ThetaY = math.Atan2(dz, dx)
	} else {
		// Se dx è zero, imposta l'angolo a 90° o -90° a seconda del segno di dz
		if dz >= 0 {
			b.ThetaY = math.Pi / 2 // 90°
		} else {
			b.ThetaY = -math.Pi / 2 // -90°
		}
	}

	if dx != 0 || dy != 0 {
		b.ThetaZ = math.Atan2(dy, dx)
	} else {
		// Se dx e dy sono entrambi zero, imposta l'angolo a 0°
		b.ThetaZ = 0
	}

}

// Metodo per restituire i gradi di libertà
func (b *Beam3D) DoF() int {
	return 3
}

// Metodo per restituire il suo numero
func (b *Beam3D) ElementNumber() int64 {
	return b.ID
}

// StiffnessMatrix calcola la matrice di rigidità del beam
func (b *Beam3D) StiffnessMatrix() *mat.Dense {

	// Calcola i coefficienti per la matrice di rigidità
	A := b.Section.Area
	L := b.Length()
	E := b.Material.YoungModulus
	Iy := b.Section.Iy
	Iz := b.Section.Iz
	J := b.Section.J
	// Crea la matrice di rigidità 6x6
	stiffnessMatrix := mat.NewDense(12, 12, nil)

	// Imposta i valori della matrice
	stiffnessMatrix.Set(0, 0, A*E/L)
	stiffnessMatrix.Set(0, 6, -A*E/L)
	stiffnessMatrix.Set(6, 0, -A*E/L)
	stiffnessMatrix.Set(6, 6, A*E/L)

	stiffnessMatrix.Set(1, 1, 12*Iz*E/(L*L*L))
	stiffnessMatrix.Set(1, 5, 6*Iz*E/(L*L))
	stiffnessMatrix.Set(1, 7, -12*Iz*E/(L*L*L))
	stiffnessMatrix.Set(1, 11, 6*Iz*E/(L*L))

	stiffnessMatrix.Set(2, 2, 12*Iy*E/(L*L*L))
	stiffnessMatrix.Set(2, 4, -6*Iy*E/(L*L))
	stiffnessMatrix.Set(2, 8, -12*Iy*E/(L*L*L))
	stiffnessMatrix.Set(2, 10, -6*Iy*E/(L*L))

	stiffnessMatrix.Set(3, 3, J*E/L)
	stiffnessMatrix.Set(3, 9, -J*E/L)
	stiffnessMatrix.Set(9, 3, -J*E/L)
	stiffnessMatrix.Set(9, 9, J*E/L)

	stiffnessMatrix.Set(4, 2, -6*Iy*E/(L*L))
	stiffnessMatrix.Set(4, 4, 4*Iy*E/L)
	stiffnessMatrix.Set(4, 8, 6*Iy*E/(L*L))
	stiffnessMatrix.Set(4, 10, 2*Iy*E/L)

	stiffnessMatrix.Set(5, 1, 6*Iz*E/(L*L))
	stiffnessMatrix.Set(5, 5, 4*Iz*E/L)
	stiffnessMatrix.Set(5, 7, -6*Iz*E/(L*L))
	stiffnessMatrix.Set(5, 11, 2*Iz*E/L)

	stiffnessMatrix.Set(7, 1, -12*Iz*E/(L*L*L))
	stiffnessMatrix.Set(7, 5, -6*Iz*E/(L*L))
	stiffnessMatrix.Set(7, 7, 12*Iz*E/(L*L*L))
	stiffnessMatrix.Set(7, 11, -6*Iz*E/(L*L))

	stiffnessMatrix.Set(8, 2, -12*Iy*E/(L*L*L))
	stiffnessMatrix.Set(8, 4, 6*Iy*E/(L*L))
	stiffnessMatrix.Set(8, 8, 12*Iy*E/(L*L*L))
	stiffnessMatrix.Set(8, 10, 6*Iy*E/(L*L))

	stiffnessMatrix.Set(10, 2, -6*Iy*E/(L*L))
	stiffnessMatrix.Set(10, 4, 2*Iy*E/L)
	stiffnessMatrix.Set(10, 8, 6*Iy*E/(L*L))
	stiffnessMatrix.Set(10, 10, 4*Iy*E/L)

	stiffnessMatrix.Set(11, 1, 6*Iz*E/(L*L))
	stiffnessMatrix.Set(11, 5, 2*Iz*E/L)
	stiffnessMatrix.Set(11, 7, -6*Iz*E/(L*L))
	stiffnessMatrix.Set(11, 11, 4*Iz*E/L)

	return stiffnessMatrix
}

// MassMatrix calcola la matrice delle masse del beam
func (b *Beam3D) MassMatrix() *mat.Dense {
	// Estrae la densità del materiale
	density := b.Material.Density

	// Calcola la lunghezza del beam
	length := b.Length()

	// Calcola l'area trasversale
	area := b.Section.Area

	// Calcola la massa per unità di lunghezza
	massPerLength := density * area * length

	// Crea la matrice delle masse 12x12
	massMatrix := mat.NewDense(12, 12, nil)

	// Imposta i valori della matrice
	massMatrix.Set(0, 0, massPerLength/3)
	massMatrix.Set(1, 1, massPerLength/3)
	massMatrix.Set(3, 3, massPerLength/3)
	massMatrix.Set(4, 4, massPerLength/3)
	massMatrix.Set(5, 5, massPerLength/3)
	massMatrix.Set(6, 6, massPerLength/3)
	massMatrix.Set(0, 3, massPerLength/6)
	massMatrix.Set(1, 4, massPerLength/6)
	massMatrix.Set(3, 0, massPerLength/6)
	massMatrix.Set(4, 1, massPerLength/6)

	return massMatrix
}

func (b *Beam3D) GlobalStiffnessMatrix() *mat.Dense {
	// Calcola gli angoli con gli assi
	b.SetAxisAngle()

	// Calcola la trasposta della matrice di trasformazione
	transformation := *transform.TransformMatrix3D(b.ThetaX, b.ThetaY, b.ThetaZ)
	transposedTransformation := transformation.T()

	// Moltiplica T^T * K_local
	var temp1 mat.Dense
	temp1.Mul(transposedTransformation, b.StiffnessMatrix())

	// Moltiplica (T^T * K_local) * T
	var globalStiffness mat.Dense
	globalStiffness.Mul(&temp1, &transformation)

	return &globalStiffness
}

// Metodo per applicare i vincoli ai nodi del beam
func (b *Beam3D) ApplyConstraints(constraints Constraint3D) {
	b.Constraints = constraints
}

// Metodo per applicare le forze ai nodi del beam
func (b *Beam3D) ApplyForces(forces Force3D) {
	b.Forces = forces
}
