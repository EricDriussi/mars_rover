package small_rock_test

import (
	"mars_rover/internal/domain/coordinate/absoluteCoordinate"
	smallRock "mars_rover/internal/domain/obstacle/small_rock"
	"mars_rover/internal/domain/size"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsWithinLimit(t *testing.T) {
	sizeLimit := &size.Size{Width: 2, Height: 2}

	for x := 0; x <= sizeLimit.Width; x++ {
		for y := 0; y <= sizeLimit.Height; y++ {
			testCoordinate := absoluteCoordinate.From(x, y)
			rock := smallRock.In(*testCoordinate)

			assert.False(t, rock.IsBeyond(*sizeLimit))
		}
	}
}

func TestIsBeyondLimit(t *testing.T) {
	sizeLimit := &size.Size{Width: 3, Height: 3}
	testCases := []struct {
		name       string
		coordinate *absoluteCoordinate.AbsoluteCoordinate
	}{
		{
			name:       "both out of bounds",
			coordinate: absoluteCoordinate.From(4, 4),
		},
		{
			name:       "X out of bounds",
			coordinate: absoluteCoordinate.From(4, 3),
		},
		{
			name:       "Y out of bounds",
			coordinate: absoluteCoordinate.From(3, 4),
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			rock := smallRock.In(*testCase.coordinate)
			assert.True(t, rock.IsBeyond(*sizeLimit))
		})
	}
}

func TestOccupiesAGivenRandomCoordinate(t *testing.T) {
	testCoordinate := absoluteCoordinate.From(rand.Int(), rand.Int())
	rock := smallRock.In(*testCoordinate)

	assert.True(t, rock.Occupies(*testCoordinate))
}

func TestDoesNotOccupyADifferentCoordinate(t *testing.T) {
	testCoordinate := absoluteCoordinate.From(1, 1)
	rock := smallRock.In(*testCoordinate)

	assert.False(t, rock.Occupies(*absoluteCoordinate.From(1, 2)))
}
