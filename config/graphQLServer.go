package config

import (
	"log"
	"net/http"

	"github.com/Guilherme415/password/business"
	"github.com/Guilherme415/password/models"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/mitchellh/mapstructure"
)

func NewGraphQLServer() {
	// Schema
	fields := schema()

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

var trackType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "Track",
	Fields: graphql.InputObjectConfigFieldMap{
		"rule": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"value": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
	},
})

func schema() graphql.Fields {
	return graphql.Fields{
		"verify": &graphql.Field{
			Type: graphql.String,
			Args: graphql.FieldConfigArgument{
				"password": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"rules": &graphql.ArgumentConfig{
					Type: graphql.NewList(trackType),
				},
			},
			Resolve: resolver,
		},
	}
}

func resolver(p graphql.ResolveParams) (interface{}, error) {
	passwordBusiness := business.NewPasswordBusiness()

	password := p.Args["password"].(string)

	rules := p.Args["rules"].([]interface{})

	rulesParsed := []models.Rule{}

	for index := range rules {
		var track models.Rule
		convertInterfaceToRule(rules[index], &track)

		rulesParsed = append(rulesParsed, track)
	}

	response := passwordBusiness.VerifyStrongPassword(models.VerifyStrongPasswordBody{
		Password: password,
		Rules:    rulesParsed,
	})

	return response, nil
}

func convertInterfaceToRule(from interface{}, to *models.Rule) {
	if trackMap, ok := from.(map[string]interface{}); ok {
		config := &mapstructure.DecoderConfig{TagName: "json", Result: &to}
		decoder, _ := mapstructure.NewDecoder(config)
		decoder.Decode(trackMap)
	}
}
