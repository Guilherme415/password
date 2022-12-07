package models

type VerifyStrongPasswordBody struct {
	Password string `json:"password"`
	Rules    []Rule `json:"rules"`
}

type VerifyStrongPasswordResponse struct {
	Verify  bool     `json:"verify"`
	NoMatch []string `json:"noMatch"`
}
