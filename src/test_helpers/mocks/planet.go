package mocks

import (
	"github.com/stretchr/testify/assert"
	"mars_rover/src/domain/obstacle/obstacles"
	. "mars_rover/src/domain/obstacle/obstacles"
	. "mars_rover/src/domain/size"
	"testing"

	. "github.com/stretchr/testify/mock"
)

type MockPlanet struct {
	Mock
}

func (this *MockPlanet) Size() Size {
	args := this.Called()
	return args.Get(0).(Size)
}

func (this *MockPlanet) Obstacles() Obstacles {
	args := this.Called()
	return args.Get(0).(Obstacles)
}

func (this *MockPlanet) Color() string {
	args := this.Called()
	return args.String(0)
}

func PlanetWithNoObstaclesOfSize(t *testing.T, size int) *MockPlanet {
	mockPlanet := new(MockPlanet)
	testSize, err := Square(size)
	assert.Nil(t, err)
	mockPlanet.On("Size").Return(*testSize)
	testObstacles := obstacles.Empty()
	mockPlanet.On("Obstacles").Return(*testObstacles)
	return mockPlanet
}
