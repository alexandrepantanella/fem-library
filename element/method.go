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