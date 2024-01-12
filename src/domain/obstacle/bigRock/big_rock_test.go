package bigRock_test

import (
	"github.com/stretchr/testify/assert"
	"mars_rover/src/domain/coordinate/absoluteCoordinate"
	. "mars_rover/src/domain/coordinate/absoluteCoordinate"
	"mars_rover/src/domain/coordinate/coordinates"
	"mars_rover/src/domain/obstacle/bigRock"
	"mars_rover/src/domain/size"
	"math/rand"
	"testing"
)

func TestBuildsWithWithinTwoAndFiveCoordinates(t *testing.T) {
	testCases := []struct {
		name        string
		coordinates []AbsoluteCoordinate
	}{
		{
			name: "two coordinates",
			coordinates: []AbsoluteCoordinate{
				*absoluteCoordinate.Build(1, 1),
				*absoluteCoordinate.Build(1, 2),
			},
		},
		{
			name: "three coordinates",
			coordinates: []AbsoluteCoordinate{
				*absoluteCoordinate.Build(1, 1),
				*absoluteCoordinate.Build(1, 2),
				*absoluteCoordinate.Build(1, 3),
			},
		},
		{
			name: "four coordinates",
			coordinates: []AbsoluteCoordinate{
				*absoluteCoordinate.Build(1, 1),
				*absoluteCoordinate.Build(1, 2),
				*absoluteCoordinate.Build(1, 3),
				*absoluteCoordinate.Build(1, 4),
			},
		},
		{
			name: "five coordinates",
			coordinates: []AbsoluteCoordinate{
				*absoluteCoordinate.Build(1, 1),
				*absoluteCoordinate.Build(1, 2),
				*absoluteCoordinate.Build(1, 3),
				*absoluteCoordinate.Build(1, 4),
				*absoluteCoordinate.Build(1, 5),
			},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			coords, err := coordinates.New(testCase.coordinates...)
			assert.Nil(t, err)

			_, err = bigRock.In(*coords)

			assert.Nil(t, err)
		})
	}
}

func TestDoesNotBuildWithLessThanTwoOrMoreThanEightCoordinates(t *testing.T) {
	testCases := []struct {
		name        string
		coordinates []AbsoluteCoordinate
	}{
		{
			name: "one coordinate",
			coordinates: []AbsoluteCoordinate{
				*absoluteCoordinate.Build(1, 1),
			},
		},
		{
			name: "nine coordinates",
			coordinates: []AbsoluteCoordinate{
				*absoluteCoordinate.Build(1, 1),
				*absoluteCoordinate.Build(1, 2),
				*absoluteCoordinate.Build(1, 3),
				*absoluteCoordinate.Build(1, 4),
				*absoluteCoordinate.Build(1, 5),
				*absoluteCoordinate.Build(1, 6),
				*absoluteCoordinate.Build(1, 7),
				*absoluteCoordinate.Build(1, 8),
				*absoluteCoordinate.Build(1, 9),
			},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			coords, err := coordinates.New(testCase.coordinates...)
			assert.Nil(t, err)

			_, err = bigRock.In(*coords)

			assert.Error(t, err)
		})
	}
}

func TestIsWithinLimit(t *testing.T) {
	coords, err := coordinates.New(
		*absoluteCoordinate.Build(1, 1),
		*absoluteCoordinate.Build(1, 2),
		*absoluteCoordinate.Build(1, 3),
	)
	assert.Nil(t, err)
	rock, err := bigRock.In(*coords)
	assert.Nil(t, err)
	sizeLimit, err := size.Square(4)
	assert.Nil(t, err)

	assert.False(t, rock.IsBeyond(*sizeLimit))
}

func TestIsBeyondLimit(t *testing.T) {
	sizeLimit, err := size.Square(3)
	assert.Nil(t, err)
	coordinateWithinLimit := absoluteCoordinate.Build(3, 3)
	testCases := []struct {
		name       string
		coordinate *AbsoluteCoordinate
	}{
		{
			name:       "X out of bounds",
			coordinate: absoluteCoordinate.Build(4, 3),
		},
		{
			name:       "Y out of bounds",
			coordinate: absoluteCoordinate.Build(3, 4),
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			coords, err := coordinates.New(*coordinateWithinLimit, *testCase.coordinate)
			assert.Nil(t, err)
			rock, err := bigRock.In(*coords)
			assert.Nil(t, err)

			assert.True(t, rock.IsBeyond(*sizeLimit))
		})
	}
}

func TestOccupiesItsCoordinates(t *testing.T) {
	x := rand.Int()
	y := rand.Int()
	coordinate1 := absoluteCoordinate.Build(x, y)
	coordinate2 := absoluteCoordinate.Build(x, y+1)
	coords, err := coordinates.New(*coordinate1, *coordinate2)
	assert.Nil(t, err)
	rock, err := bigRock.In(*coords)
	assert.Nil(t, err)

	assert.True(t, rock.Occupies(*coordinate1))
	assert.True(t, rock.Occupies(*coordinate2))
}

func TestDoesNotOccupyAnExternalCoordinate(t *testing.T) {
	coords, err := coordinates.New(
		*absoluteCoordinate.Build(1, 1),
		*absoluteCoordinate.Build(1, 2),
	)
	assert.Nil(t, err)
	rock, err := bigRock.In(*coords)
	assert.Nil(t, err)

	assert.False(t, rock.Occupies(*absoluteCoordinate.Build(1, 3)))
}
