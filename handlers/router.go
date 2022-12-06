package handlers

import (
	"github.com/Guilherme415/password/config/dependency"
	"github.com/gin-gonic/gin"
)

func Router(e *gin.Engine) {
	v1 := e.Group("/v1")

	PasswordHandler(v1.Group("verify"), dependency.PasswordController)
}
