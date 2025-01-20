package config

import (
	"github.com/rs/cors"
)

// Função para configurar o middleware CORS
func SetupCORS() *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // Permite todas as origens
	})
}
