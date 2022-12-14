package utils

import (
	"regexp"
	"strings"
)

// Remove todos caracteres exceto letras maiúsculas
func RemoveDelimAndLowerCase(str string) string {
	regExp := regexp.MustCompile(`[^[:upper:]]`)
	return regExp.ReplaceAllString(str, "")
}

// Remove todos caracteres exceto letras minúsculas
func RemoveDelimAndUpperCase(str string) string {
	regExp := regexp.MustCompile(`[^[:lower:]]`)
	return regExp.ReplaceAllString(str, "")
}

// Remove todos caracteres exceto números
func RemoveStringAndChars(str string) string {
	re := regexp.MustCompile("[0-9]+")
	numbersArray := re.FindAllString(str, -1)

	numbersString := strings.Join(numbersArray, "")

	return numbersString
}

// Remove todos caracteres exceto caracteres especiais
func RemoveNonSpecialChars(str string) string {
	re := regexp.MustCompile(`["!@#$%^&*()-+\/{}\[\]]+`)
	specialCharsArray := re.FindAllString(str, -1)

	specialCharsString := strings.Join(specialCharsArray, "")

	return specialCharsString
}
