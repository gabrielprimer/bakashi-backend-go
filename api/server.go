package main

import (
	"log"
	"os" // 
	"net/http"
	"github.com/rs/cors"
	"backend-bakashi-go/handlers"  // Corrigido para importar o pacote handlers
)

func main() {
	// Configurar as rotas
	mux := http.NewServeMux()
	mux.HandleFunc("/api/animes", handlers.GetAnimes)   // Rota para animes
	mux.HandleFunc("/api/episodes", handlers.GetEpisodes)  // Rota para episódios

	// Usando o middleware CORS
	handler := cors.Default().Handler(mux)

	// Pega a porta do ambiente (se existir) ou usa a 3000 como fallback
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000" // Default
	}

	log.Println("Servidor rodando na porta", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
