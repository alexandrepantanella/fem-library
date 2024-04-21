package test

import (
	"fmt"

	"github.com/fem-library/element"
	"github.com/fem-library/material"
	"github.com/fem-library/node"
	"github.com/fem-library/section"
	"gonum.org/v1/gonum/mat"
)

func TestBeam3DStiffnessMatrix() {
	// Definizione dei parametri per il beam
	beam := element.Beam3D{
		ID: 1,
		Material: material.Material{
			YoungModulus: 210e9, // Pa
			Density:      7850,  // kg/m^3
		},
		Section: section.GeometricProperties{
			Area: 0.01, // m^2
			Iy:   1e-5, // m^4
			Iz:   1e-5, // m^4
			J:    1e-5, // m^4
		},
		Node1: node.Node3D{
			ID:          1,
			Coordinates: [3]float64{0.0, 0.0, 0.0},
		},
		Node2: node.Node3D{
			ID:          2,
			Coordinates: [3]float64{1.0, 0.0, 0.0},
		},
	}

	// Calcolo della lunghezza del beam
	length := beam.Length()

	// Calcolo della matrice di rigidità attesa
	E := beam.Material.YoungModulus
	A := beam.Section.Area
	Iy := beam.Section.Iy
	Iz := beam.Section.Iz
	J := beam.Section.J

	stiffnessMatrixExpected := mat.NewDense(12, 12, []float64{
		A * E / length, 0, 0, 0, 0, 0, -A * E / length, 0, 0, 0, 0, 0,
		0, 12 * Iz * E / (length * length * length), 0, 0, 0, 6 * Iz * E / (length * length), 0, -12 * Iz * E / (length * length * length), 0, 0, 0, 6 * Iz * E / (length * length),
		0, 0, 12 * Iy * E / (length * length * length), 0, -6 * Iy * E / (length * length), 0, 0, 0, -12 * Iy * E / (length * length * length), 0, -6 * Iy * E / (length * length), 0,
		0, 0, 0, J * E / length, 0, 0, 0, 0, 0, -J * E / length, 0, 0,
		0, 0, -6 * Iy * E / (length * length), 0, 4 * Iy * E / length, 0, 0, 0, 6 * Iy * E / (length * length), 0, 2 * Iy * E / length, 0,
		0, 6 * Iz * E / (length * length), 0, 0, 0, 4 * Iz * E / length, 0, -6 * Iz * E / (length * length), 0, 0, 0, 2 * Iz * E / length,
		-A * E / length, 0, 0, 0, 0, 0, A * E / length, 0, 0, 0, 0, 0,
		0, -12 * Iz * E / (length * length * length), 0, 0, 0, -6 * Iz * E / (length * length), 0, 12 * Iz * E / (length * length * length), 0, 0, 0, -6 * Iz * E / (length * length),
		0, 0, -12 * Iy * E / (length * length * length), 0, 6 * Iy * E / (length * length), 0, 0, 0, 12 * Iy * E / (length * length * length), 0, 6 * Iy * E / (length * length), 0,
		0, 0, 0, -J * E / length, 0, 0, 0, 0, 0, J * E / length, 0, 0,
		0, 0, -6 * Iy * E / (length * length), 0, 2 * Iy * E / length, 0, 0, 0, 6 * Iy * E / (length * length), 0, 4 * Iy * E / length, 0,
		0, 6 * Iz * E / (length * length), 0, 0, 0, 2 * Iz * E / length, 0, -6 * Iz * E / (length * length), 0, 0, 0, 4 * Iz * E / length,
	})

	// Calcolo della matrice di rigidità effettiva
	stiffnessMatrixActual := beam.StiffnessMatrix()

	fmt.Println("Matrice:")
	fmt.Printf("%v\n", mat.Formatted(stiffnessMatrixActual))

	// Verifica della correttezza
	if !mat.Equal(stiffnessMatrixExpected, stiffnessMatrixActual) {
		fmt.Printf("Stiffness matrix calculation incorrect, expected:\n%v\ngot:\n%v", mat.Formatted(stiffnessMatrixExpected), mat.Formatted(stiffnessMatrixActual))
	}
}
