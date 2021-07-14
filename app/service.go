package app

import (
	"PhoneNumberFormater/domain"
	"log"
)

type Service struct {
	rawStorage domain.RawStorage
	validStorage domain.ValidStorage
	validator domain.Validator
	formatter domain.Formatter
	visualizer domain.Visualizer
}

func NewService(RawStorage domain.RawStorage,
	ValidStorage domain.ValidStorage,
	Validator domain.Validator,
	Formatter domain.Formatter,
	Visualizer domain.Visualizer) *Service {
	return &Service{
		RawStorage,
		ValidStorage,
		Validator,
		Formatter,
		Visualizer,
	}
}

func (s *Service) Run() {
	var temp string
	rawNumbers := s.rawStorage.GetAllRawNumbers()
	s.visualizer.Visualize(rawNumbers)
	for i := range rawNumbers {
		temp = rawNumbers[i]
		if !s.validator.Validate(rawNumbers[i]) {
			temp = s.formatter.TryToFix(rawNumbers[i])
		}
		if s.validator.Validate(temp) {
			temp = s.formatter.AddCountryCode(temp)
			s.validStorage.AddValidNumber(temp)
		} else {
			log.Printf("Can't to format number %v, please, check it!\n", rawNumbers[i])
		}
	}
	s.visualizer.Visualize(s.validStorage.GetAllValidNumbers())
}

