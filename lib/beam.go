package lib

import "math"

type Beam1 struct {
	Id    int
	Node1 Node1
	Node2 Node1
	A     float64 //Section area
	Fu1   float64 //Force applied in 1
	Fu2   float64 //Force applied in 2
	Qu    float64 //Distributed Force applied along u
	E     float64 //Modulus E
	Dt    float64 //Delta T applied

}

type Beam2 struct {
	Id    int
	Node1 Node2
	Node2 Node2
	A     float64 //Section area
	E     float64 //Modulus E
	G     float64 //Modulus G
	Fu1   float64 //Force applied in 1 along u
	Fv1   float64 //Force applied in 1 along v
	Fu2   float64 //Force applied in 2 along u
	Fv2   float64 //Force applied in 2 along v
	Qu    float64 //Distributed Force applied along u
	Qv    float64 //Distributed Force applied along v
	Dt    float64 //Delta T applied
}

type Beam3 struct {
	Id    int
	Node1 Node3
	Node2 Node3
	A     float64 //Section area
	E1    float64 //Modulus E along u
	E2    float64 //Modulus E along v
	E3    float64 //Modulus E along z
	G1    float64 //Modulus G along u
	G2    float64 //Modulus G along v
	G3    float64 //Modulus G along z
	Fu1   float64 //Force applied in 1 along u
	Fv1   float64 //Force applied in 1 along v
	Fz1   float64 //Force applied in 1 along z
	Fu2   float64 //Force applied in 2 along u
	Fv2   float64 //Force applied in 2 along u
	Fz2   float64 //Force applied in 2 along z
	Qu    float64 //Distributed Force applied along u
	Qv    float64 //Distributed Force applied along v
	Dt    float64 //Delta T applied
}

// Length of the Beam1
func (b *Beam1) L() float64 {
	return math.Abs(b.Node1.X - b.Node2.X)
}

// Length of the Beam2
func (b *Beam2) L() float64 {
	return math.Sqrt(math.Pow(b.Node1.X-b.Node2.X, 2) + math.Pow(b.Node1.Y-b.Node2.Y, 2))
}

// Length of the Beam3
func (b *Beam3) L() float64 {
	return math.Sqrt(math.Pow(b.Node1.X-b.Node2.X, 2) + math.Pow(b.Node1.Y-b.Node2.Y, 2) + math.Pow(b.Node1.Z-b.Node2.Z, 2))
}
