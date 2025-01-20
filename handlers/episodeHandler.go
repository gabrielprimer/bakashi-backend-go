// backend-bakashi-go/handlers/episodes.go
package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"backend-bakashi-go/types"
)

func GetEpisodes(w http.ResponseWriter, r *http.Request) {
	// Carregar o arquivo JSON de episódios
	data, err := ioutil.ReadFile("data/episodes.json")
	if err != nil {
		log.Println("Erro ao ler arquivo de episódios:", err)
		http.Error(w, "Erro ao ler arquivo", http.StatusInternalServerError)
		return
	}

	// Deserializar os dados
	var episodes struct {
		Episodes []types.Episode `json:"episodes"`
	}
	if err := json.Unmarshal(data, &episodes); err != nil {
		log.Println("Erro ao deserializar episódios:", err)
		http.Error(w, "Erro ao processar dados", http.StatusInternalServerError)
		return
	}

	// Definir o cabeçalho de resposta
	w.Header().Set("Content-Type", "application/json")

	// Serializar e enviar os dados de volta como JSON
	if err := json.NewEncoder(w).Encode(episodes); err != nil {
		log.Println("Erro ao enviar resposta:", err)
		http.Error(w, "Erro ao enviar resposta", http.StatusInternalServerError)
	}
}
