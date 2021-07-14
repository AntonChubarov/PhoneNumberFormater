package infrastructure

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCase struct {
	Number string
	IsValid bool
}

func TestValidate1(t *testing.T) {
	cases := []TestCase{
		{"0633562669", true},
		{"+38(063)2564446", false},
	}

	validator := NewRegExValidator()

	for i, item := range cases {
		valid := validator.Validate(item.Number)
		//if item.IsValid == valid {
		//	t.Errorf("[%d] Expected result: %v", i, valid)
		//}
		if item.IsValid != valid {
			t.Errorf("[%d] Unexpected result: %v", i, valid)
		}
	}

}

func TestValidate2(t *testing.T) {
	validator := NewRegExValidator()
	valid := validator.Validate("0633562669")
	assert.Equal(t, true, valid)
}

