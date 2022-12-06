package controllers

import (
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

func (p *PasswordController) VerifyStrongPassword(ctx *gin.Context) {
	verifyStrongPasswordBody := models.VerifyStrongPasswordBody{}
	if err := ctx.BindJSON(&verifyStrongPasswordBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})

		return
	}

	errors := p.passwordBusiness.VerifyStrongPassword(verifyStrongPasswordBody.Password)
	if errors != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": errors,
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
