package test

import (
	"mars_rover/internal/domain/coordinate"
	"mars_rover/internal/domain/size"

	"github.com/stretchr/testify/mock"
)

type MockObstacle struct {
	mock.Mock
}

func (this *MockObstacle) IsBeyond(limit size.Size) bool {
	args := this.Called(limit)
	return args.Bool(0)
}

func (this *MockObstacle) Occupies(coord coordinate.AbsoluteCoordinate) bool {
	args := this.Called(coord)
	return args.Bool(0)
}

func (this *MockObstacle) Coordinates() coordinate.AbsoluteCoordinate {
	args := this.Called()
	return args.Get(0).(coordinate.AbsoluteCoordinate)
}
