package lib

import (
	"errors"
)


func CheckStructure(a Analysis) error {
	err := checkNodeIDs(a.InputData.Node)
	if err != nil {
		return err
	}
	err = checkElementsPresent(a)
	if err != nil {
		return err
	}
	err = checkNodeIDsAssociated(a.InputData.Node, a.InputData.Element1D)
	if err != nil {
		return err
	}
	return nil
}


// Check if all Node IDs are consecutive and in order
func checkNodeIDs(nodes []Node) error {
	// Create a map to keep track of node IDs
	nodeIDs := make(map[int]bool)

	// Check node IDs
	for _, node := range nodes {
		// Check if node ID already exists in the map
		if nodeIDs[node.Id] {
			// If node ID already exists, return an error
			return errors.New("not all node IDs are unique")
		}
		// Add node ID to the map
		nodeIDs[node.Id] = true
	}

	// Check if node IDs are consecutive and in order
	for i := 1; i <= len(nodes); i++ {
		// Check if node ID exists in the map
		if !nodeIDs[i] {
			// If node ID is missing, return an error
			return errors.New("node IDs are not consecutive or in order")
		}
	}

	// If all node IDs are consecutive and in order, return nil (no error)
	return nil
}

// Check if at least one Element is present
func checkElementsPresent(a Analysis) error {
	// Check if there are elements in the Analysis structure
	if len(a.InputData.Element1D) == 0 && len(a.InputData.Element2D) == 0 && len(a.InputData.Element3D) == 0 {
		// If there are no elements, return an error
		return errors.New("no element present")
	}

	// If at least one type of element is present, return nil (no error)
	return nil
}

// Check that no Node ID is not associated with any element
//func checkNodeIDsAssociated(nodes []Node, elements []Element1D, elements2D []Element2D, elements3D []Element3D) error {
func checkNodeIDsAssociated(nodes []Node, elements []Element1D) error {
	// Create a map to keep track of node IDs associated with elements
	nodeIDsAssociated := make(map[int]bool)

	// Check Node IDs associated with Element1D
	for _, element := range elements {
		nodeIDsAssociated[element.N1] = true
		nodeIDsAssociated[element.N2] = true
	}

	// Check Node IDs associated with Element2D
	// for _, element := range elements2D {
	// 	// Add more node IDs association logic here if needed
	// }

	// Check Node IDs associated with Element3D
	// for _, element := range elements3D {
	// 	// Add more node IDs association logic here if needed
	// }

	// Check Node IDs not associated with any element
	for _, node := range nodes {
		// If Node ID is not associated with any element, return an error
		if !nodeIDsAssociated[node.Id] {
			return errors.New("Node ID not associated with any element")
		}
	}

	// If all Node IDs are associated with at least one element, return nil (no error)
	return nil
}