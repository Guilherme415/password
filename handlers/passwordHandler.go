package handlers

import (
	"github.com/Guilherme415/password/controllers"
	"github.com/gin-gonic/gin"
)

// Rotas da controller Password
func PasswordHandler(g *gin.RouterGroup, c controllers.IPasswordController) {
	g.POST("", c.VerifyStrongPassword)
}
