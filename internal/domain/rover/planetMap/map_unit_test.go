package planetMap_test

import (
	"mars_rover/internal/domain/coordinate/absoluteCoordinate"
	"mars_rover/internal/domain/obstacle"
	"mars_rover/internal/domain/obstacle/obstacles"
	obstacleTest "mars_rover/internal/domain/obstacle/test"
	planetTest "mars_rover/internal/domain/planet/test"
	"mars_rover/internal/domain/rover/planetMap"
	"mars_rover/internal/domain/size"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestReportsCollisionWithMock(t *testing.T) {
	mockObstacle := new(obstacleTest.MockObstacle)
	mockPlanet := new(planetTest.MockPlanet)
	mockPlanet.On("Obstacles").Return(*obstacles.New([]obstacle.Obstacle{mockObstacle}))
	mockPlanet.On("Size").Return(size.Size{})
	testMap := planetMap.Of(mockPlanet)

	mockObstacle.On("Occupies", mock.Anything).Return(true)

	didCollide := testMap.CollidesWithObstacle(absoluteCoordinate.AbsoluteCoordinate{})

	mockObstacle.AssertCalled(t, "Occupies", mock.Anything)
	assert.True(t, didCollide)
}

func TestReportsNoCollisionWithMock(t *testing.T) {
	mockObstacle := new(obstacleTest.MockObstacle)
	mockPlanet := new(planetTest.MockPlanet)
	mockPlanet.On("Obstacles").Return(*obstacles.New([]obstacle.Obstacle{mockObstacle}))
	mockPlanet.On("Size").Return(size.Size{})
	testMap := planetMap.Of(mockPlanet)

	mockObstacle.On("Occupies", mock.Anything).Return(false)

	didCollide := testMap.CollidesWithObstacle(absoluteCoordinate.AbsoluteCoordinate{})

	mockObstacle.AssertCalled(t, "Occupies", mock.Anything)
	assert.False(t, didCollide)
}
