package analisys

import (
	"github.com/fem-library/element"
	"github.com/fem-library/node"
)

type Analisys1D struct {
	Dim       int    //Dimensioni del modello
	Type      string // SL or SNL or DL
	Behaviour string // ISO or ORTHO (material)
	Bar1D     []element.Bar1D
	Spring1D  []element.Spring1D
	Beam1D    []element.Beam1D
	Node1D    []node.Node1D
	//Output	  []Output
}
