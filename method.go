package main

import (
	"fmt"
	"math"

	"gonum.org/v1/gonum/mat"
)

// Method to get Id
func (e *Element1D) GetID() int {
	return e.Id
}
func (e *Element2D) GetID() int {
	return e.Id
}
func (e *Element3D) GetID() int {
	return e.Id
}

// Method to set Length
func (e *Element1D) SetLength(nodes []Node) {
	node1 := findNodeByID(nodes, e.N1)
	node2 := findNodeByID(nodes, e.N2)
	if node1 != nil && node2 != nil {
		e.L = math.Sqrt(math.Pow(node2.Coordinate[0] - node1.Coordinate[0],2) + 
		math.Pow(node2.Coordinate[1] - node1.Coordinate[1],2) +
		math.Pow(node2.Coordinate[2] - node1.Coordinate[2],2))
	}
}

// Method to set Theta
func (e *Element1D) SetTheta(nodes []Node) {
	node1 := findNodeByID(nodes, e.N1)
	node2 := findNodeByID(nodes, e.N2)
	if node1 != nil && node2 != nil {
		deltaX := node2.Coordinate[0] - node1.Coordinate[0]
		deltaY := node2.Coordinate[1] - node1.Coordinate[1]
		deltaZ := node2.Coordinate[2] - node1.Coordinate[2]
		e.Theta[0] = math.Atan2(deltaY, deltaX)
		e.Theta[1] = math.Atan2(deltaZ, deltaX)
		e.Theta[2] = math.Atan2(deltaZ, deltaY)
	}
}

// Method to set Local StiffnessMatrix
func (e *Element1D) SetKL(DoF int){
	var k1, C1, C2, C3, C4, C5, C6, C7, C8, C9, C10 float64
	switch {
	case DoF == 1: //Dof u
		if e.Type == "SPRING" {
			k1 = e.K
		}
		if e.Type == "BAR" {
			k1 = (e.E * e.A) / e.L
		}
		stiffnessMatrix := mat.NewDense(2,2, nil)
		stiffnessMatrix.Set(0, 0, k1)
		stiffnessMatrix.Set(0, 1, -k1)
		stiffnessMatrix.Set(1, 0, -k1)
		stiffnessMatrix.Set(1, 1, k1)
		e.KL = stiffnessMatrix
		case DoF == 2: //Dof u,v
		if e.Type == "SPRING" {
			k1 = e.K
		}
		if e.Type == "BAR" {
			k1 = (e.E * e.A) / e.L
		}
		stiffnessMatrix := mat.NewDense(4,4, nil)
		stiffnessMatrix.Set(0, 0, k1)
		stiffnessMatrix.Set(0, 2, -k1)
		stiffnessMatrix.Set(2, 0, -k1)
		stiffnessMatrix.Set(2, 2, k1)
		e.KL = stiffnessMatrix
	case DoF == 3: //Dof u,v,phi
		if e.Type == "SPRING"{
			C1 = e.K
			C2 = 0
		}
		if e.Type == "BAR"{
			C1 = (e.E * e.A) / e.L
			C2 = 0
		}
		if e.Type == "BEAM"{
			C1 = (e.E * e.A) / e.L
			C2 = (e.E * e.Iz) / math.Pow(e.L,3)
		}
		stiffnessMatrix := mat.NewDense(6,6, nil)
		stiffnessMatrix.Set(0, 0, C1)
		stiffnessMatrix.Set(0, 3, -C1)
		stiffnessMatrix.Set(3, 0, -C1)
		stiffnessMatrix.Set(3, 3, C1)

		stiffnessMatrix.Set(1, 1, 12*C2)
		stiffnessMatrix.Set(1, 4, -12*C2)
		stiffnessMatrix.Set(4, 1, -12*C2)
		stiffnessMatrix.Set(4, 4, 12*C2)

		stiffnessMatrix.Set(1, 2, 6*C2*e.L)
		stiffnessMatrix.Set(1, 5, 6*C2*e.L)
		stiffnessMatrix.Set(4, 2, -6*C2*e.L)
		stiffnessMatrix.Set(4, 5, -6*C2*e.L)

		stiffnessMatrix.Set(2, 1, 6*C2*e.L)
		stiffnessMatrix.Set(2, 4, -6*C2*e.L)
		stiffnessMatrix.Set(5, 1, 6*C2*e.L)
		stiffnessMatrix.Set(5, 4, -6*C2*e.L)

		stiffnessMatrix.Set(2, 2, 4*C2*math.Pow(e.L,2))
		stiffnessMatrix.Set(2, 5, 2*C2*math.Pow(e.L,2))
		stiffnessMatrix.Set(5, 2, 2*C2*math.Pow(e.L,2))
		stiffnessMatrix.Set(5, 5, 4*C2*math.Pow(e.L,2))

		e.KL = stiffnessMatrix
	case DoF == 6: //Dof u,v,w,phi1,phi2,phi3
		if e.Type == "SPRING"{
			C1 = e.K
			C2 = 0
		}
		if e.Type == "BAR"{
			C1 = (e.E * e.A) / e.L
			C2 = 0
		}
		if e.Type == "BEAM"{
			C1 = (e.E * e.A) / e.L
			C2 = (e.E * e.Iz) / math.Pow(e.L,3)
			C3 = (e.E * e.Iz) / math.Pow(e.L,2)
			C4 = (e.E * e.Iy) / math.Pow(e.L,3)
			C5 = (e.E * e.Iy) / math.Pow(e.L,2)
			C6 = (e.G * e.Ix) / e.L
			C7 = (e.E * e.Iy) / math.Pow(e.L,2)
			C8 = (e.G * e.Iy) / e.L
			C9 = (e.E * e.Iz) / math.Pow(e.L,2)
			C10 = (e.E * e.Iz) / e.L
		}
		stiffnessMatrix := mat.NewDense(12,12, nil)
		stiffnessMatrix.Set(0, 0, C1)
		stiffnessMatrix.Set(0, 6, -C1)
		stiffnessMatrix.Set(6, 0, -C1)
		stiffnessMatrix.Set(6, 6, C1)

		stiffnessMatrix.Set(1, 1, 12*C2)
		stiffnessMatrix.Set(1, 7, -12*C2)
		stiffnessMatrix.Set(7, 1, -12*C2)
		stiffnessMatrix.Set(7, 7, 12*C2)

		stiffnessMatrix.Set(1, 5, 6*C3)
		stiffnessMatrix.Set(1, 11, 6*C3)
		stiffnessMatrix.Set(11, 5, -6*C3)
		stiffnessMatrix.Set(11, 11, -6*C3)

		stiffnessMatrix.Set(2, 2, 12*C4)
		stiffnessMatrix.Set(2, 8, -12*C4)
		stiffnessMatrix.Set(8, 2, -12*C4)
		stiffnessMatrix.Set(8, 8, 12*C4)

		stiffnessMatrix.Set(2, 4, -6*C5)
		stiffnessMatrix.Set(2, 10, -6*C5)
		stiffnessMatrix.Set(10, 4, 6*C5)
		stiffnessMatrix.Set(10, 10, 6*C5)

		stiffnessMatrix.Set(3, 3, C6)
		stiffnessMatrix.Set(3, 9, C6)
		stiffnessMatrix.Set(9, 3, C6)
		stiffnessMatrix.Set(9, 9, C6)

		stiffnessMatrix.Set(4, 2, -6*C7)
		stiffnessMatrix.Set(4, 8, 6*C7)
		stiffnessMatrix.Set(8, 4, -6*C7)
		stiffnessMatrix.Set(8, 8, 6*C7)

		stiffnessMatrix.Set(4, 4, 4*C8)
		stiffnessMatrix.Set(4, 10, 2*C8)
		stiffnessMatrix.Set(10, 4, 2*C8)
		stiffnessMatrix.Set(10, 10, 4*C8)

		stiffnessMatrix.Set(5, 2, 6*C9)
		stiffnessMatrix.Set(5, 8, -6*C9)
		stiffnessMatrix.Set(8, 5, -6*C9)
		stiffnessMatrix.Set(8, 8, C9)

		stiffnessMatrix.Set(5, 5, 4*C10)
		stiffnessMatrix.Set(5, 11, 2*C10)
		stiffnessMatrix.Set(11, 5, 2*C10)
		stiffnessMatrix.Set(11, 11, 4*C10)

		e.KL = stiffnessMatrix
	}
	
}

/****************************/
/*	Start Analisys methods  */
/****************************/

// Method to count and set the distinct nodes
func (s *Analysis) CountDistinctNodes() {
	nodeMap := make(map[int]bool)
	for _, element := range s.InputData.Element1D {
		nodeMap[element.N1] = true
		nodeMap[element.N2] = true
	} //TODO: Extend calc
	// for _, element := range s.InputData.Element2D {
	// 	nodeMap[element.N1] = true
	// 	nodeMap[element.N2] = true
	// }
	// for _, element := range s.InputData.Element2D {
	// 	nodeMap[element.N1] = true
	// 	nodeMap[element.N2] = true
	// }
	s.CalcData.NumNode = len(nodeMap)
}

// Method to count and set the elements
func (s *Analysis) CountElements() {
	s.CalcData.NumElement = len(s.InputData.Element1D) + len(s.InputData.Element2D) + len(s.InputData.Element3D)
}

// Method to get the constraints as a map[int]map[int]bool
func (s *Analysis) SetConstraints() {
	constraints := make(map[int]map[int]bool)
	for i := 1; i <= s.CalcData.NumNode; i++ {
		constraints[i] = make(map[int]bool)
		constraints[i][0] = false
		constraints[i][1] = false 
		constraints[i][2] = false 
	}

	for _, element := range s.InputData.Element1D {
		node1, _ := getNodeByID(s.InputData.Node, element.N1)
		node2, _ := getNodeByID(s.InputData.Node, element.N1)
		setMapConstrain(constraints, node1, node2, s.InputData.DoF)
	}
	
	
	s.CalcData.Constraints = constraints
}

func setMapConstrain(constraints map[int]map[int]bool, node1 Node, node2 Node, DoF int) {
	
	constraints[node1.Id][0] = node1.Constrains[0]
	constraints[node2.Id][0] = node2.Constrains[0]
	switch {
	case DoF == 2:
		constraints[node1.Id][1] = node1.Constrains[1]
		constraints[node2.Id][1] = node2.Constrains[1]
	case DoF == 3:
		constraints[node1.Id][1] = node1.Constrains[1]
		constraints[node2.Id][1] = node2.Constrains[1]
		constraints[node1.Id][2] = node1.Constrains[3]
		constraints[node2.Id][2] = node2.Constrains[3]
	case DoF == 6:
		constraints[node1.Id][1] = node1.Constrains[1]
		constraints[node2.Id][1] = node2.Constrains[1]
		constraints[node1.Id][2] = node1.Constrains[2]
		constraints[node2.Id][2] = node2.Constrains[2]
		constraints[node1.Id][3] = node1.Constrains[3]
		constraints[node2.Id][3] = node2.Constrains[3]
		constraints[node1.Id][4] = node1.Constrains[4]
		constraints[node2.Id][4] = node2.Constrains[4]
		constraints[node1.Id][5] = node1.Constrains[5]
		constraints[node2.Id][5] = node2.Constrains[5]
	}	
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