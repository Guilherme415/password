package utils_test

import (
	"errors"
	"testing"

	"github.com/Guilherme415/password/utils"
	"github.com/stretchr/testify/assert"
)

func TestVerifyMinStringSize(t *testing.T) {
	t.Run("Should return nil when string size is valid", func(t *testing.T) {
		resp := utils.VerifyMinStringSize("abcabc", 5)

		assert.Equal(t, nil, resp)
	})

	t.Run("Should return error when string size is not valid", func(t *testing.T) {
		expectedResponse := errors.New("minSize")

		resp := utils.VerifyMinStringSize("abcabc", 7)

		assert.Equal(t, expectedResponse, resp)
	})
}

func TestVerifyUpperCaseString(t *testing.T) {
	t.Run("Should return nil when Uppercase chars quantity is valid", func(t *testing.T) {
		resp := utils.VerifyUpperCaseString("ABCabc", 3)

		assert.Equal(t, nil, resp)
	})

	t.Run("Should return error when Uppercase chars quantity is not valid", func(t *testing.T) {
		expectedResponse := errors.New("minUppercase")

		resp := utils.VerifyUpperCaseString("abcabc", 7)

		assert.Equal(t, expectedResponse, resp)
	})
}

func TestVerifyLowerCaseString(t *testing.T) {
	t.Run("Should return nil when Lowercase chars quantity is valid", func(t *testing.T) {
		resp := utils.VerifyLowerCaseString("ABCabc", 3)

		assert.Equal(t, nil, resp)
	})

	t.Run("Should return error when Lowercase chars quantity is not valid", func(t *testing.T) {
		expectedResponse := errors.New("minLowercase")

		resp := utils.VerifyLowerCaseString("abcabc", 7)

		assert.Equal(t, expectedResponse, resp)
	})
}

func TestVerifyDigitsString(t *testing.T) {
	t.Run("Should return nil when minDigit chars quantity is valid", func(t *testing.T) {
		resp := utils.VerifyDigitsString("ABCabc123", 3)

		assert.Equal(t, nil, resp)
	})

	t.Run("Should return error when minDigit chars quantity is not valid", func(t *testing.T) {
		expectedResponse := errors.New("minDigit")

		resp := utils.VerifyDigitsString("abcabc123", 7)

		assert.Equal(t, expectedResponse, resp)
	})
}

func TestVerifySpecialChars(t *testing.T) {
	t.Run("Should return nil when SpecialChars chars quantity is valid", func(t *testing.T) {
		resp := utils.VerifySpecialChars("AB*Ca@bc!3", 3)

		assert.Equal(t, nil, resp)
	})

	t.Run("Should return error when SpecialChars chars quantity is not valid", func(t *testing.T) {
		expectedResponse := errors.New("minSpecialChars")

		resp := utils.VerifySpecialChars("AB*Ca@bc!3", 7)

		assert.Equal(t, expectedResponse, resp)
	})
}

func TestVerifyRepetedChars(t *testing.T) {
	t.Run("Should return nil when no repeted chars", func(t *testing.T) {
		resp := utils.VerifyRepetedChars("AaB*Ca@bc!3")

		assert.Equal(t, nil, resp)
	})

	t.Run("Should return error when repeted chars", func(t *testing.T) {
		expectedResponse := errors.New("noRepeted")

		resp := utils.VerifyRepetedChars("AAB*Ca@bbc!3")

		assert.Equal(t, expectedResponse, resp)
	})
}
