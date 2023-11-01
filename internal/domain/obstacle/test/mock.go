package test

import (
	"mars_rover/internal/domain/coordinate"
	"mars_rover/internal/domain/size"

	"github.com/stretchr/testify/mock"
)

type MockObstacle struct {
	mock.Mock
	coord coordinate.Coordinate
}

func (this *MockObstacle) IsBeyond(limit size.Size) bool {
	args := this.Called(limit)
	return args.Bool(0)
}

func (this *MockObstacle) Coordinate() coordinate.Coordinate {
	return this.coord
}
