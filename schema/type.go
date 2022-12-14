package schema

import "github.com/graphql-go/graphql"

// Tipagem da a model Rule para GraphQL
var rulesQraphQLType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "Rule",
	Fields: graphql.InputObjectConfigFieldMap{
		"rule": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"value": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
	},
})

// Criando tipagem para o retorno da query Verify
var verifyStrongPasswordResponseType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "VerifyStrongPasswordResponse",
		Fields: graphql.Fields{
			"verify": &graphql.Field{
				Type: graphql.Boolean,
			},
			"noMatch": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
		},
	},
)
