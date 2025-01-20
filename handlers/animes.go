package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"backend-bakashi-go/types"  // Certifique-se de que o caminho está correto
)

func GetAnimes(w http.ResponseWriter, r *http.Request) {
	// Carregar o arquivo JSON de animes
	data, err := ioutil.ReadFile("data/animes.json")
	if err != nil {
		log.Println("Erro ao ler arquivo de animes:", err)
		http.Error(w, "Erro ao ler arquivo", http.StatusInternalServerError)
		return
	}

	// Deserializar os dados
	var animes struct {
		Animes []types.Anime `json:"animes"`
	}
	if err := json.Unmarshal(data, &animes); err != nil {
		log.Println("Erro ao deserializar animes:", err)
		http.Error(w, "Erro ao processar dados", http.StatusInternalServerError)
		return
	}

	// Definir o cabeçalho de resposta
	w.Header().Set("Content-Type", "application/json")

	// Serializar e enviar os dados de volta como JSON
	if err := json.NewEncoder(w).Encode(animes); err != nil {
		log.Println("Erro ao enviar resposta:", err)
		http.Error(w, "Erro ao enviar resposta", http.StatusInternalServerError)
	}
}
