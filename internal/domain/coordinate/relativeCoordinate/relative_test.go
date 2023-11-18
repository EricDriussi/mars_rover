package relativeCoordinate_test

import (
	"mars_rover/internal/domain/coordinate/relativeCoordinate"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRelativeFromOrthogonal(t *testing.T) {
	notRelative := relativeCoordinate.New(0, 0)
	testCases := []struct {
		name string
		x, y int
	}{
		{
			name: "up",
			x:    0,
			y:    1,
		},
		{
			name: "right",
			x:    1,
			y:    0,
		},
		{
			name: "down",
			x:    0,
			y:    -1,
		},
		{
			name: "left",
			x:    -1,
			y:    0,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			coord := relativeCoordinate.New(testCase.x, testCase.y)
			areTheSame := coord.X() == notRelative.X() && coord.Y() == notRelative.Y()
			assert.False(t, areTheSame)
		})
	}
}

func TestRelativeFromNonOrthogonal(t *testing.T) {
	notRelative := relativeCoordinate.New(0, 0)
	testCases := []struct {
		name string
		x, y int
	}{
		{
			name: "up-right",
			x:    1,
			y:    1,
		},
		{
			name: "up-left",
			x:    -1,
			y:    1,
		},
		{
			name: "down-right",
			x:    1,
			y:    -1,
		},
		{
			name: "down-left",
			x:    -1,
			y:    -1,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			coord := relativeCoordinate.New(testCase.x, testCase.y)
			areTheSame := coord.X() == notRelative.X() && coord.Y() == notRelative.Y()
			assert.True(t, areTheSame)
		})
	}
}
