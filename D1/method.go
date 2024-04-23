package D1

import (
	"math"

	"gonum.org/v1/gonum/mat"
)

/**************************/
/*	Start Spring methods  */
/**************************/

// Method to get Id
func (s *Spring) GetID() int {
	return s.Id
}

// Method to set Id
func (s *Spring) SetID(Id int) {
	s.Id = Id
}

// Method to set Local StiffnessMatrix
func (s *Spring) SetKL(){
	k :=  s.K
	stiffnessMatrix := mat.NewDense(2,2, nil)
	stiffnessMatrix.Set(0, 0, k)
	stiffnessMatrix.Set(0, 1, -k)
	stiffnessMatrix.Set(1, 0, -k)
	stiffnessMatrix.Set(1, 1, k)

	s.Kl = stiffnessMatrix
}

// Method to set Global StiffnessMatrix
func (s *Spring) SetKG() {
	s.Kg = s.Kl
}

/**************************/
/*	 End Spring methods   */
/**************************/

/***********************/
/*	Start Bar methods  */
/***********************/

// Method to get Id
func (r *Bar) GetID() int {
	return r.Id
}

// Method to set Id
func (r *Bar) SetID(Id int) {
	r.Id = Id
}

// Method to set Length
func (r *Bar) SetL(nodes []Node) {
	node1 := findNodeByID(nodes, r.N1)
	node2 := findNodeByID(nodes, r.N2)
	if node1 != nil && node2 != nil {
		r.L = math.Abs(node2.X - node1.X)
	}
}

// Method to set Local StiffnessMatrix
func (r *Bar) SetKL(){
	k := (r.E * r.A) / r.L
	stiffnessMatrix := mat.NewDense(2,2, nil)
	stiffnessMatrix.Set(0, 0, k)
	stiffnessMatrix.Set(0, 1, -k)
	stiffnessMatrix.Set(1, 0, -k)
	stiffnessMatrix.Set(1, 1, k)

	r.Kl = stiffnessMatrix
}

// Method to set Mass Matrix
func (r *Bar) SetM() {
	mass :=  r.D * r.L * r.A
	massMatrix := mat.NewDense(2, 2, nil)
	massMatrix.Set(0, 0, mass/3)
	massMatrix.Set(0, 1, mass/6)
	massMatrix.Set(1, 0, mass/6)
	massMatrix.Set(1, 1, mass/3)
	r.Mass = massMatrix
}

// // Method to set Global StiffnessMatrix
func (r *Bar) SetKG() {
	r.Kg = r.Kl
}

/***********************/
/*	End Bar methods    */
/***********************/

/************************/
/*	Start Beam methods  */
/************************/

// Method to get Id
func (b *Beam) GetID() int {
	return b.Id
}

// Method to set Id
func (b *Beam) SetID(Id int) {
	b.Id = Id
}

// Method to set Length
func (b *Beam) SetL(nodes []Node) {
	node1 := findNodeByID(nodes, b.N1)
	node2 := findNodeByID(nodes, b.N2)
	if node1 != nil && node2 != nil {
		b.L = math.Abs(node2.X - node1.X)
	}
}

// Method to set Local StiffnessMatrix
func (b *Beam) SetKL(){
	k := (b.E * b.A) / b.L
	stiffnessMatrix := mat.NewDense(2,2, nil)
	stiffnessMatrix.Set(0, 0, k)
	stiffnessMatrix.Set(0, 1, -k)
	stiffnessMatrix.Set(1, 0, -k)
	stiffnessMatrix.Set(1, 1, k)

	b.Kl = stiffnessMatrix
}

// Method to set Mass Matrix
func (b *Beam) SetM() {
	mass :=  b.D * b.L * b.A
	massMatrix := mat.NewDense(2, 2, nil)
	massMatrix.Set(0, 0, mass/3)
	massMatrix.Set(0, 1, mass/6)
	massMatrix.Set(1, 0, mass/6)
	massMatrix.Set(1, 1, mass/3)
	b.Mass = massMatrix
}

// // Method to set Global StiffnessMatrix
func (b *Beam) SetKG() {
	b.Kg = b.Kl
}

/***********************/
/*	End Beam methods   */
/***********************/

/***********************/
/*	Utility functions  */
/***********************/

// Utility function
func findNodeByID(nodes []Node, nodeID int) *Node {
	for _, node := range nodes {
		if node.Id == nodeID {
			return &node
		}
	}
	return nil 
}