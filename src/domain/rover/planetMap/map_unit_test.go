package planetMap_test

import (
	"mars_rover/src/domain/coordinate/absoluteCoordinate"
	. "mars_rover/src/domain/coordinate/absoluteCoordinate"
	. "mars_rover/src/domain/obstacle"
	"mars_rover/src/domain/obstacle/obstacles"
	. "mars_rover/src/domain/obstacle/test"
	. "mars_rover/src/domain/planet/test"
	"mars_rover/src/domain/rover/planetMap"
	"mars_rover/src/domain/size"
	. "mars_rover/src/domain/size"
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

func TestIsOutOfBoundsWithMock(t *testing.T) {
	planetSize, err := size.Square(5)
	assert.Nil(t, err)
	mockPlanet := new(MockPlanet)
	mockPlanet.On("Obstacles").Return(*obstacles.FromList([]Obstacle{}))
	mockPlanet.On("Size").Return(*planetSize)
	testMap := planetMap.OfPlanet(mockPlanet)

	assert.False(t, testMap.IsOutOfBounds(*absoluteCoordinate.From(0, 0)))
	assert.False(t, testMap.IsOutOfBounds(*absoluteCoordinate.From(4, 4)))
	assert.True(t, testMap.IsOutOfBounds(*absoluteCoordinate.From(6, 6)))
	assert.True(t, testMap.IsOutOfBounds(*absoluteCoordinate.From(-1, -1)))
}
