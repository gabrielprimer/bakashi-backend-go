package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"backend-bakashi-go/api/types"
)

func AnimesHandler(w http.ResponseWriter, r *http.Request) {
	animes, err := loadAnimes()
	if err != nil {
		log.Printf("Erro ao carregar os animes: %v", err)
		http.Error(w, "Erro ao carregar os animes", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{"animes": animes})
}

func loadAnimes() ([]types.Anime, error) {
	dataPath := "api/data/animes.json"
	data, err := os.ReadFile(dataPath)
	if err != nil {
		return nil, err
	}
	var result struct {
		Animes []types.Anime `json:"animes"`
	}
	err = json.Unmarshal(data, &result)
	return result.Animes, err
}
