package lib

import (
	"fmt"
	"math"

	"gonum.org/v1/gonum/mat"
)

/****************************/
/*  Element 1D methods      */
/****************************/

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

//Method to initialize struct
func (e *Element1D) Init(DoF int) {
    e.Theta = make([]float64, DoF) 
    e.KL = mat.NewDense(2*DoF, 2*DoF, nil) 
    e.KG = mat.NewDense(2*DoF, 2*DoF, nil) 
}

// Method to set Length
func (e *Element1D) SetLength(nodes []Node, DoF int) {
	node1 := findNodeByID(nodes, e.N1)
	node2 := findNodeByID(nodes, e.N2)
	if node1 != nil && node2 != nil {
		switch {
			case DoF == 1: //Dof u
				e.L = node2.Coordinate[0] - node1.Coordinate[0]
			case DoF == 2 || DoF == 3: //Dof u,v,phi
				e.L = math.Sqrt(math.Pow(node2.Coordinate[0] - node1.Coordinate[0],2) + 
				math.Pow(node2.Coordinate[1] - node1.Coordinate[1],2))
			case DoF == 6: //Dof u,v,w,phi1,phi2,phi3
				e.L = math.Sqrt(math.Pow(node2.Coordinate[0] - node1.Coordinate[0],2) + 
				math.Pow(node2.Coordinate[1] - node1.Coordinate[1],2) +
				math.Pow(node2.Coordinate[2] - node1.Coordinate[2],2))
		}
	}
}

// Method to set Theta
func (e *Element1D) SetTheta(nodes []Node, DoF int) {
	node1 := findNodeByID(nodes, e.N1)
	node2 := findNodeByID(nodes, e.N2)
	if node1 != nil && node2 != nil {
		e.Theta = make([]float64, DoF)
		switch {
			case DoF == 1: //Dof u
				e.Theta[0] = 0
			case DoF == 2 || DoF == 3: //Dof u,v,phi
				deltaX := node2.Coordinate[0] - node1.Coordinate[0]
				deltaY := node2.Coordinate[1] - node1.Coordinate[1]
				e.Theta[0] = math.Atan2(deltaY, deltaX)
			case DoF == 6: //Dof u,v,w,phi1,phi2,phi3
				deltaX := node2.Coordinate[0] - node1.Coordinate[0]
				deltaY := node2.Coordinate[1] - node1.Coordinate[1]
				deltaZ := node2.Coordinate[2] - node1.Coordinate[2]
				e.Theta[0] = math.Atan2(deltaY, deltaX)
				e.Theta[1] = math.Atan2(deltaZ, deltaX)
				e.Theta[2] = math.Atan2(deltaZ, deltaY)
		}
	}
}

// Method to set Local StiffnessMatrix for one-dimensional element
func (e *Element1D) SetKL(DoF int){
	var stiffnessMatrix *mat.Dense
	var  C1, C2, C3, C4, C5, C6, C7, C8, C9, C10 float64
	switch {
	case DoF == 1: //Dof u
			if e.Type == "SPRING" {
				C1 = e.K
			}
			if e.Type == "BAR" {
				C1 = (e.E * e.A) / e.L
			}
			stiffnessMatrix = mat.NewDense(2*DoF, 2*DoF, []float64{C1, -C1, -C1, C1})
		case DoF == 2: //Dof u,v
			if e.Type == "SPRING" {
				C1 = e.K
			}
			if e.Type == "BAR" {
				C1 = (e.E * e.A) / e.L
			}
			stiffnessMatrix = mat.NewDense(2*DoF, 2*DoF, []float64{C1, 0, -C1, 0, 0, C1, 0, -C1, -C1, 0, C1, 0, 0, -C1, 0, C1})
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
		stiffnessMatrix = mat.NewDense(2*DoF, 2*DoF, []float64{
            C1, 0, 0, -C1, 0, 0,
            0, 12 * C2, 6 * C2 * e.L, 0, -12 * C2, 6 * C2 * e.L,
            0, 6 * C2 * e.L, 4 * C2 * math.Pow(e.L, 2), 0, -6 * C2 * e.L, 2 * C2 * math.Pow(e.L, 2),
            -C1, 0, 0, C1, 0, 0,
            0, -12 * C2, -6 * C2 * e.L, 0, 12 * C2, -6 * C2 * e.L,
            0, 6 * C2 * e.L, 2 * C2 * math.Pow(e.L, 2), 0, -6 * C2 * e.L, 4 * C2 * math.Pow(e.L, 2),
        })
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
		stiffnessMatrix = mat.NewDense(2*DoF, 2*DoF, []float64{
            C1, 0, 0, 0, 0, 0, -C1, 0, 0, 0, 0, 0,
            0, 12 * C2, 0, 0, 0, 6 * C3, 0, -12 * C2, 0, 0, 0, 6 * C3 * e.L,
            0, 0, 12 * C4, 0, -6 * C5, 0, 0, 0, -12 * C4, 0, -6 * C5 * e.L, 0,
            0, 0, 0, C6, 0, 0, 0, 0, 0, C6, 0, 0,
            0, 0, -6 * C7, 0, 4 * C8, 0, 0, 0, 6 * C7, 0, 2 * C8 * e.L, 0,
            0, 6 * C9 * e.L, 0, 0, 0, 2 * C9 * math.Pow(e.L, 2), 0, -6 * C9 * e.L, 0, 0, 4 * C10 * math.Pow(e.L, 2), 0,
            -C1, 0, 0, 0, 0, 0, C1, 0, 0, 0, 0, 0,
            0, -12 * C2, 0, 0, 0, -6 * C3 * e.L, 0, 12 * C2, 0, 0, 0, -6 * C3 * e.L,
            0, 0, -12 * C4, 0, 6 * C5 * e.L, 0, 0, 0, 12 * C4, 0, 6 * C5, 0,
            0, 0, 0, C6, 0, 0, 0, 0, 0, C6, 0, 0,
            0, 0, -6 * C7 * e.L, 0, 2 * C8 * e.L, 0, 0, 0, 6 * C7 * e.L, 0, 4 * C8, 0,
            0, 6 * C9 * e.L, 0, 0, 0, 2 * C9 * math.Pow(e.L, 2), 0, -6 * C9 * e.L, 0, 0, 4 * C10 * math.Pow(e.L, 2), 0,
        })
	}
	e.KL = 	stiffnessMatrix	
}

// Method to set Local StiffnessMatrix for one-dimensional element
func (e *Element1D) SetKG(DoF int){
	switch {
		case DoF == 1: //Dof u
		e.KG = e.KL
	}
}

/****************************/
/*  	CalcData methods    */
/****************************/

func (c *CalcData) Init() {
    c.Length = make(map[int]float64)
    c.Constraints = make(map[int]map[int]bool)
    c.LocalStiffness = make(map[int]*mat.Dense)
    c.GlobalStiffness = make(map[int]*mat.Dense)
    c.Mass = make(map[int]*mat.Dense)
    c.Force = make(map[int]*mat.VecDense)
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
	switch {
		case s.InputData.DoF == 1: //Dof u
		for i := 1; i <= s.CalcData.NumNode; i++ {
			constraints[i] = make(map[int]bool)
			constraints[i][0] = false
		}
		case s.InputData.DoF == 2: //Dof u v
		for i := 1; i <= s.CalcData.NumNode; i++ {
			constraints[i] = make(map[int]bool)
			constraints[i][0] = false
			constraints[i][1] = false 
		}
		case s.InputData.DoF == 3: //Dof u v phi
		for i := 1; i <= s.CalcData.NumNode; i++ {
			constraints[i] = make(map[int]bool)
			constraints[i][0] = false
			constraints[i][1] = false 
			constraints[i][2] = false 
		}
		case s.InputData.DoF == 6: //Dof u v w phi1 phi2 phi3
		for i := 1; i <= s.CalcData.NumNode; i++ {
			constraints[i] = make(map[int]bool)
			constraints[i][0] = false
			constraints[i][1] = false 
			constraints[i][2] = false 
			constraints[i][3] = false
			constraints[i][4] = false 
			constraints[i][5] = false 
		}
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