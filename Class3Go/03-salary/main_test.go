package main_test

import (
	salary "class3go/03-salary"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCalcSalaryCategoryC(t *testing.T) {
	// arrange
	minutes := 6000 // 100 hours
	category := "C"
	expectedSalary := float64(100 * 1000)

	// act
	sal, err := salary.CalcSalary(minutes, category)

	// assert
	require.NoError(t, err, "The error should be nil for valid input")
	require.Equal(t, expectedSalary, sal, "The salary for category C should be calculated correctly")
}

func TestCalcSalaryInvalidConditions(t *testing.T) {
	// Test for negative minutes
	{
		// arrange
		minutes := -100
		category := "A"

		// act
		_, err := salary.CalcSalary(minutes, category)

		// assert
		require.Error(t, err, "The function should return nil for negative minutes")
	}

	// Test for invalid category
	{
		// arrange
		minutes := 1000
		category := "X"

		// act
		_, err := salary.CalcSalary(minutes, category)

		// assert
		require.Error(t, err, "The function should return nil for an invalid category")
	}
}
