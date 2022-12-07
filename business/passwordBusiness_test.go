package business_test

import (
	"testing"

	"github.com/Guilherme415/password/business"
	"github.com/Guilherme415/password/config/constants"
	"github.com/Guilherme415/password/models"
	"github.com/stretchr/testify/assert"
)

func TestVerifyStrongPassword(t *testing.T) {
	mockValidStrongPassword := models.VerifyStrongPasswordBody{
		Password: "TesteSenhaForte!123&",
		Rules: []models.Rule{
			{
				Rule:  string(constants.RuleMinSize),
				Value: 8,
			},
			{
				Rule:  string(constants.RuleMinLowerCase),
				Value: 5,
			},
			{
				Rule:  string(constants.RuleMinUpperCase),
				Value: 3,
			},
			{
				Rule:  string(constants.RuleMinDigit),
				Value: 3,
			},
			{
				Rule:  string(constants.RuleMinSpecialChars),
				Value: 2,
			},
			{
				Rule:  string(constants.RuleNoRepeted),
				Value: 0,
			},
		},
	}

	mockInvalidStrongPassword := models.VerifyStrongPasswordBody{
		Password: "TesteSenhaaForte!123&",
		Rules: []models.Rule{
			{
				Rule:  string(constants.RuleMinSize),
				Value: 30,
			},
			{
				Rule:  string(constants.RuleMinLowerCase),
				Value: 20,
			},
			{
				Rule:  string(constants.RuleMinUpperCase),
				Value: 20,
			},
			{
				Rule:  string(constants.RuleMinDigit),
				Value: 20,
			},
			{
				Rule:  string(constants.RuleMinSpecialChars),
				Value: 30,
			},
			{
				Rule:  string(constants.RuleNoRepeted),
				Value: 0,
			},
		},
	}

	t.Run("Should return a valid strong password", func(t *testing.T) {
		expectedResponse := models.VerifyStrongPasswordResponse{
			Verify:  true,
			NoMatch: []string{},
		}

		passwordBusiness := business.NewPasswordBusiness()

		resp := passwordBusiness.VerifyStrongPassword(mockValidStrongPassword)

		assert.Equal(t, expectedResponse, resp)
	})

	t.Run("Should return an invalid strong password", func(t *testing.T) {
		expectedResponse := models.VerifyStrongPasswordResponse{
			Verify: false,
			NoMatch: []string{
				string(constants.RuleMinSize),
				string(constants.RuleMinLowerCase),
				string(constants.RuleMinUpperCase),
				string(constants.RuleMinDigit),
				string(constants.RuleMinSpecialChars),
				string(constants.RuleNoRepeted),
			},
		}

		passwordBusiness := business.NewPasswordBusiness()

		resp := passwordBusiness.VerifyStrongPassword(mockInvalidStrongPassword)

		assert.Equal(t, expectedResponse, resp)
	})
}
