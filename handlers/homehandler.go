package handlers

import (
	"fmt"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Bienvenue sur la liste de courses</h1><p>Utilisez /shopping pour voir et modifier votre liste de courses.</p>")
}
