package business

type IPasswordBusiness interface {
	VerifyStrongPassword(password string) []error
}

type PasswordBusiness struct {
}

func NewPasswordBusiness() IPasswordBusiness {
	return &PasswordBusiness{}
}

func (p *PasswordBusiness) VerifyStrongPassword(password string) []error {
	return nil
}
