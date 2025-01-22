package handler

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"backend-bakashi-go/api" // Importa o pacote api com os manipuladores
)

// Função que Vercel espera exportada para rodar o servidor
func Handler(w http.ResponseWriter, r *http.Request) {
	// Definindo os handlers para os endpoints
	http.HandleFunc("/api/animes", api.AnimesHandler)
	http.HandleFunc("/api/episodes", api.EpisodesHandler)

	// Verifica a porta para o Vercel, caso contrário, usa a padrão 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Porta padrão para o Vercel
	}

	// Mensagem de log para iniciar o servidor
	fmt.Printf("Servidor rodando em http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
