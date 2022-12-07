package utils

import (
	"errors"

	"github.com/Guilherme415/password/config/constants"
)

func VerifyMinStringSize(paramString string, minSize int) error {
	if len(paramString) < minSize {
		return errors.New(string(constants.RuleMinSize))
	}

	return nil
}

func VerifyUpperCaseString(paramString string, minSize int) error {
	upperCaseChars := RemoveDelimAndLowerCase(paramString)

	if len(upperCaseChars) < minSize {
		return errors.New(string(constants.RuleMinUpperCase))
	}

	return nil
}

func VerifyLowerCaseString(paramString string, minSize int) error {
	lowerCaseChars := RemoveDelimAndUpperCase(paramString)

	if len(lowerCaseChars) < minSize {
		return errors.New(string(constants.RuleMinLowerCase))
	}

	return nil
}

func VerifyDigitsString(paramString string, minSize int) error {
	digits := RemoveStringAndChars(paramString)

	if len(digits) < minSize {
		return errors.New(string(constants.RuleMinDigit))
	}

	return nil
}

func VerifySpecialChars(paramString string, minSize int) error {
	specialChars := RemoveNonSpecialChars(paramString)

	if len(specialChars) < minSize {
		return errors.New(string(constants.RuleMinSpecialChars))
	}

	return nil
}

func VerifyRepetedChars(paramString string, minSize int) error {
	for i := range paramString {
		if i+1 == len(paramString) {
			break
		}

		if paramString[i] == paramString[i+1] {
			return errors.New(string(constants.RuleNoRepeted))
		}
	}

	return nil
}
