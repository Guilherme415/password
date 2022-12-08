package controllers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Guilherme415/password/business"
	"github.com/Guilherme415/password/config/constants"
	"github.com/Guilherme415/password/controllers"
	"github.com/Guilherme415/password/models"
	"github.com/gin-gonic/gin"
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

	t.Run("/verify should return Ok", func(t *testing.T) {
		expectedResponse := models.VerifyStrongPasswordResponse{
			Verify:  true,
			NoMatch: []string{},
		}

		passwordBusiness := business.NewPasswordBusiness()
		passwordController := controllers.NewPasswordController(passwordBusiness)

		// Server start
		g := gin.Default()
		g.POST("/verify", passwordController.VerifyStrongPassword)
		w := httptest.NewRecorder()

		body, _ := json.Marshal(mockValidStrongPassword)

		req, _ := http.NewRequest(http.MethodPost, "/verify", bytes.NewBuffer(body))
		req.Header.Add("Content-Type", "application/json")
		g.ServeHTTP(w, req)

		response := models.VerifyStrongPasswordResponse{}
		err := json.Unmarshal(w.Body.Bytes(), &response)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, expectedResponse, response)
	})

	t.Run("/verify should return BadRequest when has a invalid body", func(t *testing.T) {
		expectedResponse := models.Response{
			Code:    fmt.Sprint(http.StatusBadRequest),
			Message: "invalid character 'a' looking for beginning of value",
		}

		passwordBusiness := business.NewPasswordBusiness()
		passwordController := controllers.NewPasswordController(passwordBusiness)

		// Server start
		g := gin.Default()
		g.POST("/verify", passwordController.VerifyStrongPassword)
		w := httptest.NewRecorder()

		req, _ := http.NewRequest(http.MethodPost, "/verify", bytes.NewBufferString("abcabac"))
		req.Header.Add("Content-Type", "application/json")
		g.ServeHTTP(w, req)

		response := models.Response{}
		err := json.Unmarshal(w.Body.Bytes(), &response)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, expectedResponse, response)
	})
}
