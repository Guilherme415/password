package constants

type Rules string

const (
	RuleMinLowerCase    Rules = "minLowerCase"
	RuleMinUpperCase    Rules = "minUpperCase"
	RuleMinSize         Rules = "minSize"
	RuleMinDigit        Rules = "minDigit"
	RuleMinSpecialChars Rules = "minSpecialChars"
	RuleNoRepeted       Rules = "noRepeted"
)
