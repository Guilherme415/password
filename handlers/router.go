package handlers

import (
	"github.com/Guilherme415/password/config/dependency"
	"github.com/gin-gonic/gin"
)

// Rotas da API
// Com pre-fix na rota e já realizando a injeção de dependencia
func Router(e *gin.Engine) {
	PasswordHandler(e.Group("verify"), dependency.PasswordController)
}
