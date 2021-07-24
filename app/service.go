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

func NewService(rawStorage domain.RawStorage,
	validStorage domain.ValidStorage,
	validator domain.Validator,
	formatter domain.Formatter,
	visualizer domain.Visualizer) *Service {
	return &Service{
		rawStorage,
		validStorage,
		validator,
		formatter,
		visualizer,
	}
}

func (s *Service) Run() {
	var temp string
	var err error
	rawNumbers := s.rawStorage.GetAllRawNumbers()
	s.visualizer.Visualize(rawNumbers)
	for i := range rawNumbers {
		temp = rawNumbers[i]
		if !s.validator.Validate(rawNumbers[i]) {
			temp = s.formatter.TryToFix(rawNumbers[i])
		}
		if s.validator.Validate(temp) {
			temp, err = s.formatter.AddCountryCode(temp)
			if err != nil {
				log.Println(err)
			}
			s.validStorage.AddValidNumber(temp)
		} else {
			log.Printf("Can't to format number %v, please, check it!\n", rawNumbers[i])
		}
	}
	s.visualizer.Visualize(s.validStorage.GetAllValidNumbers())
}

