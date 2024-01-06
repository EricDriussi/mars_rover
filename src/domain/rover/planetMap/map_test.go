package planetMap_test

import (
	"mars_rover/src/domain/coordinate/absoluteCoordinate"
	. "mars_rover/src/domain/coordinate/absoluteCoordinate"
	"mars_rover/src/domain/obstacle/obstacles"
	"mars_rover/src/domain/rover/planetMap"
	"mars_rover/src/domain/size"
	. "mars_rover/src/domain/size"
	. "mars_rover/src/test_helpers/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestReportsThatObstacleIsInGivenCoordinate(t *testing.T) {
	mockObstacle := new(MockObstacle)
	mockPlanet := new(MockPlanet)
	mockPlanet.On("Obstacles").Return(*obstacles.FromList(mockObstacle))
	mockPlanet.On("Size").Return(Size{})
	mockObstacle.On("Occupies", mock.Anything).Return(true)
	testMap := planetMap.OfPlanet(mockPlanet)

	assert.True(t, testMap.HasObstacleIn(AbsoluteCoordinate{}))
}

func TestReportsThatObstacleIsNotInGivenCoordinate(t *testing.T) {
	mockObstacle := new(MockObstacle)
	mockPlanet := new(MockPlanet)
	mockPlanet.On("Obstacles").Return(*obstacles.FromList(mockObstacle))
	mockPlanet.On("Size").Return(Size{})
	mockObstacle.On("Occupies", mock.Anything).Return(false)
	testMap := planetMap.OfPlanet(mockPlanet)

	assert.False(t, testMap.HasObstacleIn(AbsoluteCoordinate{}))
}

func TestReportsCoordinateAsOutOfBoundsWhenOutsideThePlanet(t *testing.T) {
	planetSize, err := size.Square(5)
	assert.Nil(t, err)
	mockPlanet := new(MockPlanet)
	mockPlanet.On("Obstacles").Return(*obstacles.FromList())
	mockPlanet.On("Size").Return(*planetSize)
	testMap := planetMap.OfPlanet(mockPlanet)

	assert.False(t, testMap.IsOutOfBounds(*absoluteCoordinate.Build(0, 0)))
	assert.False(t, testMap.IsOutOfBounds(*absoluteCoordinate.Build(4, 4)))
	assert.True(t, testMap.IsOutOfBounds(*absoluteCoordinate.Build(6, 6)))
	assert.True(t, testMap.IsOutOfBounds(*absoluteCoordinate.Build(-1, -1)))
}
