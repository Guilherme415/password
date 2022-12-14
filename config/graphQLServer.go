package config

import (
	"log"
	"net/http"

	"github.com/Guilherme415/password/schema"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

func NewGraphQLServer() {
	// Schema
	fields := schema.Schema()

	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}

	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}

	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	h := handler.New(&handler.Config{
		Schema:     &schema,
		Pretty:     true,
		GraphiQL:   false,
		Playground: true,
	})

	http.Handle("/graphql", h)
	http.ListenAndServe(":8080", nil)
}
