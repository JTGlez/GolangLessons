package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMinFunc(t *testing.T) {
	// arrange
	grades := []float64{10, 20, 30, 5, 15}
	expectedMin := 5.0

	// act
	minimum := minFunc(grades...)

	// assert
	require.Equal(t, expectedMin, minimum, "The minimum value should be calculated correctly")
}

func TestMaxFunc(t *testing.T) {
	// arrange
	grades := []float64{10, 20, 30, 5, 15}
	expectedMax := 30.0

	// act
	maximum := maxFunc(grades...)

	// assert
	require.Equal(t, expectedMax, maximum, "The maximum value should be calculated correctly")
}

func TestAverageFunc(t *testing.T) {
	// arrange
	grades := []float64{10, 20, 30, 5, 15}
	expectedAverage := 16.0

	// act
	average := averageFunc(grades...)

	// assert
	require.Equal(t, expectedAverage, average, "The average value should be calculated correctly")
}
