package handler

import (
	"net/http"
	"backend-bakashi-go/api"
)

// Handler é a função exportada que a Vercel utiliza como ponto de entrada.
func Handler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/api/animes":
		handlers.AnimesHandler(w, r)
	case "/api/episodes":
		handlers.EpisodesHandler(w, r)
	default:
		http.NotFound(w, r)
	}
}
