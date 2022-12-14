package schema

import (
	"github.com/Guilherme415/password/business"
	"github.com/Guilherme415/password/models"
	"github.com/graphql-go/graphql"
	"github.com/mitchellh/mapstructure"
)

func VerifyResolver(p graphql.ResolveParams) (interface{}, error) {
	passwordBusiness := business.NewPasswordBusiness()

	verifyStrongPasswordBody := paramsToVerifyStrongPasswordBody(p.Args)

	response := passwordBusiness.VerifyStrongPassword(verifyStrongPasswordBody)

	return response, nil
}

func paramsToVerifyStrongPasswordBody(args map[string]interface{}) models.VerifyStrongPasswordBody {
	password := args["password"].(string)
	rules := args["rules"].([]interface{})

	rulesParsed := []models.Rule{}

	for index := range rules {
		rule := models.Rule{}
		convertInterfaceToRule(rules[index], &rule)

		rulesParsed = append(rulesParsed, rule)
	}

	verifyStrongPasswordBody := models.VerifyStrongPasswordBody{
		Password: password,
		Rules:    rulesParsed,
	}

	return verifyStrongPasswordBody
}

func convertInterfaceToRule(from interface{}, to *models.Rule) {
	if trackMap, ok := from.(map[string]interface{}); ok {
		config := &mapstructure.DecoderConfig{TagName: "json", Result: &to}
		decoder, _ := mapstructure.NewDecoder(config)
		decoder.Decode(trackMap)
	}
}
