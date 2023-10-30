package coordinate_test

import (
	"mars_rover/internal/domain/coordinate"
	"mars_rover/internal/domain/size"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsWithinAGivenSize(t *testing.T) {
	sizeLimit, _ := size.From(2, 2)
	for x := 0; x <= sizeLimit.Width; x++ {
		for y := 0; y <= sizeLimit.Height; y++ {
			validCoordinate := coordinate.New(x, y)

			assert.True(t, validCoordinate.IsWithin(*sizeLimit))
		}
	}
}

func TestIsNotWithinAGivenSize(t *testing.T) {
	sizeLimit, _ := size.From(3, 3)
	testCases := []struct {
		name string
		x    int
		y    int
	}{
		{
			name: "Both out of bounds",
			x:    4,
			y:    4,
		},
		{
			name: "X out of bounds",
			x:    4,
			y:    3,
		},
		{
			name: "Y out of bounds",
			x:    3,
			y:    4,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			outOfBoundsCoordinate := coordinate.New(testCase.x, testCase.y)

			assert.False(t, outOfBoundsCoordinate.IsWithin(*sizeLimit))
		})
	}
}
