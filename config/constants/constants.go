package constants

type Rules string

const (
	RuleMinLowerCase    Rules = "minLowercase"
	RuleMinUpperCase    Rules = "minUppercase"
	RuleMinSize         Rules = "minSize"
	RuleMinDigit        Rules = "minDigit"
	RuleMinSpecialChars Rules = "minSpecialChars"
	RuleNoRepeted       Rules = "noRepeted"
)
