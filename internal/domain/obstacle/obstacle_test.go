package obstacle_test

import (
	"mars_rover/internal/domain/location"
	"mars_rover/internal/domain/obstacle"
	"mars_rover/internal/domain/size"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsWithinLimit(t *testing.T) {
	testLocation, _ := location.From(1, 2)
	testObstacle := obstacle.In(testLocation)
	sizeLimit, _ := size.From(5, 5)

	assert.True(t, testObstacle.IsWithinLimit(*sizeLimit))
}

func TestIsNotWithinLimit(t *testing.T) {
	sizeLimit, _ := size.From(3, 3)
	testCases := []struct {
		name string
		x    int
		y    int
	}{
		{
			name: "Both out of bounds",
			x:    4,
			y:    4,
		},
		{
			name: "X out of bounds",
			x:    4,
			y:    3,
		},
		{
			name: "Y out of bounds",
			x:    3,
			y:    4,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			outOfBoundsLocation, _ := location.From(testCase.x, testCase.y)
			outOfBoundsObstacle := obstacle.In(outOfBoundsLocation)

			assert.False(t, outOfBoundsObstacle.IsWithinLimit(*sizeLimit))
		})
	}
}
