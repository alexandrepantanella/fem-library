package analisys

import (
	"github.com/fem-library/element"
	"gonum.org/v1/gonum/mat"
)
type Analisys struct {
	Dim       int    //Dimensioni del modello
	Type      string // SL or SNL or DL
	Behaviour string // ISO or ORTHO (material)
	Analisys1D	Analisys1D
	Analisys2D	Analisys2D
	Analisys3D	Analisys3D
	GlobalStiffnessMatrix mat.Dense //Matrice di rigidezza globale
}

type Analisys1D struct {
	Bar1D     []element.Bar1D
	Spring1D  []element.Spring1D
	Beam1D    []element.Beam1D
}
type Analisys2D struct {
	//Bar1D     []element.Bar2D TODO:Add element
	Spring1D  []element.Spring2D
	Beam1D    []element.Beam2D
}

type Analisys3D struct {
	//Bar1D     []element.Bar3D TODO:Add element
	Spring1D  []element.Spring3D
	Beam1D    []element.Beam3D
}
