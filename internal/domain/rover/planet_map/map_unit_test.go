package planetMap_test

import (
	"mars_rover/internal/domain/coordinate"
	"mars_rover/internal/domain/obstacle"
	obstacleTest "mars_rover/internal/domain/obstacle/test"
	planetTest "mars_rover/internal/domain/planet/test"
	planetMap "mars_rover/internal/domain/rover/planet_map"
	"mars_rover/internal/domain/size"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestReportsCollisionWithMock(t *testing.T) {
	mockObstacle := new(obstacleTest.MockObstacle)
	mockPlanet := new(planetTest.MockPlanet)
	mockPlanet.On("Obstacles").Return([]obstacle.Obstacle{mockObstacle})
	mockPlanet.On("Size").Return(size.Size{})
	planetMap := planetMap.Of(mockPlanet)

	mockObstacle.On("Occupies", mock.Anything).Return(true)

	didCollide := planetMap.CheckCollision(coordinate.AbsoluteCoordinate{})

	mockObstacle.AssertCalled(t, "Occupies", mock.Anything)
	assert.True(t, didCollide)
}

func TestReportsNoCollisionWithMock(t *testing.T) {
	mockObstacle := new(obstacleTest.MockObstacle)
	mockPlanet := new(planetTest.MockPlanet)
	mockPlanet.On("Obstacles").Return([]obstacle.Obstacle{mockObstacle})
	mockPlanet.On("Size").Return(size.Size{})
	planetMap := planetMap.Of(mockPlanet)

	mockObstacle.On("Occupies", mock.Anything).Return(false)

	didCollide := planetMap.CheckCollision(coordinate.AbsoluteCoordinate{})

	mockObstacle.AssertCalled(t, "Occupies", mock.Anything)
	assert.False(t, didCollide)
}
