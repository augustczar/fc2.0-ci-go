package test

import (
	"testing"
	"math_test.go"
)

func TestSoma(t *testing.T) {
	total := Soma(15, 15)

	if total != 30 {
		t.Errorf("Sum result is invalid: Result: %d. Expected: %d", total, 30)
	}
}