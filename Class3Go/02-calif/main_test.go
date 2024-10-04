package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCalculateAverage_NoGrades(t *testing.T) {
	// arrange
	var grades []float64
	expectedError := "no grades provided"

	// act
	_, err := calculateAverage(grades...)

	// assert
	require.Error(t, err)
	require.EqualError(t, err, expectedError)
}

func TestCalculateAverage_InvalidGrade(t *testing.T) {
	// arrange
	grades := []float64{10, -5, 20}
	expectedError := "invalid grade"

	// act
	_, err := calculateAverage(grades...)

	// assert
	require.Error(t, err)
	require.EqualError(t, err, expectedError)
}

func TestCalculateAverage_ValidGrades(t *testing.T) {
	// arrange
	grades := []float64{10, 20, 30}
	expectedAverage := 20.0

	// act
	average, err := calculateAverage(grades...)

	// assert
	require.NoError(t, err)
	require.Equal(t, expectedAverage, average)
}
