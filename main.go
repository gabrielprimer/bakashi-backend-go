package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"backend-bakashi-go/api" // Importe o pacote api
)

// Função principal de handler para o Vercel
func Handler(w http.ResponseWriter, r *http.Request) {
	// Defina os handlers para os endpoints
	http.HandleFunc("/api/animes", api.AnimesHandler)
	http.HandleFunc("/api/episodes", api.EpisodesHandler)

	// Verifica a porta para o Vercel, caso contrário, usa a padrão 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Porta padrão para o Vercel
	}

	// Inicia o servidor
	fmt.Printf("Servidor rodando em http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func main() {
	// Apenas chamando o Handler para a Vercel
	http.HandleFunc("/", Handler)
}
