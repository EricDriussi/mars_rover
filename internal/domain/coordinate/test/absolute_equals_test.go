package coordinate_test

import (
	"mars_rover/internal/domain/coordinate"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEqualsBasedOnValues(t *testing.T) {
	aCoordinate := coordinate.NewAbsolute(1, 1)
	anEqualCoordinate := coordinate.NewAbsolute(1, 1)

	areTheSame := aCoordinate.Equals(anEqualCoordinate)
	assert.True(t, areTheSame)
}

func TestNotEqualsBasedOnValues(t *testing.T) {
	aCoordinate := coordinate.NewAbsolute(1, 1)
	testCases := []struct {
		name                string
		differentCoordinate *coordinate.AbsoluteCoordinate
	}{
		{
			name:                "neither X nor Y can differ",
			differentCoordinate: coordinate.NewAbsolute(2, 2),
		},
		{
			name:                "X cannot differ",
			differentCoordinate: coordinate.NewAbsolute(2, 1),
		},
		{
			name:                "Y cannot differ",
			differentCoordinate: coordinate.NewAbsolute(1, 2),
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			areTheSame := aCoordinate.Equals(testCase.differentCoordinate)
			assert.False(t, areTheSame)
		})
	}
}
