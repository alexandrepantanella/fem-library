package main

import (
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
	Node			[]Node				// Array of nodes
	Element1D		[]Element1D			// Array of elements
	Element2D		[]Element2D			// Array of elements
	Element3D		[]Element3D			// Array of elements
	ElementForce	[]ElementForce		// Array of force elements
	NodeForce		[]NodalForce		// Array of force elements
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

type Node struct {
	Id          int          // Unique identifier of the node
	Coordinate  []float64    // Node coordinate X,Y,Z
	Constrains  []bool       // Constraints on the node
}

type NodalForce struct {
	Id          int            // Unique identifier of the node
	F           []float64      // Force applied on node
	M           []float64      // Moments applied on node
}

type ElementForce struct {
	Id          int            // Unique identifier of the element
	F           []float64      // Distributed Force applied on element
	M           []float64      // Distributed Moments applied on element
}

type Element1D struct {			// One dimensional elements
	Id      int                 // Element ID
	Type	string				// Element type: BAR - SPRING - BEAM 
	K		float64				// Spring elastic constant
	L       float64             // Length 
	A       float64             // Cross-sectional area
	E       float64             // Modulus of elasticity
	G		float64				//
	Ix 		float64				// Torsion constant 
	Iy		float64				// Second moment of area 
	Iz		float64				// Second moment of area  
	D       float64             // Density
	N1      int                 // Initial node
	N2      int                 // Final node
	Theta	[]float64			// Element angles with x,y and z
	KL		*mat.Dense			// Local Stifness Matrix
}

type Element2D struct {			// Two dimension elements
	Id      int                 // Element ID
	Type	string				// Element type: 
}

type Element3D struct {			// Three dimension elements
	Id      int                 // Element ID
	Type	string				// Element type: 
}