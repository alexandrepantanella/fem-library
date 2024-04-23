package D1

import (
	"gonum.org/v1/gonum/mat"
)

type Analysis struct {
	Name		string				// Analisys name
	Description	string				// Problem description or references
	Dim         int         		// Model dimensions
	Type        string      		// SL or SNL or DL
	Spring      []Spring			// Spring elements
	Bar         []Bar				// Bar elements
	Beam        []Beam				// Beam elements
	Node    	[]Node				// Spring elements
	NumNode		int					// Number of nodes
	NumElement	int					// Number of elements
	Constraints	map[int]bool		// Map of node costrains
	Kg 			mat.Dense   		// Global stiffness matrix
	KgRed		mat.Dense   		// Global stiffness matrix reduced
	Force		mat.VecDense		// Vector of forces applied to nodes in the global reference system
	ForceRed	mat.VecDense		// Vector of forces reduced
	Output		Output				// Output
}

type Spring struct {
	Id      int                 // Spring ID
	K       float64             // Spring constant
	N1      int                 // Initial node
	N2      int                 // Final node
	Kl      *mat.Dense          // Local stiffness matrix
	Kg      *mat.Dense          // Global stiffness matrix
	F1      float64        		// Force on node 1
	F2      float64        		// Force on node 2
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
	F1      float64        		// Force on node 1
	F2      float64        		// Force on node 2
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
	F1      float64        		// Force on node 1
	F2      float64        		// Force on node 2
}

// Node1D represents a node in a one-dimensional system
type Node struct {
	Id          int          // Unique identifier of the node
	X           float64      // Node coordinate
	C           bool         // Constraints on the node
}

// Type Output
type Output struct {
	Displacement	mat.VecDense	// Vector of displacement
	Reaction	mat.VecDense		// Vector of reactions
}

