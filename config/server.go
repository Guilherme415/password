package config

import (
	"github.com/Guilherme415/password/config/dependency"
	"github.com/Guilherme415/password/handlers"
	"github.com/gin-gonic/gin"
)

func NewServer() {
	r := gin.Default()

	dependency.Load()

	port := ":8080"

	handlers.Router(r)

	r.Run(port)
}
