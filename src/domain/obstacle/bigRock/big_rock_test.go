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

func TestIsWithinLimit(t *testing.T) {
	sizeLimit, _ := size.Square(4)
	coordinates := []AbsoluteCoordinate{
		*absoluteCoordinate.Build(1, 1),
		*absoluteCoordinate.Build(1, 2),
		*absoluteCoordinate.Build(1, 3),
	}

	rock := bigRock.In(coordinates)

	assert.False(t, rock.IsBeyond(*sizeLimit))
}

func TestIsBeyondLimit(t *testing.T) {
	sizeLimit, _ := size.Square(3)
	testCases := []struct {
		name       string
		coordinate *AbsoluteCoordinate
	}{
		{
			name:       "both out of bounds",
			coordinate: absoluteCoordinate.Build(4, 4),
		},
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
			rock := bigRock.In([]AbsoluteCoordinate{*testCase.coordinate})
			assert.True(t, rock.IsBeyond(*sizeLimit))
		})
	}
}

func TestOccupiesAGivenRandomCoordinate(t *testing.T) {
	testCoordinate := absoluteCoordinate.Build(rand.Int(), rand.Int())
	rock := bigRock.In([]AbsoluteCoordinate{*testCoordinate})

	assert.True(t, rock.Occupies(*testCoordinate))
}

func TestDoesNotOccupyADifferentCoordinate(t *testing.T) {
	testCoordinate := absoluteCoordinate.Build(1, 1)
	rock := bigRock.In([]AbsoluteCoordinate{*testCoordinate})

	assert.False(t, rock.Occupies(*absoluteCoordinate.Build(1, 2)))
}

func TestGetCoordinates(t *testing.T) {
	testCoordinate := []AbsoluteCoordinate{*absoluteCoordinate.Build(2, 1)}
	rock := bigRock.In(testCoordinate)

	assert.Equal(t, rock.Coordinates(), testCoordinate)
}
