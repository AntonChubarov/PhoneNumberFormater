package app

import (
	"PhoneNumberFormater/domain"
	"PhoneNumberFormater/infrastructure"
	"github.com/golang/mock/gomock"
	"testing"
)

type Case struct {
	desc string
	rawStorage domain.RawStorage
	validStorage domain.ValidStorage
	validator domain.Validator
	formatter domain.Formatter
	visualizer domain.Visualizer
}

func TestService(t *testing.T)  {
	ctrl := gomock.NewController(t)

	rawNumber := "8-050-203-04-50"
	validNumber := "0502030450"
	formattedNumber := "+380502030450"

	cases := []Case {
		{
			desc:"Function GetAllRawNumbers() of RawStorage should be called once when service Run() executing",
			rawStorage: func(ctrl *gomock.Controller) *domain.MockRawStorage {
				mock:=domain.NewMockRawStorage(ctrl)
				mock.EXPECT().GetAllRawNumbers().Return([]string{}).Times(1)
				return mock
			}(ctrl),
			validStorage: infrastructure.NewValidRAMstorage(),
			validator: infrastructure.NewRegExValidator(),
			formatter: infrastructure.NewUkraineFormater(),
			visualizer: infrastructure.NewConsoleVisualizer(),
		},
		{
			desc:"Function Visualize() of Visualizer should be called twice when service Run() executing",
			rawStorage: func(ctrl *gomock.Controller) *domain.MockRawStorage {
				mock:=domain.NewMockRawStorage(ctrl)
				mock.EXPECT().GetAllRawNumbers().Return([]string{})
				return mock
			}(ctrl),
			validStorage: infrastructure.NewValidRAMstorage(),
			validator: infrastructure.NewRegExValidator(),
			formatter: infrastructure.NewUkraineFormater(),
			visualizer: func(ctrl *gomock.Controller) *domain.MockVisualizer {
				mock:=domain.NewMockVisualizer(ctrl)
				mock.EXPECT().Visualize([]string{}).Times(2)
				return mock
			}(ctrl),
		},
		{
			desc:"Function Validate() of Validator should be called twice when service Run() executing",
			rawStorage: func(ctrl *gomock.Controller) *domain.MockRawStorage {
				mock:=domain.NewMockRawStorage(ctrl)
				mock.EXPECT().GetAllRawNumbers().Return([]string{""})
				return mock
			}(ctrl),
			validStorage: infrastructure.NewValidRAMstorage(),
			validator: func(ctrl *gomock.Controller) *domain.MockValidator {
				mock:=domain.NewMockValidator(ctrl)
				mock.EXPECT().Validate("").Times(2)
				return mock
			}(ctrl),
			formatter: infrastructure.NewUkraineFormater(),
			visualizer: infrastructure.NewConsoleVisualizer(),
		},
		{
			desc:`Function TryToFix() and AddCountryCode() of Formatter should be once
when GetAllRawNumbers() of RawStorage returns one invalid number`,
			rawStorage: func(ctrl *gomock.Controller) *domain.MockRawStorage {
				mock:=domain.NewMockRawStorage(ctrl)
				mock.EXPECT().GetAllRawNumbers().Return([]string{rawNumber})
				return mock
			}(ctrl),
			validStorage: infrastructure.NewValidRAMstorage(),
			validator: infrastructure.NewRegExValidator(),
			formatter: func(ctrl *gomock.Controller) *domain.MockFormatter {
				mock:=domain.NewMockFormatter(ctrl)
				mock.EXPECT().TryToFix(rawNumber).Return(validNumber).Times(1)
				mock.EXPECT().AddCountryCode(validNumber).Return(formattedNumber, nil).Times(1)
				return mock
			}(ctrl),
			visualizer: infrastructure.NewConsoleVisualizer(),
		},
		{
			desc:`Function AddValidNumber() and GetAllValidNumbers() of ValidStorage
should be called once when GetAllRawNumbers() of RawStorage returns one number`,
			rawStorage: func(ctrl *gomock.Controller) *domain.MockRawStorage {
				mock:=domain.NewMockRawStorage(ctrl)
				mock.EXPECT().GetAllRawNumbers().Return([]string{rawNumber})
				return mock
			}(ctrl),
			validStorage: func(ctrl *gomock.Controller) *domain.MockValidStorage {
				mock:=domain.NewMockValidStorage(ctrl)
				mock.EXPECT().AddValidNumber(formattedNumber).Times(1)
				mock.EXPECT().GetAllValidNumbers().Return([]string{formattedNumber}).Times(1)
				return mock
			}(ctrl),
			validator: infrastructure.NewRegExValidator(),
			formatter: infrastructure.NewUkraineFormater(),
			visualizer: infrastructure.NewConsoleVisualizer(),
		},
	}

	for i := range cases {
		t.Run(cases[i].desc, func(t *testing.T) {
			service := NewService(cases[i].rawStorage, cases[i].validStorage, cases[i].validator,
				cases[i].formatter,	cases[i].visualizer)
			service.Run()
		})
	}
}
