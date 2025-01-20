package main

import (
    "github.com/rs/cors"
    "log"
    "net/http"
    "backend-bakashi-go/handlers"
)

func main() {
    mux := http.NewServeMux()

    // Carregar dados de animes e episódios
    mux.HandleFunc("/api/animes", handlers.GetAnimes)
    mux.HandleFunc("/api/episodes", handlers.GetEpisodes)

    // Configurar CORS
    c := cors.New(cors.Options{
        AllowedOrigins: []string{"*"}, // Permitir todas as origens
    })

    handler := c.Handler(mux)

    log.Println("Servidor rodando em http://localhost:5000")
    log.Fatal(http.ListenAndServe(":5000", handler))
}
