package controllers

import (
	"fmt"
	"net/http"

	"github.com/Guilherme415/password/business"
	"github.com/Guilherme415/password/models"
	"github.com/gin-gonic/gin"
)

type IPasswordController interface {
	VerifyStrongPassword(ctx *gin.Context)
}

type PasswordController struct {
	passwordBusiness business.IPasswordBusiness
}

func NewPasswordController(passwordBusiness business.IPasswordBusiness) IPasswordController {
	return &PasswordController{passwordBusiness}
}

// Rota POST /verify
// Body: {
//	"password": string,
//	"rules": [
//		{"rule": string,"value": int}
//	]
// }
// StatusCodes possíveis: 400 (BadRequest), 200 (Ok)
// Response: {
//	"verify": boolean,
//	"noMatch": [string]
// }
func (p *PasswordController) VerifyStrongPassword(ctx *gin.Context) {
	verifyStrongPasswordBody := models.VerifyStrongPasswordBody{}
	if err := ctx.BindJSON(&verifyStrongPasswordBody); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Code:    fmt.Sprint(http.StatusBadRequest),
			Message: err.Error(),
		})

		return
	}

	response := p.passwordBusiness.VerifyStrongPassword(verifyStrongPasswordBody)

	ctx.JSON(http.StatusOK, response)
}
