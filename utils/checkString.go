package utils

import (
	"errors"

	"github.com/Guilherme415/password/config/constants"
)

// Verifica se o tamanho da string passada como parametro
// é menor que o tamanho mínimo, passado como parametro
func VerifyMinStringSize(paramString string, minSize int) error {
	if len(paramString) < minSize {
		return errors.New(string(constants.RuleMinSize))
	}

	return nil
}

// Verifica se o número de letras maiúsculas da string passada como parametro
// é menor que a quantidade mínima, passada como parametro
func VerifyUpperCaseString(paramString string, minSize int) error {
	upperCaseChars := RemoveDelimAndLowerCase(paramString)

	if len(upperCaseChars) < minSize {
		return errors.New(string(constants.RuleMinUpperCase))
	}

	return nil
}

// Verifica se o número de letras minusculas da string passada como parametro
// é menor que a quantidade mínima, passada como parametro
func VerifyLowerCaseString(paramString string, minSize int) error {
	lowerCaseChars := RemoveDelimAndUpperCase(paramString)

	if len(lowerCaseChars) < minSize {
		return errors.New(string(constants.RuleMinLowerCase))
	}

	return nil
}

// Verifica se a quantidade de números na string passada como parametro
// é menor que a quantidade mínima, passada como parametro
func VerifyDigitsString(paramString string, minSize int) error {
	digits := RemoveStringAndChars(paramString)

	if len(digits) < minSize {
		return errors.New(string(constants.RuleMinDigit))
	}

	return nil
}

// Verifica se o número de caracteres especiais na string passada como parametro
// é menor que a quantidade mínima, passada como parametro
func VerifySpecialChars(paramString string, minSize int) error {
	specialChars := RemoveNonSpecialChars(paramString)

	if len(specialChars) < minSize {
		return errors.New(string(constants.RuleMinSpecialChars))
	}

	return nil
}

// Verifica se há letras repetidas na sequencia
// aab retorna o erro: noRepeted, indicando que existe letras repetidas
// aba retorna nil, indicando que não há letra repetidas
func VerifyRepetedChars(paramString string) error {
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
