package main

import (
	"log"
	"net/http"
	"backend-bakashi-go/handlers"
)

func main() {
	// Definir as rotas diretamente no main
	http.HandleFunc("/api/animes", handlers.GetAnimes)
	http.HandleFunc("/api/episodes", handlers.GetEpisodes)

	// Configurar o servidor na porta padrão da Vercel
	log.Println("Servidor rodando na Vercel...")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}
