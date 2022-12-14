package config

import (
	"github.com/Guilherme415/password/config/dependency"
	"github.com/Guilherme415/password/handlers"
	"github.com/gin-gonic/gin"
)

// Gerando servidor Rest na porta :8080
func NewRestServer() {
	r := gin.Default()

	dependency.Load()

	port := ":8080"

	handlers.Router(r)

	r.Run(port)
}
