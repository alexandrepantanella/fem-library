package main

import (
	"github.com/fem-library/element"
	"gonum.org/v1/gonum/mat"
)

type Analysis struct {
	InputData	InputData			// Input data
	CalcData	CalcData			// Calculated data
	OutputData	OutputData			// Output data
}

type InputData struct {
	Name			string						// Analisys name
	Description		string						// Problem description or references
	DoF         	int         				// Model Degree of Freedom (1 - 2 - 3 - 6)
	Type        	string      				// Analisys type (STATICLINEAR - STATICNONLINEAR - DYNAMICLINEAR - DYNAMICNONLINEAR - BUCKLING)
	Subtype			string						// Subtype	(SPRING1D - BAR1D - TRUSS2D - BEAM2D)
	Node			[]element.Node				// Array of nodes
	Element1D		[]element.Element1D			// Array of elements
	Element2D		[]element.Element2D			// Array of elements
	Element3D		[]element.Element3D			// Array of elements
	ElementForce	[]element.ElementForce		// Array of force elements
	NodeForce		[]element.NodalForce		// Array of force elements
}

type CalcData struct {
	NumNode			int						// Number of nodes
	NumElement		int						// Number of elements
	Length       	map[int]float64         // Length of elements
	Constraints		map[int]map[int]bool	// Map of node costrains map[node]map[dof][boolvalue] //ex. [12][0]true
	LocalStiffness	map[int]*mat.Dense		// Map of Local Stiffness matrix for elements
	GlobalStiffness	map[int]*mat.Dense		// Map of Global stiffness matrix for elements
	Mass  			map[int]*mat.Dense		// Map of Mass matrix for elements
	Force			map[int]*mat.VecDense	// Map of Vector of forces applied to nodes in the global reference system
}

type OutputData struct {
	Displacement	mat.VecDense		// Vector of displacement
	Reaction		mat.VecDense		// Vector of reactions
}