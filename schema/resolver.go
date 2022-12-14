package schema

import (
	"github.com/Guilherme415/password/business"
	"github.com/Guilherme415/password/models"
	"github.com/graphql-go/graphql"
	"github.com/mitchellh/mapstructure"
)

// Resolver responsavél pela query "verify"
// Verifica se a senha passada está de acordo com as regras
// Se tiver de acordo, a senha será considerada forte
func VerifyResolver(p graphql.ResolveParams) (interface{}, error) {
	passwordBusiness := business.NewPasswordBusiness()

	verifyStrongPasswordBody := paramsToVerifyStrongPasswordBody(p.Args)

	response := passwordBusiness.VerifyStrongPassword(verifyStrongPasswordBody)

	return response, nil
}

// Converte os parametros passado na query "verify" para struct usada na função passwordBusiness.VerifyStrongPassword()
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

// Converte a interface de Rule vinda da query "verify" para a model Rule
func convertInterfaceToRule(from interface{}, to *models.Rule) {
	if trackMap, ok := from.(map[string]interface{}); ok {
		config := &mapstructure.DecoderConfig{TagName: "json", Result: &to}
		decoder, _ := mapstructure.NewDecoder(config)
		decoder.Decode(trackMap)
	}
}
