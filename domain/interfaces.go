package domain

//go:generate mockgen -source=interfaces.go -destination=mocks.go -package=domain

type RawStorage interface {
	GetAllRawNumbers() []string
}

type ValidStorage interface {
	AddValidNumber(validNumber ...string)
	GetAllValidNumbers() []string
}

type Validator interface {
	Validate(number string) (isValid bool)
}

type Formatter interface {
	TryToFix(number string) (validNumber string)
	AddCountryCode(validNumber string) (formatedNumber string, err error)
}

type Visualizer interface {
	Visualize(number []string)
}