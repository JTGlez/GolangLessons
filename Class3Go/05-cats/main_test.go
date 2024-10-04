package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCalcFoodForSpiders(t *testing.T) {
	// arrange
	spider := Animal{name: "Spider", kg: 0.15}
	count := 6
	expectedFood := 0.9
	tolerance := 0.0001

	// act
	calcFood := spider.animal()
	totalFood := calcFood(count)

	// assert
	require.InEpsilon(t, expectedFood, totalFood, tolerance, "The total food for spiders should be calculated correctly")
}
