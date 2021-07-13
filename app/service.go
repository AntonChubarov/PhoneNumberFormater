package app

import (
	"PhoneNumberFormater/domain"
	"PhoneNumberFormater/infrastructure"
)

type Service struct {
	domain.RawStorage
	domain.ValidStorage
	domain.Validator
	domain.Formater
	domain.Visualizer
}

func MakeService() *Service {
	s := &Service{}
	s.RawStorage = infrastructure.NewRawRAMstorage()
	s.ValidStorage = infrastructure.NewValidRAMstorage()
	s.Validator = infrastructure.NewRegExValidator()
	s.Formater = infrastructure.NewUkraineFormater()
	s.Visualizer = infrastructure.NewVonsoleVisualizer()
	return s
}

func (s *Service) Run(path string) {
	var temp string
	s.RawStorage.GetNumbersFromFile(path)
	s.Visualizer.Visualize(s.RawStorage.GetAllNumbers())
	for _, number := range s.RawStorage.GetAllNumbers() {
		temp = number
		if !s.Validator.Validate(number) {
			temp = s.Formater.TryToFix(number)
		}
		if s.Validator.Validate(temp) {
			temp = s.Formater.AddCountryCode(temp)
			s.ValidStorage.AddValidNumber(temp)
		}
	}
	s.Visualizer.Visualize(s.ValidStorage.GetAllNumbers())
}

