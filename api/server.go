package main

import (
	"log"
	"net/http"
	"backend-bakashi-go/handlers"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Carregar dados de animes e episódios
	mux := http.NewServeMux()
	mux.HandleFunc("/api/animes", handlers.GetAnimes)
	mux.HandleFunc("/api/episodes", handlers.GetEpisodes)

	// Definir o cabeçalho de resposta
	w.Header().Set("Content-Type", "application/json")

	// Chamar o handler apropriado
	mux.ServeHTTP(w, r)
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("Servidor rodando na Vercel...")
}
