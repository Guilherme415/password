package dependency

import (
	"github.com/Guilherme415/password/business"
	"github.com/Guilherme415/password/controllers"
)

// Controllers
var (
	PasswordController controllers.IPasswordController
)

// Business
var (
	PasswordBusiness business.IPasswordBusiness
)

// Injetando as dependencias da aplicação
func Load() {
	// Repositories

	// Business
	PasswordBusiness = business.NewPasswordBusiness()

	// Controllers
	PasswordController = controllers.NewPasswordController(PasswordBusiness)
}
