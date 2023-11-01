package planetMap_test

import (
	coordTest "mars_rover/internal/domain/coordinate/test"
	planetTest "mars_rover/internal/domain/planet/test"
	planetMap "mars_rover/internal/domain/rover/planet_map"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestReportsCollisionWithMock(t *testing.T) {
	mockplanet := new(planetTest.MockPlanet)
	mockCoordinate := new(coordTest.MockCoordinate)
	mockCoordinate.On("Equals", mock.Anything).Return(true)

	planetMap := planetMap.Of(mockplanet)
	didCollide := planetMap.CheckCollision(mockCoordinate)

	mockCoordinate.AssertCalled(t, "Equals", mock.Anything)
	assert.True(t, didCollide)
}

func TestReportsNoCollisionWithMock(t *testing.T) {
	mockplanet := new(planetTest.MockPlanet)
	mockCoordinate := new(coordTest.MockCoordinate)
	mockCoordinate.On("Equals", mock.Anything).Return(false)

	planetMap := planetMap.Of(mockplanet)
	didCollide := planetMap.CheckCollision(mockCoordinate)

	mockCoordinate.AssertCalled(t, "Equals", mock.Anything)
	assert.False(t, didCollide)
}
