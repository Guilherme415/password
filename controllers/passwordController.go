package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type IPasswordController interface {
	VerifyStrongPassword(ctx *gin.Context)
}

type PasswordController struct {
}

func NewPasswordController() IPasswordController {
	return &PasswordController{}
}

func (p *PasswordController) VerifyStrongPassword(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
