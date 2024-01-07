package coordinates_test

import (
	"github.com/stretchr/testify/assert"
	"mars_rover/src/domain/coordinate/absoluteCoordinate"
	. "mars_rover/src/domain/coordinate/absoluteCoordinate"
	"mars_rover/src/domain/coordinate/coordinates"
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

func TestDoesNotCreateWithNoCoordinates(t *testing.T) {
	coords, err := coordinates.New()

	assert.Nil(t, coords)
	assert.Error(t, err)
}

func TestAllCoordinatesAreWithinLimit(t *testing.T) {
	sizeLimit, _ := size.Square(4)
	coords, err := coordinates.New(
		*absoluteCoordinate.Build(1, 1),
		*absoluteCoordinate.Build(1, 2),
		*absoluteCoordinate.Build(1, 3),
	)

	assert.Nil(t, err)
	assert.False(t, coords.Overflow(*sizeLimit))
}

func TestSomeCoordinatesAreBeyondLimit(t *testing.T) {
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

func TestAnyCoordinateOccupiesAGivenCoordinate(t *testing.T) {
	x := rand.Int()
	y := rand.Int()
	aCoordinate := absoluteCoordinate.Build(x, y)
	coords, err := coordinates.New(*aCoordinate)

	assert.Nil(t, err)
	assert.True(t, coords.Contain(*absoluteCoordinate.Build(x, y)))
}

func TestNoCoordinateOccupiesAGivenCoordinate(t *testing.T) {
	coords, err := coordinates.New(*absoluteCoordinate.Build(1, 1))

	assert.Nil(t, err)
	assert.False(t, coords.Contain(*absoluteCoordinate.Build(1, 2)))
}

func TestAnyOfGivenCoordinatesAreOccupied(t *testing.T) {
	coords, err := coordinates.New(
		*absoluteCoordinate.Build(1, 1),
		*absoluteCoordinate.Build(1, 2),
	)
	assert.Nil(t, err)
	otherCoords, err := coordinates.New(
		*absoluteCoordinate.Build(1, 1),
		*absoluteCoordinate.Build(1, 3),
	)
	assert.Nil(t, err)

	assert.True(t, coords.ContainAnyOf(*otherCoords))
}

func TestNoneOfGivenCoordinatesAreOccupied(t *testing.T) {
	coords, err := coordinates.New(
		*absoluteCoordinate.Build(1, 1),
		*absoluteCoordinate.Build(1, 2),
	)
	assert.Nil(t, err)
	otherCoords, err := coordinates.New(
		*absoluteCoordinate.Build(1, 3),
		*absoluteCoordinate.Build(1, 4),
	)
	assert.Nil(t, err)

	assert.False(t, coords.ContainAnyOf(*otherCoords))
}

func TestDeterminesIfAllCoordinatesAreContiguous(t *testing.T) {
	coords, err := coordinates.New(
		*absoluteCoordinate.Build(1, 1),
		*absoluteCoordinate.Build(3, 3),
		*absoluteCoordinate.Build(1, 2),
		*absoluteCoordinate.Build(2, 2),
		*absoluteCoordinate.Build(2, 3),
	)

	assert.Nil(t, err)
	assert.True(t, coords.AreContiguous())
}

func TestDeterminesIfAnyCoordinateIsNonContiguous(t *testing.T) {
	coords, err := coordinates.New(
		*absoluteCoordinate.Build(1, 1),
		*absoluteCoordinate.Build(1, 3),
	)

	assert.Nil(t, err)
	assert.False(t, coords.AreContiguous())
}
