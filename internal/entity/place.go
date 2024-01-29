package entity

import (
	"errors"
	"unicode"
)

type Place struct {
	City       string `json:"city"`
	PostalCode string `json:"postal_code"`
}

func NormalizePostalCode(postalCode string) (string, error) {
	var normalized_cep string

	for _, char := range postalCode {
		if unicode.IsDigit(char) {
			normalized_cep += string(char)
		}
	}

	if normalized_cep == "" {
		return "", errors.New("empty postal code")
	}

	return normalized_cep, nil
}
