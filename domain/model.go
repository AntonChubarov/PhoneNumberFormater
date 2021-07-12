package model

type RawPhoneNumberStorage interface {
	AddNumber(number ...string)
	GetNumber() string
	GetAllNumbers() []string
	RemoveNumber(number string)
}

type ValidPhoneNumberStorage interface {
	AddValidNumber(validNumber ...string)
	GetAllNumbers() []string
}

type PhoneNumberValidator interface {
	Validate(number string) (validNumber string, err error)
}

type PhoneNumberFormater interface {
	AddCountryCode(validNumber string) (formatedNumber string)
}

type Printer interface {
	Print(number ...string)
}