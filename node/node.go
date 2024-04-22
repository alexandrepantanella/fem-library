package node

// Node1D rappresenta un nodo in un sistema monodimensionale
type Node1D struct {
	ID         int          // Identificativo univoco del nodo
	Coordinate float64      // Coordinata del nodo
	Constraint Constraint1D // Vincoli sul nodo
}

// Constraint rappresenta i vincoli applicati a un nodo
type Constraint1D struct {
	X  int // Vincoli sulle direzioni x
	}

// Vincoli
func (n *Node1D) SetFreeCostrain() { // Libero
	n.Constraint.X = 0
}
func (n *Node1D) SetFixedCostrain() { // Incastro
	n.Constraint.X = 1
}

// Node2D rappresenta un nodo in un sistema bidimensionale
type Node2D struct {
	ID          int          // Identificativo univoco del nodo
	Coordinates [2]float64   // Coordinate x, y del nodo
	Constraint  Constraint2D // Vincoli sul nodo
}

// Constraint rappresenta i vincoli applicati a un nodo
type Constraint2D struct {
	X, Y int // Vincoli sulle direzioni x, y
	RZ   int // Vincoli di rotazione attorno a z
}

// Vincoli
func (n *Node2D) SetFreeCostrain() { // Libero
	n.Constraint.X = 0
	n.Constraint.Y = 0
	n.Constraint.RZ = 0
}
func (n *Node2D) SetFixedCostrain() { // Incastro
	n.Constraint.X = 1
	n.Constraint.Y = 1
	n.Constraint.RZ = 1
}
func (n *Node2D) SetPinnedCostrain() { // Cerniera
	n.Constraint.X = 1
	n.Constraint.Y = 1
	n.Constraint.RZ = 0
}
func (n *Node2D) SetXRollerCostrain() { // Carrello lungo x
	n.Constraint.X = 0
	n.Constraint.Y = 1
	n.Constraint.RZ = 0
}
func (n *Node2D) SetYRollerCostrain() { // Carrello lungo y
	n.Constraint.X = 1
	n.Constraint.Y = 0
	n.Constraint.RZ = 0
}
func (n *Node2D) SetArbitratyCostrain(Values [3]int) { //Arbitrario
	n.Constraint.X = Values[0]
	n.Constraint.Y = Values[1]
	n.Constraint.RZ = Values[2]
}

// Node3D rappresenta un nodo in un sistema tridimensionale
type Node3D struct {
	ID          int          // Identificativo univoco del nodo
	Coordinates [3]float64   // Coordinate x, y, z del nodo
	Constraint  Constraint3D // Vincoli sul nodo
}

// Constraint rappresenta i vincoli applicati a un nodo
type Constraint3D struct {
	X, Y, Z    int // Vincoli sulle direzioni x, y, z
	RX, RY, RZ int // Vincoli di rotazione
}

// Vincoli
func (n *Node3D) SetFreeCostrain() { // Libero
	n.Constraint.X = 0
	n.Constraint.Y = 0
	n.Constraint.Z = 0
	n.Constraint.RX = 0
	n.Constraint.RY = 0
	n.Constraint.RZ = 0
}
func (n *Node3D) SetFixedCostrain() { // Incastro
	n.Constraint.X = 1
	n.Constraint.Y = 1
	n.Constraint.Z = 1
	n.Constraint.RX = 1
	n.Constraint.RY = 1
	n.Constraint.RZ = 1
}
func (n *Node3D) SetPinnedCostrain() { // Cerniera
	n.Constraint.X = 1
	n.Constraint.Y = 1
	n.Constraint.Z = 1
	n.Constraint.RX = 0
	n.Constraint.RY = 0
	n.Constraint.RZ = 0
}
func (n *Node3D) SetXRollerCostrain() { // Carrello lungo x
	n.Constraint.X = 0
	n.Constraint.Y = 1
	n.Constraint.Z = 1
	n.Constraint.RX = 0
	n.Constraint.RY = 0
	n.Constraint.RZ = 0
}
func (n *Node3D) SetYRollerCostrain() { // Carrello lungo y
	n.Constraint.X = 1
	n.Constraint.Y = 0
	n.Constraint.Z = 1
	n.Constraint.RX = 0
	n.Constraint.RY = 0
	n.Constraint.RZ = 0
}
func (n *Node3D) SetZRollerCostrain() { // Carrello lungo z
	n.Constraint.X = 1
	n.Constraint.Y = 1
	n.Constraint.Z = 0
	n.Constraint.RX = 0
	n.Constraint.RY = 0
	n.Constraint.RZ = 0
}
func (n *Node3D) SetArbitratyCostrain(Values [6]int) { //Arbitrario
	n.Constraint.X = Values[0]
	n.Constraint.Y = Values[1]
	n.Constraint.Z = Values[2]
	n.Constraint.RX = Values[3]
	n.Constraint.RY = Values[4]
	n.Constraint.RZ = Values[5]
}
