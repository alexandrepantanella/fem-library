package D1

import (
	"fmt"
	"math"

	"gonum.org/v1/gonum/mat"
)

/****************************/
/*	Start Analisys methods  */
/****************************/

// Method to count and set the distinct nodes
func (s *Analysis) CountDistinctNodes() {
	nodeMap := make(map[int]bool)
	for _, spring := range s.Spring {
		nodeMap[spring.N1] = true
		nodeMap[spring.N2] = true
	}
	for _, beam := range s.Beam {
		nodeMap[beam.N1] = true
		nodeMap[beam.N2] = true
	}
	for _, bar := range s.Bar {
		nodeMap[bar.N1] = true
		nodeMap[bar.N2] = true
	}
	s.NumNode = len(nodeMap)
}

// Method to count and set the elements
func (s *Analysis) CountElements() {
	s.NumElement = len(s.Spring) + len(s.Bar) + len(s.Beam)
}

//Method to set all length
func (s *Analysis) SetLength() {
	for _, element := range s.Bar {
		element.SetL(s.Node)
	}
	for _, element := range s.Beam {
		element.SetL(s.Node)
	}
}

// Method to get the constraints (nodes with C=true) as a map[int]bool
func (s *Analysis) SetConstraints() {
	constraints := make(map[int]bool, s.NumNode)
	for i := 1; i <= s.NumNode; i++ {
		constraints[i] = false
	}
	for _, element := range s.Spring {
		node1, _ := getNodeByID(s.Node, element.N1)
		node2, _ := getNodeByID(s.Node, element.N2)
		if node1.C {
			constraints[node1.Id] = true
		}
		if node2.C {
			constraints[node2.Id] = true
		}
	}
	for _, element := range s.Bar {
		node1, _ := getNodeByID(s.Node, element.N1)
		node2, _ := getNodeByID(s.Node, element.N2)
		if node1.C {
			constraints[node1.Id] = true
		}
		if node2.C {
			constraints[node2.Id] = true
		}
	}
	for _, element := range s.Beam {
		node1, _ := getNodeByID(s.Node, element.N1)
		node2, _ := getNodeByID(s.Node, element.N2)
		if node1.C {
			constraints[node1.Id] = true
		}
		if node2.C {
			constraints[node2.Id] = true
		}
	}
	s.Constraints = constraints
}

// Method to Get the final diplacements
func (s *Analysis) CalculateDisplacements(u mat.VecDense) {
	newVec := mat.NewVecDense(s.NumNode, nil)
	j:=0
	for i := 0; i < s.NumNode; i++ {
        if s.Constraints[i+1] { 
            newVec.SetVec(i, 0)
        } else {
            newVec.SetVec(i, u.AtVec(j))
			j++
        }
    }
	s.Output.Displacement = *newVec
}

// Method to Get the final reactions
func (s *Analysis) CalculateReactions() {
	reaction := mat.NewVecDense(s.NumNode, nil)
	reaction.MulVec(&s.Kg, &s.Output.Displacement)
	s.Output.Reaction = *reaction
}

// Auxiliary Function to get a node by its ID
func getNodeByID(nodes []Node, id int) (Node, error) {
	for _, node := range nodes {
		if node.Id == id {
			return node, nil
		}
	}
	return Node{}, fmt.Errorf("node with ID %d not found", id)
}

/****************************/
/*	 End Analisys methods   */
/****************************/

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