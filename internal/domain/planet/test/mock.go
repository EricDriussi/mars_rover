package test

import (
	. "mars_rover/internal/domain/obstacle/obstacles"
	. "mars_rover/internal/domain/size"

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
