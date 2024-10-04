package main_test

import (
	taxes "class3go/01-taxes"
	"errors"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCalculateTaxesSuite(t *testing.T) {

	suite := []struct {
		name        string
		salary      float64
		expectedTax float64
		err         error
	}{
		{
			name:        "should return a tax of 0",
			salary:      40000.0,
			expectedTax: 0.,
			err:         nil,
		},
		{
			name:        "should return a tax of 0 for salary below 50000",
			salary:      40000.0,
			expectedTax: 0.,
			err:         nil,
		},
		{
			name:        "should return a tax of 17% for salary between 50000 and 150000",
			salary:      100000.0,
			expectedTax: 100000.0 * 0.17,
			err:         nil,
		},
		{
			name:        "should return a tax of 27% for salary above 150000",
			salary:      200000.0,
			expectedTax: 200000.0 * 0.27,
			err:         nil,
		},
		{
			name:        "should return an error for invalid salary",
			salary:      -5000.0,
			expectedTax: 0.,
			err:         errors.New("invalid salary"),
		},
	}

	for _, c := range suite {
		t.Run(c.name, func(t *testing.T) {
			res, err := taxes.CalculateTaxes(c.salary)
			if c.err != nil {
				require.Error(t, err)
				require.Equal(t, c.err.Error(), err.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, c.expectedTax, res)
			}
		})
	}
}
