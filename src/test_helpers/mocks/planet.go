package mocks

import (
	. "mars_rover/src/domain/obstacle/obstacles"
	. "mars_rover/src/domain/size"

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

func PlanetWithNoObstaclesOfSize(size int) *MockPlanet {
	mockPlanet := new(MockPlanet)
	testSize, _ := Square(size)
	mockPlanet.On("Size").Return(*testSize)
	testObstacles := FromList()
	mockPlanet.On("Obstacles").Return(*testObstacles)
	return mockPlanet
}
