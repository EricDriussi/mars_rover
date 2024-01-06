package relativeCoordinate_test

import (
	"github.com/stretchr/testify/assert"
	. "mars_rover/src/domain/coordinate/relativeCoordinate"
	"testing"
)

func TestIsAlwaysOrthogonal(t *testing.T) {
	testCases := []struct {
		name        string
		constructor func() *RelativeCoordinate
		x, y        int
	}{
		{
			name:        "up",
			constructor: Up,
			x:           0,
			y:           1,
		},
		{
			name:        "right",
			constructor: Right,
			x:           1,
			y:           0,
		},
		{
			name:        "down",
			constructor: Down,
			x:           0,
			y:           -1,
		},
		{
			name:        "left",
			constructor: Left,
			x:           -1,
			y:           0,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			coordinate := testCase.constructor()
			assert.Equal(t, testCase.x, coordinate.X())
			assert.Equal(t, testCase.y, coordinate.Y())
		})
	}
}
