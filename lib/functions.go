package lib

import (
	"math"

	"gonum.org/v1/gonum/mat"
)

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

func TransformMatrix(theta float64, DoF int) *mat.Dense {
	
	cosTheta := math.Cos(theta)
	sinTheta := math.Sin(theta)

	switch {
	case DoF == 2:
		transformMatrix := mat.NewDense(4, 4, nil)
		transformMatrix.Set(0, 0, cosTheta)
		transformMatrix.Set(0, 1, -sinTheta)
		transformMatrix.Set(1, 0, sinTheta)
		transformMatrix.Set(1, 1, cosTheta)
		transformMatrix.Set(2, 2, cosTheta)
		transformMatrix.Set(2, 3, -sinTheta)
		transformMatrix.Set(3, 2, sinTheta)
		transformMatrix.Set(3, 3, cosTheta)
		return transformMatrix
	case DoF == 3:
		transformMatrix := mat.NewDense(6, 6, nil)
 		transformMatrix.Set(0, 0, cosTheta)
		transformMatrix.Set(0, 1, -sinTheta)
		transformMatrix.Set(1, 0, sinTheta)
		transformMatrix.Set(1, 1, cosTheta)
		transformMatrix.Set(2, 2, 1)
		transformMatrix.Set(3, 3, cosTheta)
		transformMatrix.Set(3, 4, -sinTheta)
		transformMatrix.Set(4, 3, sinTheta)
		transformMatrix.Set(4, 4, cosTheta)
		transformMatrix.Set(5, 5, 1)
		return transformMatrix
	case DoF == 6:
		transformMatrix := mat.NewDense(12, 12, nil)
		//TODO:
		return transformMatrix
	} 
	transformMatrix := mat.NewDense(1, 1, nil)
	return transformMatrix
}