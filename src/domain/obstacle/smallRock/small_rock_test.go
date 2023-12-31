package smallRock_test

import (
	"mars_rover/src/domain/coordinate/absoluteCoordinate"
	. "mars_rover/src/domain/coordinate/absoluteCoordinate"
	"mars_rover/src/domain/obstacle/smallRock"
	. "mars_rover/src/domain/obstacle/smallRock"
	"mars_rover/src/domain/size"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsWithinLimit(t *testing.T) {
	sizeLimit, _ := size.Square(2)

	for x := 0; x <= sizeLimit.Width(); x++ {
		for y := 0; y <= sizeLimit.Height(); y++ {
			rock := smallRock.In(*absoluteCoordinate.Build(x, y))

			assert.False(t, rock.IsBeyond(*sizeLimit))
		}
	}
}

func TestIsBeyondLimit(t *testing.T) {
	sizeLimit, _ := size.Square(3)
	testCases := []struct {
		name string
		rock SmallRock
	}{
		{
			name: "both out of bounds",
			rock: smallRock.In(*absoluteCoordinate.Build(4, 4)),
		},
		{
			name: "X out of bounds",
			rock: smallRock.In(*absoluteCoordinate.Build(4, 3)),
		},
		{
			name: "Y out of bounds",
			rock: smallRock.In(*absoluteCoordinate.Build(3, 4)),
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.True(t, testCase.rock.IsBeyond(*sizeLimit))
		})
	}
}

func TestOccupiesItsCoordinate(t *testing.T) {
	testCoordinate := *absoluteCoordinate.Build(rand.Int(), rand.Int())
	rock := smallRock.In(testCoordinate)

	assert.True(t, rock.Occupies(testCoordinate))
}

func TestDoesNotOccupyAnExternalCoordinate(t *testing.T) {
	testCoordinate := *absoluteCoordinate.Build(1, 1)
	rock := smallRock.In(testCoordinate)

	assert.False(t, rock.Occupies(*absoluteCoordinate.Build(1, 2)))
}

func TestListsOccupiedCoordinates(t *testing.T) {
	testCoordinate := *absoluteCoordinate.Build(2, 1)
	rock := smallRock.In(testCoordinate)

	assert.Equal(t, rock.Coordinates(), []AbsoluteCoordinate{testCoordinate})
}
