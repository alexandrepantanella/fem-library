package main

import "alexcloud.it/dsm/test"

func main() {
	// Imposta il gestore per la rotta di benvenuto
	// http.HandleFunc("/", handler.WelcomeHandler)
	// http.HandleFunc("/testBar1D", handler.Bar1DHandler)
	// http.HandleFunc("/testBeam1D", handler.Beam1DHandler)
	// http.HandleFunc("/testBeam2D", handler.Beam2DHandler)
	// http.HandleFunc("/TestTransformMatrix3D", handler.TestTransformMatrix3D)

	//Avvia il server HTTP sulla porta 8082
	// go func() {
	// 	log.Fatal(http.ListenAndServe(":8082", handler.LogHandler(http.DefaultServeMux)))
	// }()

	// Attendi indefinitamente
	// select {}
	test.TestBeam3DEulerStiffnessMatrix()
}
