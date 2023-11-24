package planetMap_test

import (
	. "mars_rover/internal/domain/coordinate/absoluteCoordinate"
	. "mars_rover/internal/domain/obstacle"
	"mars_rover/internal/domain/obstacle/obstacles"
	. "mars_rover/internal/domain/obstacle/test"
	. "mars_rover/internal/domain/planet/test"
	"mars_rover/internal/domain/rover/planetMap"
	. "mars_rover/internal/domain/size"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestReportsCollisionWithMock(t *testing.T) {
	mockObstacle := new(MockObstacle)
	mockPlanet := new(MockPlanet)
	mockPlanet.On("Obstacles").Return(*obstacles.FromList([]Obstacle{mockObstacle}))
	mockPlanet.On("Size").Return(Size{})
	testMap := planetMap.OfPlanet(mockPlanet)

	mockObstacle.On("Occupies", mock.Anything).Return(true)

	didCollide := testMap.HasObstacleIn(AbsoluteCoordinate{})

	mockObstacle.AssertCalled(t, "Occupies", mock.Anything)
	assert.True(t, didCollide)
}

func TestReportsNoCollisionWithMock(t *testing.T) {
	mockObstacle := new(MockObstacle)
	mockPlanet := new(MockPlanet)
	mockPlanet.On("Obstacles").Return(*obstacles.FromList([]Obstacle{mockObstacle}))
	mockPlanet.On("Size").Return(Size{})
	testMap := planetMap.OfPlanet(mockPlanet)

	mockObstacle.On("Occupies", mock.Anything).Return(false)

	didCollide := testMap.HasObstacleIn(AbsoluteCoordinate{})

	mockObstacle.AssertCalled(t, "Occupies", mock.Anything)
	assert.False(t, didCollide)
}
