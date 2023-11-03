package small_rock_test

import (
	"mars_rover/internal/domain/coordinate"
	smallRock "mars_rover/internal/domain/obstacle/small_rock"
	"mars_rover/internal/domain/size"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsWithinLimit(t *testing.T) {
	sizeLimit := &size.Size{Width: 2, Height: 2}

	for x := 0; x <= sizeLimit.Width; x++ {
		for y := 0; y <= sizeLimit.Height; y++ {
			testCoordinate := coordinate.New(x, y)
			rock := smallRock.In(*testCoordinate)

			assert.False(t, rock.IsBeyond(*sizeLimit))
		}
	}
}

func TestIsBeyondLimit(t *testing.T) {
	sizeLimit := &size.Size{Width: 3, Height: 3}
	testCases := []struct {
		name       string
		coordinate *coordinate.Coordinate
	}{
		{
			name:       "both out of bounds",
			coordinate: coordinate.New(4, 4),
		},
		{
			name:       "X out of bounds",
			coordinate: coordinate.New(4, 3),
		},
		{
			name:       "Y out of bounds",
			coordinate: coordinate.New(3, 4),
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			rock := smallRock.In(*testCase.coordinate)
			assert.True(t, rock.IsBeyond(*sizeLimit))
		})
	}
}
