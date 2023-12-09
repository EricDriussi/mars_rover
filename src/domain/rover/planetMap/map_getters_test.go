package planetMap_test

import (
	. "mars_rover/src/domain/obstacle"
	"mars_rover/src/domain/obstacle/obstacles"
	"mars_rover/src/domain/rover/planetMap"
	. "mars_rover/src/domain/size"
	. "mars_rover/src/test_helpers/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetWidthHeightAndObstacles(t *testing.T) {
	mockObstacle := new(MockObstacle)
	mockPlanet := new(MockPlanet)
	mockPlanet.On("Obstacles").Return(*obstacles.FromList([]Obstacle{mockObstacle}))
	mockPlanet.On("Size").Return(Size{})
	testMap := planetMap.OfPlanet(mockPlanet)

	planetSize := mockPlanet.Size()
	assert.Equal(t, testMap.Width(), planetSize.Width())
	assert.Equal(t, testMap.Height(), planetSize.Height())
	assert.Equal(t, testMap.Obstacles(), mockPlanet.Obstacles())
}
