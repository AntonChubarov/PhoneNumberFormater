package infrastructure

import (
	"log"
	"regexp"
)

const (
	phoneRegEx = "^[0][0-9]{9}$"
)

type RegExValidator struct {

}

func (r *RegExValidator) Validate(number string) (isValid bool) {
	isValid, err := regexp.MatchString(phoneRegEx, number)
	if err != nil {
		log.Println(err)
	}
	return
}

func NewRegExValidator() *RegExValidator {
	return &RegExValidator{}
}
