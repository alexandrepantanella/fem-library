package element

import "gonum.org/v1/gonum/mat"

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