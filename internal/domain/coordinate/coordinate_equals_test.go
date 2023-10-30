package coordinate_test

import (
	"mars_rover/internal/domain/coordinate"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEqualsBasedOnValues(t *testing.T) {
	aCoordinate := coordinate.New(1, 1)
	anEqualCoordinate := coordinate.New(1, 1)

	areTheSame := aCoordinate.Equals(*anEqualCoordinate)
	assert.True(t, areTheSame)
}

func TestNotEqualsBasedOnValues(t *testing.T) {
	aCoordinate := coordinate.New(1, 1)
	testCases := []struct {
		name string
		x    int
		y    int
	}{
		{
			name: "neither x nor y can differ",
			x:    2,
			y:    2,
		},
		{
			name: "x cannot differ",
			x:    2,
			y:    1,
		},
		{
			name: "y cannot differ",
			x:    1,
			y:    2,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			aDifferentCoordinate := coordinate.New(testCase.x, testCase.y)
			areTheSame := aCoordinate.Equals(*aDifferentCoordinate)
			assert.False(t, areTheSame)
		})
	}
}
