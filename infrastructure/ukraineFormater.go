package infrastructure

import "strings"

const (
	ukraineCode = "+38"
	symbolsToRemove =" +()-"
)

type UkraineFormater struct {

}

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

func (u *UkraineFormater) AddCountryCode(validNumber string) (formatedNumber string) {
	formatedNumber = ukraineCode + validNumber
	return
}

func NewUkraineFormater() *UkraineFormater {
	return &UkraineFormater{}
}
