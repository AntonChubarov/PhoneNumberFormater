package infrastructure

import (
	"errors"
	"strings"
)

const (
	ukraineCode = "+38"
	symbolsToRemove =" +()-"
)

type UkraineFormater struct {

}

var ErrNumberToShort = errors.New("Valid number should contain 10 digits")
var ErrNumberToLong = errors.New("Valid number should contain 10 digits")

func (u *UkraineFormater) TryToFix(number string) (validNumber string) {
	validNumber = number
	for _, symbol := range symbolsToRemove {
		validNumber = strings.ReplaceAll(validNumber, string(symbol), "")
	}
	if len(validNumber) > 10 {
		validNumber = validNumber[len(validNumber)-10:]
	}
	return
}

func (u *UkraineFormater) AddCountryCode(validNumber string) (formatedNumber string, err error) {
	if len(validNumber) < 10 {
		return "", ErrNumberToShort
	}

	if len(validNumber) > 10 {
		return "", ErrNumberToLong
	}

	formatedNumber = ukraineCode + validNumber
	return
}

func NewUkraineFormater() *UkraineFormater {
	return &UkraineFormater{}
}
