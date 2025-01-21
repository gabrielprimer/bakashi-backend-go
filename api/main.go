package main

import (
	"log"
	"net/http"
	"os"
	"github.com/gorilla/mux"
	"encoding/json"
	"backend-bakashi-go/types"
	"io/ioutil"
)

// Função para carregar os animes
func loadAnimes() ([]types.Anime, error) {
	data, err := ioutil.ReadFile("data/animes.json")
	if err != nil {
		return nil, err
	}
	var animes struct {
		Animes []types.Anime `json:"animes"`
	}
	err = json.Unmarshal(data, &animes)
	if err != nil {
		return nil, err
	}
	return animes.Animes, nil
}

// Handler para listar os animes
func AnimesHandler(w http.ResponseWriter, r *http.Request) {
	animes, err := loadAnimes()
	if err != nil {
		http.Error(w, "Erro ao carregar os animes", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{"animes": animes})
}

// Função para carregar os episódios
func loadEpisodes() ([]types.Episode, error) {
	data, err := ioutil.ReadFile("data/episodes.json")
	if err != nil {
		return nil, err
	}
	var episodes struct {
		Episodes []types.Episode `json:"episodes"`
	}
	err = json.Unmarshal(data, &episodes)
	if err != nil {
		return nil, err
	}
	return episodes.Episodes, nil
}

// Handler para listar os episódios
func EpisodesHandler(w http.ResponseWriter, r *http.Request) {
	episodes, err := loadEpisodes()
	if err != nil {
		http.Error(w, "Erro ao carregar os episódios", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{"episodes": episodes})
}

func main() {
	r := mux.NewRouter()

	// Rota para animes
	r.HandleFunc("/api/animes", AnimesHandler).Methods("GET")

	// Rota para episódios
	r.HandleFunc("/api/episodes", EpisodesHandler).Methods("GET")

	// Definindo a porta do servidor
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Servidor rodando na porta %s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
