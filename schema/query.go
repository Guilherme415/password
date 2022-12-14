package schema

import "github.com/graphql-go/graphql"

func Schema() graphql.Fields {
	return graphql.Fields{
		"verify": &graphql.Field{
			Type: graphql.String,
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

var rulesQraphQLType = graphql.NewInputObject(graphql.InputObjectConfig{
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
