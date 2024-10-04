package main_test

import (
	calif "class3go/02-calif"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCalculateAverageNoGrades(t *testing.T) {
	// arrange
	var grades []float64
	expectedError := "no grades provided"

	// act
	_, err := calif.CalculateAverage(grades...)

	// assert
	require.Error(t, err)
	require.EqualError(t, err, expectedError)
}

func TestCalculateAverageInvalidGrade(t *testing.T) {
	// arrange
	grades := []float64{10, -5, 20}
	expectedError := "invalid grade"

	// act
	_, err := calif.CalculateAverage(grades...)

	// assert
	require.Error(t, err)
	require.EqualError(t, err, expectedError)
}

func TestCalculateAverageValidGrades(t *testing.T) {
	// arrange
	grades := []float64{10, 20, 30}
	expectedAverage := 20.0

	// act
	average, err := calif.CalculateAverage(grades...)

	// assert
	require.NoError(t, err)
	require.Equal(t, expectedAverage, average)
}

func TestCalculateAverageGradeOutOfRange(t *testing.T) {
	// arrange
	grades := []float64{5, -1, 11}
	expectedError := "invalid grade"

	// act
	_, err := calif.CalculateAverage(grades...)

	// assert
	require.Error(t, err)
	require.EqualError(t, err, expectedError)
}
