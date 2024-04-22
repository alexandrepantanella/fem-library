package force

// Force1D rappresenta le forze applicate a un elemento monodimensionale
type Force1D struct {
	X1, X2	float64		// Forza assegnata lungo x sui nodi	
}

// Force2D rappresenta le forze applicate a un elemento bidimensionale
type Force2D struct {
	X, Y float64 // Forze nelle direzioni x,y
	Mz   float64 //Momento
}

// Force2D rappresenta le forze applicate a un elemento tridimensionale
type Force3D struct {
	X, Y, Z    float64 // Forze nelle direzioni x,y,z
	Mx, My, Mz float64 // Momenti sui tre assi
}