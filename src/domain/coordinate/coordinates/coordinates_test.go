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
	coords, err := coordinates.New(
		*absoluteCoordinate.Build(1, 1),
		*absoluteCoordinate.Build(1, 1),
		*absoluteCoordinate.Build(1, 3),
	)

	assert.Nil(t, err)
	assert.Len(t, coords.List(), 2)
}

func TestAreWithinLimit(t *testing.T) {
	sizeLimit, _ := size.Square(4)
	coords, err := coordinates.New(
		*absoluteCoordinate.Build(1, 1),
		*absoluteCoordinate.Build(1, 2),
		*absoluteCoordinate.Build(1, 3),
	)

	assert.Nil(t, err)
	assert.False(t, coords.Overflow(*sizeLimit))
}

func TestDoesNotCreateWithEmptyList(t *testing.T) {
	coords, err := coordinates.New()

	assert.Nil(t, coords)
	assert.Error(t, err)
}

func TestAreBeyondLimit(t *testing.T) {
	sizeLimit, _ := size.Square(3)
	testCases := []struct {
		name       string
		coordinate AbsoluteCoordinate
	}{
		{
			name:       "both out of bounds",
			coordinate: *absoluteCoordinate.Build(4, 4),
		},
		{
			name:       "X out of bounds",
			coordinate: *absoluteCoordinate.Build(4, 3),
		},
		{
			name:       "Y out of bounds",
			coordinate: *absoluteCoordinate.Build(3, 4),
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			coords, err := coordinates.New(testCase.coordinate)
			assert.Nil(t, err)
			assert.True(t, coords.Overflow(*sizeLimit))
		})
	}
}

func TestOccupiesAnIncludedCoordinate(t *testing.T) {
	testCoordinate := absoluteCoordinate.Build(rand.Int(), rand.Int())
	coords, err := New(*testCoordinate)

	assert.Nil(t, err)
	assert.True(t, coords.Contain(*testCoordinate))
}

func TestDoesNotOccupyANotIncludedCoordinate(t *testing.T) {
	testCoordinate := absoluteCoordinate.Build(1, 1)
	coords, err := New(*testCoordinate)

	assert.Nil(t, err)
	assert.False(t, coords.Contain(*absoluteCoordinate.Build(1, 2)))
}

func TestListsContainedCoordinates(t *testing.T) {
	testCoordinate := absoluteCoordinate.Build(rand.Int(), rand.Int())
	coords, err := New(*testCoordinate)

	assert.Nil(t, err)
	assert.Equal(t, []AbsoluteCoordinate{*testCoordinate}, coords.List())
}

func TestDeterminesIfContiguous(t *testing.T) {
	coords, err := New(
		*absoluteCoordinate.Build(1, 1),
		*absoluteCoordinate.Build(3, 3),
		*absoluteCoordinate.Build(1, 2),
		*absoluteCoordinate.Build(2, 2),
		*absoluteCoordinate.Build(2, 3),
	)

	assert.Nil(t, err)
	assert.True(t, coords.Contiguous())
}

func TestDeterminesIfNotContiguous(t *testing.T) {
	coords, err := New(
		*absoluteCoordinate.Build(1, 1),
		*absoluteCoordinate.Build(1, 3),
	)

	assert.Nil(t, err)
	assert.False(t, coords.Contiguous())
}
