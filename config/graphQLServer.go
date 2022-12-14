package config

import (
	"log"
	"net/http"

	"github.com/Guilherme415/password/schema"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

// Gerando servidor GraphQL na porta :8080 e na rota "/graphql"
func NewGraphQLServer() {
	// Schema
	fields := schema.Schema()

	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}

	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}

	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	graphQLHandler := handler.New(&handler.Config{
		Schema:     &schema,
		Pretty:     true,
		GraphiQL:   false,
		Playground: true,
	})

	http.Handle("/graphql", graphQLHandler)
	http.ListenAndServe(":8080", nil)
}
