package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCalcSalary_CategoryC(t *testing.T) {
	// arrange
	minutes := 6000 // 100 hours
	category := "C"
	expectedSalary := float64(100 * 1000) // base salary, no bonus

	// act
	salary := calcSalary(minutes, category)

	// assert
	require.Equal(t, expectedSalary, salary, "The salary for category C should be calculated correctly")
}
