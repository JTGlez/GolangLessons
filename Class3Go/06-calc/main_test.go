package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAdd(t *testing.T) {
	// arrange
	val1, val2 := 5.0, 3.0
	expected := 8.0

	// act
	result := add(val1, val2)

	// assert
	require.Equal(t, expected, result, "The addition result should be correct")
}

func TestSubtract(t *testing.T) {
	// arrange
	val1, val2 := 5.0, 3.0
	expected := 2.0

	// act
	result := subtract(val1, val2)

	// assert
	require.Equal(t, expected, result, "The subtraction result should be correct")
}

func TestMultiply(t *testing.T) {
	// arrange
	val1, val2 := 5.0, 3.0
	expected := 15.0

	// act
	result := multiply(val1, val2)

	// assert
	require.Equal(t, expected, result, "The multiplication result should be correct")
}

func TestDivide(t *testing.T) {
	// arrange
	val1, val2 := 6.0, 3.0
	expected := 2.0

	// act
	result := divide(val1, val2)

	// assert
	require.Equal(t, expected, result, "The division result should be correct")
}

func TestDivideByZero(t *testing.T) {
	// arrange
	val1, val2 := 6.0, 0.0
	expected := 0.0

	// act
	result := divide(val1, val2)

	// assert
	require.Equal(t, expected, result, "The division by zero should return 0")
}
