package dependency

import "github.com/Guilherme415/password/controllers"

// Controllers
var (
	PasswordController controllers.IPasswordController
)

func Load() {
	// Repositories

	// Business

	// Controllers
	PasswordController = controllers.NewPasswordController()
}
