package coordinate_test

import (
	"mars_rover/internal/domain/coordinate"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEqualsBasedOnValues(t *testing.T) {
	aCoordinate := coordinate.New(1, 1)
	anEqualCoordinate := coordinate.New(1, 1)

	areTheSame := aCoordinate.Equals(anEqualCoordinate)
	assert.True(t, areTheSame)
}

func TestNotEqualsBasedOnValues(t *testing.T) {
	aCoordinate := coordinate.New(1, 1)
	testCases := []struct {
		name                string
		differentCoordinate *coordinate.Coordinate2D
	}{
		{
			name:                "neither X nor Y can differ",
			differentCoordinate: coordinate.New(2, 2),
		},
		{
			name:                "X cannot differ",
			differentCoordinate: coordinate.New(2, 1),
		},
		{
			name:                "Y cannot differ",
			differentCoordinate: coordinate.New(1, 2),
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.False(t, aCoordinate.Equals(testCase.differentCoordinate))
		})
	}
}
