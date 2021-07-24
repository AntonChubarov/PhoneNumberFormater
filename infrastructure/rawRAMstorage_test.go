package infrastructure

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAllRawNumbers(t *testing.T) {
	rawStorage := NewRawRAMstorage()

	numbers := rawStorage.GetAllRawNumbers()
	numbersCount := len(numbers)

	assert.NotEqual(t, 0, numbersCount)

	for i := range numbers {
		assert.NotEqual(t, 0, len(numbers[i]))
	}
}