package D1

import (
	"gonum.org/v1/gonum/mat"
)

type Analysis struct {
	Dim         int         	// Model dimensions
	Type        string      	// SL or SNL or DL
	Behaviour	string      	// ISO or ORTHO (material)
	Spring      []Spring		// Spring elements
	Bar         []Bar			// Bar elements
	Beam        []Beam			// Beam elements
	Node    	[]Node			// Spring elements
	Kg 			mat.Dense   	// Global stiffness matrix
	Force	*	mat.VecDense	// Vector of forces applied to nodes in the global reference system
}

type Spring struct {
	Id      int                 // Spring ID
	K       float64             // Spring constant
	N1      int                 // Initial node
	N2      int                 // Final node
	Kl      *mat.Dense          // Local stiffness matrix
	Kg      *mat.Dense          // Global stiffness matrix
}

type Bar struct {
	Id      int                 // Bar ID
	L       float64             // Length 
	A       float64             // Cross-sectional area
	E       float64             // Modulus of elasticity
	D       float64             // Density
	N1      int                 // Initial node
	N2      int                 // Final node
	Kl      *mat.Dense          // Local stiffness matrix
	Kg      *mat.Dense          // Global stiffness matrix
	Mass    *mat.Dense          // Mass matrix
	//F    force.Force1D         // Vector of forces applied to the element
}

type Beam struct {
	Id      int                 // Beam ID
	L       float64             // Length 
	A       float64             // Cross-sectional area
	E       float64             // Modulus of elasticity
	D       float64             // Density
	N1      int                 // Initial node
	N2      int                 // Final node
	Kl      *mat.Dense          // Local stiffness matrix
	Kg      *mat.Dense          // Global stiffness matrix
	Mass    *mat.Dense          // Mass matrix
	//F    force.Force1D         // Vector of forces applied to the element
}

// Node1D represents a node in a one-dimensional system
type Node struct {
	Id          int          // Unique identifier of the node
	X           float64      // Node coordinate
	C           bool         // Constraints on the node
}
