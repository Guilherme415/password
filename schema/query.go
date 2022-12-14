package schema

import (
	"github.com/graphql-go/graphql"
)

// Gerando os Schemas para GraphQL
func Schema() graphql.Fields {
	return graphql.Fields{
		// Query "verify"
		"verify": &graphql.Field{
			Type: verifyStrongPasswordResponseType,
			Args: graphql.FieldConfigArgument{
				"password": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"rules": &graphql.ArgumentConfig{
					Type: graphql.NewList(rulesQraphQLType),
				},
			},
			Resolve: VerifyResolver,
		},
	}
}
