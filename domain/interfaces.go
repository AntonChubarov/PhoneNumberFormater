package domain

type RawStorage interface {
	GetNumbersFromFile(path string)
	GetAllNumbers() []string
}

type ValidStorage interface {
	AddValidNumber(validNumber ...string)
	GetAllNumbers() []string
}

type Validator interface {
	Validate(number string) (isValid bool)
}

type Formater interface {
	TryToFix(number string) (validNumber string)
	AddCountryCode(validNumber string) (formatedNumber string)
}

type Visualizer interface {
	Visualize(number []string)
}