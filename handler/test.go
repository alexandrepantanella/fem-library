package handler

// import (
// 	"fmt"
// 	"math"
// 	"net/http"

// 	"gonum.org/v1/gonum/mat"

// 	"alexcloud.it/dsm/element"
// 	"alexcloud.it/dsm/material"
// 	"alexcloud.it/dsm/node"
// 	"alexcloud.it/dsm/section"
// 	"alexcloud.it/dsm/utility"
// )

// func Bar1DHandler(w http.ResponseWriter, r *http.Request) {
// 	// Definizione del materiale del beam
// 	steel := material.Material{YoungModulus: 200e9, Density: 7850}

// 	// Definizione delle prop geom del bar
// 	section := section.GeometricProperties{Area: 0.15}

// 	// Definizione dei nodi del bar
// 	node1 := node.Node1D{Coordinate: 1.0}
// 	node2 := node.Node1D{Coordinate: 5.0}

// 	// Creazione del bar con materiale in acciaio e nodi specifici
// 	bar := element.Bar1D{ID: 1, Material: steel, Node1: node1, Node2: node2, Section: section}

// 	// Calcolo della lunghezza del bar
// 	length := bar.Length()

// 	// Calcolo della matrice di rigidezza del bar
// 	stiffnessMatrixLocal := bar.StiffnessMatrix()

// 	// Calcolo della matrice di rigidezza globale del bar
// 	stiffnessMatrixGlobal := bar.GlobalStiffnessMatrix()

// 	// Calcolo della matrice delle masse del bar
// 	massMatrix := bar.MassMatrix()

// 	//Stampa a schermo
// 	fmt.Fprintf(w, "Lunghezza dell'asta: %.2f m\n", length)
// 	fmt.Fprintln(w)
// 	fmt.Fprintf(w, "Matrice di rigidezza dell'asta:")
// 	fmt.Fprintln(w)
// 	utility.PrintMatrixHttp(w, stiffnessMatrixLocal)
// 	fmt.Fprintln(w)
// 	fmt.Fprintf(w, "Matrice delle masse dell'asta:")
// 	fmt.Fprintln(w)
// 	utility.PrintMatrixHttp(w, massMatrix)
// 	fmt.Fprintln(w)
// 	fmt.Fprintf(w, "Matrice di rigidezza globale dell'asta:")
// 	fmt.Fprintln(w)
// 	utility.PrintMatrixHttp(w, stiffnessMatrixGlobal)
// 	fmt.Fprintln(w)

// }

// func Beam1DHandler(w http.ResponseWriter, r *http.Request) {
// 	// Definizione del materiale del beam
// 	steel := material.Material{YoungModulus: 200e9, Density: 7850}

// 	// Definizione delle prop geom del beam
// 	section := section.GeometricProperties{Area: 0.15}

// 	// Definizione dei nodi del beam
// 	node1 := node.Node1D{Coordinate: 1.0}
// 	node2 := node.Node1D{Coordinate: 5.0}

// 	// Creazione del bar con materiale in acciaio e nodi specifici
// 	beam := element.Beam1D{ID: 1, Material: steel, Node1: node1, Node2: node2, Section: section}

// 	// Calcolo della lunghezza del beam
// 	length := beam.Length()

// 	// Calcolo della matrice di rigidezza del beam
// 	stiffnessMatrixLocal := beam.StiffnessMatrix()

// 	// Calcolo della matrice di rigidezza globale del beam
// 	stiffnessMatrixGlobal := beam.GlobalStiffnessMatrix()

// 	// Calcolo della matrice delle masse del bar
// 	massMatrix := beam.MassMatrix()

// 	//Stampa a schermo
// 	fmt.Fprintf(w, "Lunghezza del bar: %.2f m\n", length)
// 	fmt.Fprintln(w)
// 	fmt.Fprintf(w, "Matrice di rigidezza del bar:")
// 	fmt.Fprintln(w)
// 	utility.PrintMatrixHttp(w, stiffnessMatrixLocal)
// 	fmt.Fprintln(w)
// 	fmt.Fprintf(w, "Matrice delle masse del bar:")
// 	fmt.Fprintln(w)
// 	utility.PrintMatrixHttp(w, massMatrix)
// 	fmt.Fprintln(w)
// 	fmt.Fprintf(w, "Matrice di rigidezza globale del bar:")
// 	fmt.Fprintln(w)
// 	utility.PrintMatrixHttp(w, stiffnessMatrixGlobal)
// 	fmt.Fprintln(w)

// }

// func Beam2DHandler(w http.ResponseWriter, r *http.Request) {
// 	// Definizione del materiale del beam
// 	steel := material.Material{YoungModulus: 200e9, Density: 7850}

// 	// Definizione dei nodi del beam
// 	// Definizione dei nodi del beam
// 	node1 := node.Node2D{ID: 1, Coordinates: [2]float64{1.0, 2.0}}
// 	node2 := node.Node2D{ID: 2, Coordinates: [2]float64{3.0, 4.0}}

// 	// Creazione del beam con materiale in acciaio e nodi specifici
// 	beam := element.Beam2D{Material: steel, Node1: node1, Node2: node2}

// 	// Calcolo della lunghezza del beam
// 	length := beam.Length()

// 	// Calcolo della matrice di rigidità del beam
// 	stiffnessMatrix := beam.StiffnessMatrix()

// 	// Calcolo della matrice delle masse del beam
// 	massMatrix := beam.MassMatrix()

// 	//Stampa a schermo
// 	fmt.Fprintf(w, "Lunghezza del beam: %.2f m\n", length)
// 	fmt.Fprintln(w)
// 	fmt.Fprintf(w, "Matrice di rigidezza del beam:")
// 	fmt.Fprintln(w)
// 	utility.PrintMatrixHttp(w, stiffnessMatrix)
// 	fmt.Fprintln(w)
// 	fmt.Fprintf(w, "Matrice delle masse del beam:")
// 	fmt.Fprintln(w)
// 	utility.PrintMatrixHttp(w, massMatrix)
// 	fmt.Fprintln(w)

// }

// func TestTransformMatrix3D(w http.ResponseWriter, r *http.Request) {
// 	thetaX := math.Pi / 2
// 	thetaY := math.Pi
// 	thetaZ := math.Pi / 4
// 	// Costruiamo la matrice di rotazione
// 	rotX := mat.NewDense(3, 3, []float64{
// 		1, 0, 0,
// 		0, math.Cos(thetaX), -math.Sin(thetaX),
// 		0, math.Sin(thetaX), math.Cos(thetaX),
// 	})
// 	rotY := mat.NewDense(3, 3, []float64{
// 		math.Cos(thetaY), 0, math.Sin(thetaY),
// 		0, 1, 0,
// 		-math.Sin(thetaY), 0, math.Cos(thetaY),
// 	})
// 	rotZ := mat.NewDense(3, 3, []float64{
// 		math.Cos(thetaZ), -math.Sin(thetaZ), 0,
// 		math.Sin(thetaZ), math.Cos(thetaZ), 0,
// 		0, 0, 1,
// 	})

// 	// Effettuiamo il prodotto delle matrici di rotazione
// 	transformMatrix := mat.NewDense(3, 3, nil)
// 	transformMatrix.Mul(rotZ, rotY)
// 	transformMatrix.Mul(transformMatrix, rotX)

// 	// Estendiamo la matrice di rotazione a 12x12
// 	transformMatrixExtended := mat.NewDense(12, 12, nil)
// 	for i := 0; i < 3; i++ {
// 		for j := 0; j < 3; j++ {
// 			transformMatrixExtended.Set(i*4+j, i*4+j, transformMatrix.At(i, j))
// 		}
// 	}

// 	// Stampiamo la matrice di trasformazione
// 	fmt.Println("Matrice di trasformazione:")
// 	fmt.Println(mat.Formatted(transformMatrixExtended))

// }
