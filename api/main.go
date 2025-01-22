package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"backend-bakashi-go/api/types"
)

// Estrutura para representar a lista de animes
type Animes struct {
	Animes []types.Anime `json:"animes"`
}

// Estrutura para representar a lista de episódios
type Episodes struct {
	Episodes []types.Episode `json:"episodes"`
}

// Handler exportada que o Vercel irá reconhecer
func Handler(w http.ResponseWriter, r *http.Request) {
	// Abrir o arquivo de dados JSON para animes
	file, err := os.Open("./data/animes.json")
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao abrir animes.json: %v", err), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Decodificar o JSON para a estrutura Animes
	var animes Animes
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&animes)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao decodificar animes.json: %v", err), http.StatusInternalServerError)
		return
	}

	// Abrir o arquivo de dados JSON para episódios
	episodeFile, err := os.Open("./data/episodes.json")
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao abrir episodes.json: %v", err), http.StatusInternalServerError)
		return
	}
	defer episodeFile.Close()

	// Decodificar o JSON para a estrutura Episodes
	var episodes Episodes
	episodeDecoder := json.NewDecoder(episodeFile)
	err = episodeDecoder.Decode(&episodes)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao decodificar episodes.json: %v", err), http.StatusInternalServerError)
		return
	}

	// Definir a rota para retornar os animes
	http.HandleFunc("/api/animes", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(animes)
	})

	// Definir a rota para retornar os episódios
	http.HandleFunc("/api/episodes", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(episodes)
	})

	// Mensagem de log informando que o servidor foi iniciado
	fmt.Println("Servidor rodando em http://localhost:5000")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
