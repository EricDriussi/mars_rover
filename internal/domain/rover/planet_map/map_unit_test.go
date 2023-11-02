package planetMap_test

import (
	coordTest "mars_rover/internal/domain/coordinate/test"
	"mars_rover/internal/domain/obstacle"
	obstacleTest "mars_rover/internal/domain/obstacle/test"
	planetTest "mars_rover/internal/domain/planet/test"
	planetMap "mars_rover/internal/domain/rover/planet_map"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReportsCollisionWithMock(t *testing.T) {
	mockCoordinate := new(coordTest.MockCoordinate)
	mockObstacle := new(obstacleTest.MockObstacle)
	mockObstacle.On("Occupies", mockCoordinate).Return(true)
	mockplanet := new(planetTest.MockPlanet)
	mockplanet.SetObstacles([]obstacle.Obstacle{mockObstacle})

	planetMap := planetMap.Of(mockplanet)
	didCollide := planetMap.CheckCollision(mockCoordinate)

	mockObstacle.AssertCalled(t, "Occupies", mockCoordinate)
	assert.True(t, didCollide)
}

func TestReportsNoCollisionWithMock(t *testing.T) {
	mockCoordinate := new(coordTest.MockCoordinate)
	mockObstacle := new(obstacleTest.MockObstacle)
	mockObstacle.On("Occupies", mockCoordinate).Return(false)
	mockplanet := new(planetTest.MockPlanet)
	mockplanet.SetObstacles([]obstacle.Obstacle{mockObstacle})

	planetMap := planetMap.Of(mockplanet)
	didCollide := planetMap.CheckCollision(mockCoordinate)

	mockObstacle.AssertCalled(t, "Occupies", mockCoordinate)
	assert.False(t, didCollide)
}
