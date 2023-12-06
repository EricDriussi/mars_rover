package coordinates

import (
	"github.com/stretchr/testify/assert"
	"mars_rover/src/domain/coordinate/absoluteCoordinate"
	. "mars_rover/src/domain/coordinate/absoluteCoordinate"
	"mars_rover/src/domain/size"
	"math/rand"
	"testing"
)

func TestIsWithinLimit(t *testing.T) {
	sizeLimit, _ := size.Square(4)
	coordinates := New([]AbsoluteCoordinate{
		*absoluteCoordinate.From(1, 1),
		*absoluteCoordinate.From(1, 2),
		*absoluteCoordinate.From(1, 3),
	})

	assert.False(t, coordinates.GoBeyond(*sizeLimit))
}

func TestIsBeyondLimit(t *testing.T) {
	sizeLimit, _ := size.Square(3)
	testCases := []struct {
		name       string
		coordinate *AbsoluteCoordinate
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
			coordinates := New([]AbsoluteCoordinate{*testCase.coordinate})
			assert.True(t, coordinates.GoBeyond(*sizeLimit))
		})
	}
}

func TestOccupiesAGivenRandomCoordinate(t *testing.T) {
	testCoordinate := absoluteCoordinate.From(rand.Int(), rand.Int())
	coordinates := New([]AbsoluteCoordinate{*testCoordinate})

	assert.True(t, coordinates.Contain(*testCoordinate))
}

func TestDoesNotOccupyADifferentCoordinate(t *testing.T) {
	testCoordinate := absoluteCoordinate.From(1, 1)
	coordinates := New([]AbsoluteCoordinate{*testCoordinate})

	assert.False(t, coordinates.Contain(*absoluteCoordinate.From(1, 2)))
}

func TestGetsList(t *testing.T) {
	testCoordinates := []AbsoluteCoordinate{*From(1, 1)}
	coordinates := New(testCoordinates)

	assert.Equal(t, testCoordinates, coordinates.List())
}