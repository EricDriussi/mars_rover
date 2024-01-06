package mocks

import (
	. "mars_rover/src/domain/coordinate/absoluteCoordinate"
	. "mars_rover/src/domain/coordinate/coordinates"
	. "mars_rover/src/domain/size"

	. "github.com/stretchr/testify/mock"
)

type MockObstacle struct {
	Mock
}

func (this *MockObstacle) IsBeyond(limit Size) bool {
	args := this.Called(limit)
	return args.Bool(0)
}

func (this *MockObstacle) Occupies(coord AbsoluteCoordinate) bool {
	args := this.Called(coord)
	return args.Bool(0)
}

func (this *MockObstacle) Coordinates() Coordinates {
	args := this.Called()
	return args.Get(0).(Coordinates)
}
