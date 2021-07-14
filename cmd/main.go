package main

import (
	"PhoneNumberFormater/app"
	"PhoneNumberFormater/infrastructure"
)

func main() {
	RawStorage := infrastructure.NewRawRAMstorage()
	ValidStorage := infrastructure.NewValidRAMstorage()
	Validator := infrastructure.NewRegExValidator()
	Formatter := infrastructure.NewUkraineFormater()
	Visualizer := infrastructure.NewConsoleVisualizer()

	service := app.NewService(RawStorage, ValidStorage, Validator, Formatter, Visualizer)
	service.Run()
}
