package transform

import (
	"math"

	"gonum.org/v1/gonum/mat"
)

// TransformMatrix2D restituisce la matrice di trasformazione locale-globale per un dato angolo theta in radianti in 2D
func TransformMatrix2D(theta float64) *mat.Dense {
	cosTheta := math.Cos(theta)
	sinTheta := math.Sin(theta)

	// Inizializza la matrice con dimensione 6x6
	transformMatrix := mat.NewDense(6, 6, nil)

	// Imposta i valori sulla matrice di trasformazione
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
}

// TransformMatrix3D restituisce la matrice di trasformazione locale-globale per tre angoli di rotazione in radianti in 3D
func TransformMatrix3D(thetaX, thetaY, thetaZ float64) *mat.Dense {
	cosX, sinX := math.Cos(thetaX), math.Sin(thetaX)
	cosY, sinY := math.Cos(thetaY), math.Sin(thetaY)
	cosZ, sinZ := math.Cos(thetaZ), math.Sin(thetaZ)

	// Inizializza la matrice con dimensione 12x12
	transformMatrix := mat.NewDense(12, 12, nil)

	// Imposta i valori sulla matrice di trasformazione
	transformMatrix.Set(0, 0, cosY*cosZ)
	transformMatrix.Set(0, 1, cosY*sinZ)
	transformMatrix.Set(0, 2, -sinY)
	transformMatrix.Set(3, 3, cosY*cosZ)
	transformMatrix.Set(3, 4, cosY*sinZ)
	transformMatrix.Set(3, 5, -sinY)
	transformMatrix.Set(6, 6, cosY*cosZ)
	transformMatrix.Set(6, 7, cosY*sinZ)
	transformMatrix.Set(6, 8, -sinY)
	transformMatrix.Set(9, 9, cosY*cosZ)
	transformMatrix.Set(9, 10, cosY*sinZ)
	transformMatrix.Set(9, 11, -sinY)

	transformMatrix.Set(1, 0, sinX*sinY*cosZ-cosX*sinZ)
	transformMatrix.Set(1, 1, sinX*sinY*sinZ+cosX*cosZ)
	transformMatrix.Set(1, 2, sinX*cosY)
	transformMatrix.Set(4, 3, sinX*sinY*cosZ-cosX*sinZ)
	transformMatrix.Set(4, 4, sinX*sinY*sinZ+cosX*cosZ)
	transformMatrix.Set(4, 5, sinX*cosY)
	transformMatrix.Set(7, 6, sinX*sinY*cosZ-cosX*sinZ)
	transformMatrix.Set(7, 7, sinX*sinY*sinZ+cosX*cosZ)
	transformMatrix.Set(7, 8, sinX*cosY)
	transformMatrix.Set(10, 9, sinX*sinY*cosZ-cosX*sinZ)
	transformMatrix.Set(10, 10, sinX*sinY*sinZ+cosX*cosZ)
	transformMatrix.Set(10, 11, sinX*cosY)

	transformMatrix.Set(2, 3, cosX*sinY*cosZ+sinX*sinZ)
	transformMatrix.Set(2, 4, cosX*sinY*sinZ-sinX*cosZ)
	transformMatrix.Set(2, 5, cosX*cosY)
	transformMatrix.Set(5, 6, cosX*sinY*cosZ+sinX*sinZ)
	transformMatrix.Set(5, 7, cosX*sinY*sinZ-sinX*cosZ)
	transformMatrix.Set(5, 8, cosX*cosY)
	transformMatrix.Set(8, 9, cosX*sinY*cosZ+sinX*sinZ)
	transformMatrix.Set(8, 10, cosX*sinY*sinZ-sinX*cosZ)
	transformMatrix.Set(8, 11, cosX*cosY)
	transformMatrix.Set(11, 9, cosX*sinY*cosZ+sinX*sinZ)
	transformMatrix.Set(11, 10, cosX*sinY*sinZ-sinX*cosZ)
	transformMatrix.Set(11, 11, cosX*cosY)

	return transformMatrix
}
