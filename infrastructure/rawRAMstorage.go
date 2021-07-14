package infrastructure

import (
	"bufio"
	"log"
	"os"
)

const sourceFilePath = "G:\\DevEducation\\GoLang\\PhoneNumberFormater\\phones.txt"

type RawRAMstorage struct {
	PhoneNumbers []string
}

func (r *RawRAMstorage) getNumbersFromFile(path string) {
	r.PhoneNumbers = []string{}
	file, err := os.Open(path)
	if err != nil {
		log.Println("Failed to open file")
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		r.PhoneNumbers = append(r.PhoneNumbers, scanner.Text())
	}
	file.Close()
}

func (r *RawRAMstorage) GetAllRawNumbers() []string {
	r.getNumbersFromFile(sourceFilePath)
	return r.PhoneNumbers
}

func NewRawRAMstorage() *RawRAMstorage {
	return &RawRAMstorage{}
}