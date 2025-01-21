package main

import (
	"log"
	"net/http"

	"backend-bakashi-go/api/handlers"
)

func main() {
	// Define as rotas
	http.HandleFunc("/api/animes", handlers.AnimesHandler)
	http.HandleFunc("/api/episodes", handlers.EpisodesHandler)

	// Inicia o servidor
	log.Println("Servidor rodando na porta 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
