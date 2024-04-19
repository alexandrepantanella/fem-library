package handler

import (
	"fmt"
	"net/http"
)

func WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Benvenuto al mio server!")
}
