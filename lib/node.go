package lib

type Node1 struct {
	Idx int
	X   float64 //X coordinate
	Ux  int     //X Constrain (Translation)
	Fx  float64 //Nodal Force applied
}

type Node2 struct {
	Idx int
	X   float64 //X coordinate
	Y   float64 //Y coordinate
	Ux  int     //X Constrain (Translation)
	Uy  int     //Y Constrain (Translation)
	Rz  int     //Rz Constrain  (Rotation)
	Fx  float64 //Nodal Force applied
	Fy  float64 //Nodal Force applied
	Mz  float64 //Nodal Moment applied
}

type Node3 struct {
	Idx int
	X   float64 //X coordinate
	Y   float64 //Y coordinate
	Z   float64 //Z coordinate
	Ux  int     //X Constrain (Translation)
	Uy  int     //Y Constrain (Translation)
	Uz  int     //Z Constrain (Translation)
	Rx  int     //Rx Constrain  (Rotation)
	Ry  int     //Ry Constrain  (Rotation)
	Rz  int     //Rz Constrain  (Rotation)
	Fx  float64 //Nodal Force applied
	Fy  float64 //Nodal Force applied
	Fz  float64 //Nodal Force applied
	Mx  float64 //Nodal Moment applied
	My  float64 //Nodal Moment applied
	Mz  float64 //Nodal Moment applied
}
