package test

import (
	"mars_rover/internal/domain/obstacle/obstacles"
	"mars_rover/internal/domain/size"

	"github.com/stretchr/testify/mock"
)

type MockPlanet struct {
	mock.Mock
}

func (this *MockPlanet) Size() size.Size {
	args := this.Called()
	return args.Get(0).(size.Size)
}

func (this *MockPlanet) Obstacles() obstacles.Obstacles {
	args := this.Called()
	return args.Get(0).(obstacles.Obstacles)
}
