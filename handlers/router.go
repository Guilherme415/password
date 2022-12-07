package handlers

import (
	"github.com/Guilherme415/password/config/dependency"
	"github.com/gin-gonic/gin"
)

func Router(e *gin.Engine) {
	PasswordHandler(e.Group("verify"), dependency.PasswordController)
}
