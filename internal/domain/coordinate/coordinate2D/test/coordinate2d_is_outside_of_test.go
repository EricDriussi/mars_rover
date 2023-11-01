package coordinate2d_test

import (
	coordinate2d "mars_rover/internal/domain/coordinate/coordinate2D"
	"mars_rover/internal/domain/size"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsWithinAGivenSize(t *testing.T) {
	sizeLimit, _ := size.From(2, 2)
	for x := 0; x <= sizeLimit.Width; x++ {
		for y := 0; y <= sizeLimit.Height; y++ {
			validCoordinate := coordinate2d.New(x, y)

			assert.False(t, validCoordinate.IsOutsideOf(*sizeLimit))
		}
	}
}

func TestIsOutsideOfAGivenSize(t *testing.T) {
	sizeLimit, _ := size.From(3, 3)
	testCases := []struct {
		name              string
		invalidCoordinate *coordinate2d.Coordinate2D
	}{
		{
			name:              "both out of bounds",
			invalidCoordinate: coordinate2d.New(4, 4),
		},
		{
			name:              "X out of bounds",
			invalidCoordinate: coordinate2d.New(4, 3),
		},
		{
			name:              "Y out of bounds",
			invalidCoordinate: coordinate2d.New(3, 4),
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.True(t, testCase.invalidCoordinate.IsOutsideOf(*sizeLimit))
		})
	}
}