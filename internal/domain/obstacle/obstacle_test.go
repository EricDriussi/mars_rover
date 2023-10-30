package obstacle_test

import (
	"mars_rover/internal/domain/coordinate"
	"mars_rover/internal/domain/obstacle"
	"mars_rover/internal/domain/size"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsWithinLimit(t *testing.T) {
	testCoordinate := coordinate.New(1, 2)
	testObstacle := obstacle.In(testCoordinate)
	sizeLimit, _ := size.From(5, 5)

	assert.True(t, testObstacle.IsWithinLimit(*sizeLimit))
}

func TestIsNotWithinLimit(t *testing.T) {
	sizeLimit, _ := size.From(3, 3)
	testCases := []struct {
		name       string
		coordinate *coordinate.Coordinate
	}{
		{
			name:       "Both out of bounds",
			coordinate: coordinate.New(4, 4),
		},
		{
			name:       "X out of bounds",
			coordinate: coordinate.New(4, 3),
		},
		{
			name:       "Y out of bounds",
			coordinate: coordinate.New(3, 4),
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			outOfBoundsObstacle := obstacle.In(testCase.coordinate)

			assert.False(t, outOfBoundsObstacle.IsWithinLimit(*sizeLimit))
		})
	}
}
