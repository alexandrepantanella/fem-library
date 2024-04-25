package element

import (
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
	var k1, k2, k3, k4 float64
	switch {
	case DoF == 2:
		if e.Type == "SPRING" {
			k1 = e.K
			k2 = -e.K
			k3 = -e.K
			k4 = e.K
		}
		if e.Type == "BAR" || e.Type == "BEAM" {
			k1 = (e.E * e.A) / e.L
			k2 = -k1
			k3 = -k1
			k4 = k1
		}
		stiffnessMatrix := mat.NewDense(2,2, nil)
		stiffnessMatrix.Set(0, 0, k1)
		stiffnessMatrix.Set(0, 3, k2)
		stiffnessMatrix.Set(3, 0, k3)
		stiffnessMatrix.Set(3, 3, k4)
		e.KL = stiffnessMatrix
	case DoF == 3:
		if e.Type == "BEAM"{
			C1 := (e.E * e.A) / e.L
			C2 := (e.E * e.Iz) / math.Pow(e.L,3)
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
		}
	}
	
}