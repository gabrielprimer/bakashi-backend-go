package handler

import (
	
	"net/http"

	"backend-bakashi-go/api" // Importa os manipuladores de API
)

// Exportando um handler para o Vercel
func Handler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/api/animes":
		api.AnimesHandler(w, r)
	case "/api/episodes":
		api.EpisodesHandler(w, r)
	default:
		http.NotFound(w, r)
	}
}
