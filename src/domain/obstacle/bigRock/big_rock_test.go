package bigRock_test

import (
	"github.com/stretchr/testify/assert"
	"mars_rover/src/domain/coordinate/absoluteCoordinate"
	. "mars_rover/src/domain/coordinate/absoluteCoordinate"
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
			_, err := bigRock.In(testCase.coordinates...)
			assert.Nil(t, err)
		})
	}
}

func TestDoesNotBuildWithLessThanTwoOrMoreThanFiveCoordinates(t *testing.T) {
	testCases := []struct {
		name        string
		coordinates []AbsoluteCoordinate
	}{
		{
			name:        "no coordinates",
			coordinates: []AbsoluteCoordinate{},
		},
		{
			name: "one coordinate",
			coordinates: []AbsoluteCoordinate{
				*absoluteCoordinate.Build(1, 1),
			},
		},
		{
			name: "six coordinates",
			coordinates: []AbsoluteCoordinate{
				*absoluteCoordinate.Build(1, 1),
				*absoluteCoordinate.Build(1, 2),
				*absoluteCoordinate.Build(1, 3),
				*absoluteCoordinate.Build(1, 4),
				*absoluteCoordinate.Build(1, 5),
				*absoluteCoordinate.Build(1, 6),
			},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			_, err := bigRock.In(testCase.coordinates...)
			assert.NotNil(t, err)
		})
	}
}

func TestIsWithinLimit(t *testing.T) {
	sizeLimit, _ := size.Square(4)
	rock, err := bigRock.In(
		*Build(1, 1),
		*Build(1, 2),
		*Build(1, 3),
	)
	assert.Nil(t, err)

	assert.False(t, rock.IsBeyond(*sizeLimit))
}

func TestIsBeyondLimit(t *testing.T) {
	sizeLimit, _ := size.Square(3)
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
			rock, err := bigRock.In(*coordinateWithinLimit, *testCase.coordinate)
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
	rock, err := bigRock.In(*coordinate1, *coordinate2)
	assert.Nil(t, err)

	assert.True(t, rock.Occupies(*coordinate1))
	assert.True(t, rock.Occupies(*coordinate2))
}

func TestDoesNotOccupyAnExternalCoordinate(t *testing.T) {
	rock, err := bigRock.In(
		*absoluteCoordinate.Build(1, 1),
		*absoluteCoordinate.Build(1, 2),
	)
	assert.Nil(t, err)

	assert.False(t, rock.Occupies(*absoluteCoordinate.Build(1, 3)))
}

func TestListsOccupiedCoordinates(t *testing.T) {
	testCoordinate := []AbsoluteCoordinate{*absoluteCoordinate.Build(2, 1), *absoluteCoordinate.Build(2, 2)}
	rock, err := bigRock.In(testCoordinate...)
	assert.Nil(t, err)

	assert.Equal(t, rock.Coordinates(), testCoordinate)
}
