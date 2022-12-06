package handlers

import (
	"github.com/Guilherme415/password/controllers"
	"github.com/gin-gonic/gin"
)

func PasswordHandler(g *gin.RouterGroup, c controllers.IPasswordController) {
	g.GET("", c.VerifyStrongPassword)
}
