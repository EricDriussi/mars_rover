package planetMap_test

import (
	coordTest "mars_rover/internal/domain/coordinate/test"
	"mars_rover/internal/domain/obstacle"
	obstacleTest "mars_rover/internal/domain/obstacle/test"
	planetTest "mars_rover/internal/domain/planet/test"
	planetMap "mars_rover/internal/domain/rover/planet_map"
	"mars_rover/internal/domain/size"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReportsCollisionWithMock(t *testing.T) {
	mockObstacle := new(obstacleTest.MockObstacle)
	mockPlanet := new(planetTest.MockPlanet)
	mockPlanet.On("Obstacles").Return([]obstacle.Obstacle{mockObstacle})
	mockPlanet.On("Size").Return(size.Size{})
	planetMap := planetMap.Of(mockPlanet)

	mockCoordinate := new(coordTest.MockCoordinate)
	mockObstacle.On("Occupies", mockCoordinate).Return(true)

	didCollide := planetMap.CheckCollision(mockCoordinate)

	mockObstacle.AssertCalled(t, "Occupies", mockCoordinate)
	assert.True(t, didCollide)
}

func TestReportsNoCollisionWithMock(t *testing.T) {
	mockObstacle := new(obstacleTest.MockObstacle)
	mockPlanet := new(planetTest.MockPlanet)
	mockPlanet.On("Obstacles").Return([]obstacle.Obstacle{mockObstacle})
	mockPlanet.On("Size").Return(size.Size{})
	planetMap := planetMap.Of(mockPlanet)

	mockCoordinate := new(coordTest.MockCoordinate)
	mockObstacle.On("Occupies", mockCoordinate).Return(false)

	didCollide := planetMap.CheckCollision(mockCoordinate)

	mockObstacle.AssertCalled(t, "Occupies", mockCoordinate)
	assert.False(t, didCollide)
}
