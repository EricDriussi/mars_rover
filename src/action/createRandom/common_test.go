package randomCreator_test

import (
	"github.com/stretchr/testify/assert"
	. "mars_rover/src/action/createRandom"
	"mars_rover/src/domain/coordinate/absoluteCoordinate"
	"mars_rover/src/domain/obstacle/obstacles"
	"mars_rover/src/domain/size"
	. "mars_rover/src/test_helpers"
	"testing"
)

func TestGeneratesRandomColor(t *testing.T) {
	for i := 0; i < 10; i++ {
		color := RandomColor()
		assert.NotEmpty(t, color)
		assert.Contains(t, []string{"red", "green", "blue"}, color)
	}
}

func TestLoopsUntilRoverLanded(t *testing.T) {
	planet := SetupPlanetOfSizeWithObstacleIn(t, 2, *absoluteCoordinate.Build(1, 1))
	for i := 0; i < 10; i++ {
		rover := LoopUntilRoverLanded(planet)
		assert.NotNil(t, rover)
	}
}

func TestGeneratesRandomCoordinate(t *testing.T) {
	testSize, err := size.Square(3)
	assert.Nil(t, err)
	for i := 0; i < 10; i++ {
		coordinate := RandomCoordinateWithin(*testSize)
		assert.NotNil(t, coordinate)
	}
}

func TestLoopsUntilAbleToAddRandomObstacle(t *testing.T) {
	testSize, err := size.Square(3)
	assert.Nil(t, err)
	list := obstacles.Empty()
	for i := 0; i < 10; i++ {
		obs := LoopUntilAbleToAddRandomObstacle(*testSize, *list)
		assert.NotNil(t, obs)
		assert.Equal(t, 1, obs.Amount())
	}
}
