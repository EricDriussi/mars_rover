package test

import (
	"mars_rover/internal/domain/coordinate"
	"mars_rover/internal/domain/obstacle"
	"mars_rover/internal/domain/size"

	"github.com/stretchr/testify/mock"
)

type MockPlanet struct {
	mock.Mock
	coord     coordinate.Coordinate
	obstacles []obstacle.Obstacle
}

func (this *MockPlanet) Size() size.Size {
	args := this.Called()
	return args.Get(0).(size.Size)
}

func (this *MockPlanet) Obstacles() []obstacle.Obstacle {
	args := this.Called()
	return args.Get(0).([]obstacle.Obstacle)
}
