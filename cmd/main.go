package main

import (
	"PhoneNumberFormater/app"
)

func main() {
	service := app.MakeService()
	service.Run("G:\\DevEducation\\GoLang\\PhoneNumberFormater\\phones.txt")
}
