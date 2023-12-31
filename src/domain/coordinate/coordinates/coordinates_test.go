package coordinates_test

import (
	"github.com/stretchr/testify/assert"
	"mars_rover/src/domain/coordinate/absoluteCoordinate"
	. "mars_rover/src/domain/coordinate/absoluteCoordinate"
	"mars_rover/src/domain/coordinate/coordinates"
	. "mars_rover/src/domain/coordinate/coordinates"
	"mars_rover/src/domain/size"
	"math/rand"
	"testing"
)

func TestFiltersOutDuplicateCoordinates(t *testing.T) {
	coords := coordinates.New(
		*absoluteCoordinate.Build(1, 1),
		*absoluteCoordinate.Build(1, 1),
		*absoluteCoordinate.Build(1, 3),
	)

	assert.Len(t, coords.List(), 2)
}

func TestAreWithinLimit(t *testing.T) {
	sizeLimit, _ := size.Square(4)
	coords := coordinates.New(
		*absoluteCoordinate.Build(1, 1),
		*absoluteCoordinate.Build(1, 2),
		*absoluteCoordinate.Build(1, 3),
	)

	assert.False(t, coords.Overflow(*sizeLimit))
}

func TestAreBeyondLimit(t *testing.T) {
	sizeLimit, _ := size.Square(3)
	testCases := []struct {
		name        string
		coordinates *Coordinates
	}{
		{
			name:        "both out of bounds",
			coordinates: coordinates.New(*absoluteCoordinate.Build(4, 4)),
		},
		{
			name:        "X out of bounds",
			coordinates: coordinates.New(*absoluteCoordinate.Build(4, 3)),
		},
		{
			name:        "Y out of bounds",
			coordinates: coordinates.New(*absoluteCoordinate.Build(3, 4)),
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.True(t, testCase.coordinates.Overflow(*sizeLimit))
		})
	}
}

func TestOccupiesAnIncludedCoordinate(t *testing.T) {
	testCoordinate := absoluteCoordinate.Build(rand.Int(), rand.Int())
	coords := New(*testCoordinate)

	assert.True(t, coords.Contain(*testCoordinate))
}

func TestDoesNotOccupyANotIncludedCoordinate(t *testing.T) {
	testCoordinate := absoluteCoordinate.Build(1, 1)
	coords := New(*testCoordinate)

	assert.False(t, coords.Contain(*absoluteCoordinate.Build(1, 2)))
}

func TestListsContainedCoordinates(t *testing.T) {
	testCoordinate := absoluteCoordinate.Build(rand.Int(), rand.Int())
	coords := New(*testCoordinate)

	assert.Equal(t, []AbsoluteCoordinate{*testCoordinate}, coords.List())
}
