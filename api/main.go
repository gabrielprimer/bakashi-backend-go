package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"backend-bakashi-go/types"
)


// Estrutura para representar a lista de animes
type Animes struct {
	Animes []types.Anime `json:"animes"`
}

// Estrutura para representar a lista de episódios
type Episodes struct {
	Episodes []types.Episode `json:"episodes"`
}

func main() {
	// Abrir o arquivo de dados JSON para animes
	file, err := os.Open("./data/animes.json")
	if err != nil {
		log.Fatalf("Erro ao abrir animes.json: %v", err)
	}
	defer file.Close()

	// Decodificar o JSON para a estrutura Animes
	var animes Animes
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&animes)
	if err != nil {
		log.Fatalf("Erro ao decodificar animes.json: %v", err)
	}

	// Abrir o arquivo de dados JSON para episódios
	episodeFile, err := os.Open("./data/episodes.json")
	if err != nil {
		log.Fatalf("Erro ao abrir episodes.json: %v", err)
	}
	defer episodeFile.Close()

	// Decodificar o JSON para a estrutura Episodes
	var episodes Episodes
	episodeDecoder := json.NewDecoder(episodeFile)
	err = episodeDecoder.Decode(&episodes)
	if err != nil {
		log.Fatalf("Erro ao decodificar episodes.json: %v", err)
	}

	// Rota para retornar os animes
	http.HandleFunc("/api/animes", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(animes)
	})

	// Rota para retornar os episódios
	http.HandleFunc("/api/episodes", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(episodes)
	})

	// Iniciar o servidor na porta 5000
	fmt.Println("Servidor rodando em http://localhost:5000")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
