package constants

type Rules string

// Declarando as constantes para as regras das senhas
const (
	RuleMinLowerCase    Rules = "minLowercase"
	RuleMinUpperCase    Rules = "minUppercase"
	RuleMinSize         Rules = "minSize"
	RuleMinDigit        Rules = "minDigit"
	RuleMinSpecialChars Rules = "minSpecialChars"
	RuleNoRepeted       Rules = "noRepeted"
)
