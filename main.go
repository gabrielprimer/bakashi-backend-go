package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"backend-bakashi-go/api"  // Importa apenas o pacote "api"
)

func main() {
	// Definir os handlers para os endpoints
	http.HandleFunc("/api/animes", api.AnimesHandler)
	http.HandleFunc("/api/episodes", api.EpisodesHandler)

	// Vercel espera que o servidor escute na porta fornecida por uma variável de ambiente
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"  // Porta padrão para o Vercel
	}

	// Mensagem de log para iniciar o servidor
	fmt.Printf("Servidor rodando em http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
