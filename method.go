package main

import (
	"fmt"

	"github.com/fem-library/element"
)

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
		setMapConstrain(constraints, node1, node2)
	}
	
	
	s.CalcData.Constraints = constraints
}

func setMapConstrain(constraints map[int]map[int]bool, element interface{}) {
	constraints[node1.Id][0] = node1.Constrains[0]
	constraints[node1.Id][1] = node1.Constrains[1]
	constraints[node1.Id][2] = node1.Constrains[2]
	constraints[node1.Id][3] = node1.Constrains[3]
	constraints[node1.Id][4] = node1.Constrains[4]
	constraints[node1.Id][5] = node1.Constrains[5]
	constraints[node2.Id][0] = node2.Constrains[0]
	constraints[node2.Id][1] = node2.Constrains[1]
	constraints[node2.Id][2] = node2.Constrains[2]
	constraints[node2.Id][3] = node2.Constrains[3]
	constraints[node2.Id][4] = node2.Constrains[4]
	constraints[node2.Id][5] = node2.Constrains[5]
}

// Auxiliary Function to get a node by its ID
func getNodeByID(nodes []element.Node, id int) (element.Node, error) {
	for _, node := range nodes {
		if node.Id == id {
			return node, nil
		}
	}
	return element.Node{}, fmt.Errorf("node with ID %d not found", id)
}