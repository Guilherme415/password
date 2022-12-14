package business

import (
	"github.com/Guilherme415/password/config/constants"
	"github.com/Guilherme415/password/models"
	"github.com/Guilherme415/password/utils"
)

type IPasswordBusiness interface {
	VerifyStrongPassword(verifyStrongPasswordBody models.VerifyStrongPasswordBody) models.VerifyStrongPasswordResponse
}

type PasswordBusiness struct {
}

func NewPasswordBusiness() IPasswordBusiness {
	return &PasswordBusiness{}
}

// Função que recebe uma senha e as regras
// Retorna se a senha é forte (verify: boolean)
// Quando a senha é fraca, em "noMatch", é retornado o que faltou para ficar forte
func (p *PasswordBusiness) VerifyStrongPassword(verifyStrongPasswordBody models.VerifyStrongPasswordBody) models.VerifyStrongPasswordResponse {
	verify := true

	invalidRules := strongPasswordRules(verifyStrongPasswordBody)
	if len(invalidRules) > 0 {
		verify = false
	}

	response := models.VerifyStrongPasswordResponse{
		Verify:  verify,
		NoMatch: invalidRules,
	}

	return response
}

// Função que recebe uma senha e as regras
// Retorna o que faltou para senha ficar forte
// Caso retornado nada ([]), a senha é forte
func strongPasswordRules(verifyStrongPasswordBody models.VerifyStrongPasswordBody) []string {
	invalidRules := []string{}

	for _, rule := range verifyStrongPasswordBody.Rules {
		err := verifyRule(verifyStrongPasswordBody.Password, rule)
		if err != nil {
			invalidRules = append(invalidRules, err.Error())
		}
	}

	return invalidRules
}

// Recebe uma senha e uma regra
// Retorna se a senha cumpre a regra específica
// Caso cumpra, retornará nil
// Caso não cumpra, retornará o motivo
func verifyRule(password string, rule models.Rule) error {
	switch rule.Rule {
	case string(constants.RuleMinSize):
		err := utils.VerifyMinStringSize(password, rule.Value)
		if err != nil {
			return err
		}

	case string(constants.RuleMinUpperCase):
		err := utils.VerifyUpperCaseString(password, rule.Value)
		if err != nil {
			return err
		}

	case string(constants.RuleMinLowerCase):
		err := utils.VerifyLowerCaseString(password, rule.Value)
		if err != nil {
			return err
		}

	case string(constants.RuleMinDigit):
		err := utils.VerifyDigitsString(password, rule.Value)
		if err != nil {
			return err
		}

	case string(constants.RuleMinSpecialChars):
		err := utils.VerifySpecialChars(password, rule.Value)
		if err != nil {
			return err
		}

	case string(constants.RuleNoRepeted):
		err := utils.VerifyRepetedChars(password)
		if err != nil {
			return err
		}
	}

	return nil
}
