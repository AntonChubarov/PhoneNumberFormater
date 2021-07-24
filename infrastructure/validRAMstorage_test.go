package infrastructure

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddNumbersToValidStorage(t *testing.T) {
	var validNumbers = []string{
		"+380633562669",
		"+380982222222",
		"+380632564446",
		"+380632564446",
	}

	validStorage := NewValidRAMstorage()

	var numbers []string
	var numbersCount int

	for i := 0; i < len(validNumbers); i++ {
		numbers = validStorage.GetAllValidNumbers()
		numbersCount = len(numbers)
		assert.Equal(t, i, numbersCount)
		validStorage.AddValidNumber(validNumbers[i])
	}
}

func TestSeveralAddsOfNumberToValidStorage(t *testing.T) {
	var singleValidNumber = "0502030450"

	validStorage := NewValidRAMstorage()

	var numbers []string
	var numbersCount int
	n := 0

	for i := 0; i < 5; i++ {
		numbers = validStorage.GetAllValidNumbers()
		numbersCount = len(numbers)
		if i > 0 {
			n = 1
		}
		assert.Equal(t, n, numbersCount)
		validStorage.AddValidNumber(singleValidNumber)
	}
}
