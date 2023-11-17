package bigRock_test

import (
	"github.com/stretchr/testify/assert"
	"mars_rover/internal/domain/coordinate/absoluteCoordinate"
	"mars_rover/internal/domain/obstacle/bigRock"
	"mars_rover/internal/domain/size"
	"math/rand"
	"testing"
)

func TestIsWithinLimit(t *testing.T) {
	sizeLimit, _ := size.Square(4)
	coordinates := []absoluteCoordinate.AbsoluteCoordinate{
		*absoluteCoordinate.From(1, 1),
		*absoluteCoordinate.From(1, 2),
		*absoluteCoordinate.From(1, 3),
	}

	rock := bigRock.In(coordinates)

	assert.False(t, rock.IsBeyond(*sizeLimit))
}

func TestIsBeyondLimit(t *testing.T) {
	sizeLimit, _ := size.Square(3)
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
			rock := bigRock.In([]absoluteCoordinate.AbsoluteCoordinate{*testCase.coordinate})
			assert.True(t, rock.IsBeyond(*sizeLimit))
		})
	}
}

func TestOccupiesAGivenRandomCoordinate(t *testing.T) {
	testCoordinate := absoluteCoordinate.From(rand.Int(), rand.Int())
	rock := bigRock.In([]absoluteCoordinate.AbsoluteCoordinate{*testCoordinate})

	assert.True(t, rock.Occupies(*testCoordinate))
}

func TestDoesNotOccupyADifferentCoordinate(t *testing.T) {
	testCoordinate := absoluteCoordinate.From(1, 1)
	rock := bigRock.In([]absoluteCoordinate.AbsoluteCoordinate{*testCoordinate})

	assert.False(t, rock.Occupies(*absoluteCoordinate.From(1, 2)))
}
