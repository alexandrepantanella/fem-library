package section

type GeometricProperties struct {
	Area      float64 //Area della sezione
	Length    float64 // Lunghezza dell'elemento
	Width     float64 // Larghezza dell'elemento
	Height    float64 // Altezza dell'elemento
	Thickness float64 // Spessore dell'elemento
	Iy        float64
	Iz        float64
	J         float64
}
