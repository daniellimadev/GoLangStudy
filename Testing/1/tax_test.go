package tax

import (
	"testing"
)

func TestCalculateTax(t *testing.T) {
	amount := 500.0
	expect := 5.0

	result := CalculateTax(amount)

	if result != expect {
		t.Errorf("Expected %f but got %f", expect, result)
	}
}
