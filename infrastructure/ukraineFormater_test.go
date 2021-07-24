package infrastructure

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type FormatterTestCase struct {
	RawNumber string
	FixedNumber string
	FormatedNumber string
	Error error
}

func TestTryToFix(t *testing.T) {
	cases := []FormatterTestCase{
		{"0633562669", "0633562669", "+380633562669", nil},
		{"0982222222", "0982222222", "+380982222222", nil},
		{"+38(063) 256-44-46", "0632564446", "+380632564446", nil},
		{" +38-063-256-44-46", "0632564446", "+380632564446", nil},
		{"22-60", "2260", "", ErrNumberToShort},
	}

	formatter := NewUkraineFormater()

	for i := range cases {
		fixedNumber := formatter.TryToFix(cases[i].RawNumber)
		if fixedNumber != cases[i].FixedNumber {
			t.Errorf("[%d] Unexpected result: %v, but should be %v", i, fixedNumber, cases[i].FixedNumber)
		}
	}

}

func TestAddCountryCodeAssert(t *testing.T) {
	cases := []FormatterTestCase{
		{"0633562669", "0633562669", "+380633562669", nil},
		{"0982222222", "0982222222", "+380982222222", nil},
		{"+38(063) 256-44-46", "0632564446", "+380632564446", nil},
		{" +38-063-256-44-46", "0632564446", "+380632564446", nil},
	}

	formatter := NewUkraineFormater()

	var formattedNumber string
	var err error

	for i := range cases {
		formattedNumber, err = formatter.AddCountryCode(cases[i].FixedNumber)
		assert.Equal(t, cases[i].FormatedNumber, formattedNumber)
		assert.Equal(t, cases[i].Error, err)
	}
}

func TestAddCountryCodeErrors(t *testing.T) {
	cases := []FormatterTestCase{
		{"", "2260", "", ErrNumberToShort},
		{"", "063555444332", "", ErrNumberToLong},
	}
	formatter := NewUkraineFormater()

	var formattedNumber string
	var err error

	for i := range cases {
		formattedNumber, err = formatter.AddCountryCode(cases[i].FixedNumber)
		assert.Equal(t, cases[i].FormatedNumber, formattedNumber)
		assert.Equal(t, cases[i].Error, err)
	}
}