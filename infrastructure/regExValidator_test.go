package infrastructure

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type ValidatorTestCase struct {
	Number string
	IsValid bool
}

func TestValidate(t *testing.T) {
	cases := []ValidatorTestCase{
		{"0633562669", true},
		{"0982222222", true},
		{"+38(063)2564446", false},
		{"+38-063-256-44-46", false},
		{"22-60", false},
	}

	validator := NewRegExValidator()

	for i, item := range cases {
		valid := validator.Validate(item.Number)
		if item.IsValid != valid {
			t.Errorf("[%d] Unexpected result: %v", i, valid)
		}
	}

}

func TestValidateAssert(t *testing.T) {
	cases := []ValidatorTestCase{
		{"0633562669", true},
		{"0982222222", true},
		{"+38(063)2564446", false},
		{"+38-063-256-44-46", false},
		{"22-60", false},
	}

	validator := NewRegExValidator()

	var valid bool

	for i := range cases {
		valid = validator.Validate(cases[i].Number)
		assert.Equal(t, cases[i].IsValid, valid)
	}
}

