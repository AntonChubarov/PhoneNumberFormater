package app

import (
	"PhoneNumberFormater/domain"
	"PhoneNumberFormater/infrastructure"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestServiceExecution(t *testing.T)  {
	ctrl := gomock.NewController(t)
	//rawNumber := "8-050-203-04-50"
	//validNumber := "0502030450"
	//formattedNumber := "+380502030450"

	cases := []struct {
		desc string
		rawStorage domain.RawStorage
		validStorage domain.ValidStorage
		validator domain.Validator
		formatter domain.Formatter
		visualizer domain.Visualizer
	} {
		{
			desc:"Some description",
			rawStorage: func(ctrl *gomock.Controller) *domain.MockRawStorage {
				mock:=domain.NewMockRawStorage(ctrl)
				mock.EXPECT().GetAllRawNumbers().Return([]string{}).Times(1)
				return mock
			}(ctrl),
		},
	}

	for _, c := range cases {
		t.Run(c.desc, func(t *testing.T) {
			service := NewService(
				c.rawStorage,
				infrastructure.NewValidRAMstorage(),
				infrastructure.NewRegExValidator(),
				infrastructure.NewUkraineFormater(),
				infrastructure.NewConsoleVisualizer(),
				)
			// act
			service.Run()
			// assert
			//assert.Equal(t, c.wantError, gotError)
			//assert.Equal(t, c.wantUser, gotUser)
		})
	}
}
