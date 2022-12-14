package main

import (
	"github.com/Guilherme415/password/config"
)

func main() {

	// Inicia servidor GraphQL
	config.NewGraphQLServer()

	// Inicia servidor Rest
	//config.NewRestServer()
}
