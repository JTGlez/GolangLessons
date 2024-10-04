package main_test

import (
	taxes "class3go/01-taxes"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCalculateTaxesBelow50000(t *testing.T) {
	// arrange
	salary := 40000.0
	expectedTax := 0.0

	// act
	tax := taxes.CalculateTaxes(salary)

	// assert
	require.Equal(t, expectedTax, tax, "The tax for a salary below $50,000 should be 0")
}

func TestCalculateTaxesBetween50000And150000(t *testing.T) {
	// arrange
	salary := 100000.0
	expectedTax := salary * 0.17

	// act
	tax := taxes.CalculateTaxes(salary)

	// assert
	require.Equal(t, expectedTax, tax, "The tax for a salary between $50,000 and $150,000 should be 17%")
}

func TestCalculateTaxesAbove150000(t *testing.T) {
	// arrange
	salary := 200000.0
	expectedTax := salary * 0.27

	// act
	tax := taxes.CalculateTaxes(salary)

	// assert
	require.Equal(t, expectedTax, tax, "The tax for a salary above $150,000 should be 27%")
}
