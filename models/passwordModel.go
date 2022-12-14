package models

// Body para funções que verificam se a senha é forte ou não
type VerifyStrongPasswordBody struct {
	Password string `json:"password"`
	Rules    []Rule `json:"rules"`
}

// Response para funções que verificam se a senha é forte ou não
type VerifyStrongPasswordResponse struct {
	Verify  bool     `json:"verify"`
	NoMatch []string `json:"noMatch"`
}
