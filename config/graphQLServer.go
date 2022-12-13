package config

import (
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

func NewGraphQLServer() {
	schema, _ := graphql.NewSchema(graphql.SchemaConfig{})

	h := handler.New(&handler.Config{
		Schema:     &schema,
		Pretty:     true,
		GraphiQL:   false,
		Playground: true,
	})

	http.Handle("/graphql", h)
	http.ListenAndServe(":8080", nil)
}
