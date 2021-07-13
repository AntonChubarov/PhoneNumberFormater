package infrastructure

import (
	"bufio"
	"log"
	"os"
)

type RawRAMstorage struct {
	PhoneNumbers []string
}

func (r *RawRAMstorage) GetNumbersFromFile(path string) {
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

func (r *RawRAMstorage) GetAllNumbers() []string {
	return r.PhoneNumbers
}

func NewRawRAMstorage() *RawRAMstorage {
	r := &RawRAMstorage{}
	r.PhoneNumbers = []string{}
	return r
}