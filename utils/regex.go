package utils

import (
	"regexp"
	"strings"
)

func RemoveDelimAndLowerCase(str string) string {
	regExp := regexp.MustCompile(`[^[:upper:]]`)
	return regExp.ReplaceAllString(str, "")
}

func RemoveDelimAndUpperCase(str string) string {
	regExp := regexp.MustCompile(`[^[:lower:]]`)
	return regExp.ReplaceAllString(str, "")
}

func RemoveStringAndChars(str string) string {
	re := regexp.MustCompile("[0-9]+")
	numbersArray := re.FindAllString(str, -1)

	numbersString := strings.Join(numbersArray, "")

	return numbersString
}

func RemoveNonSpecialChars(str string) string {
	re := regexp.MustCompile(`["!@#$%^&*()-+\/{}\[\]]+`)
	specialCharsArray := re.FindAllString(str, -1)

	specialCharsString := strings.Join(specialCharsArray, "")

	return specialCharsString
}
