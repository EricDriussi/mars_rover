package test

import (
	"mars_rover/internal/domain/coordinate"
	"mars_rover/internal/domain/size"

	"github.com/stretchr/testify/mock"
)

type MockCoordinate struct {
	mock.Mock
	x, y int
}

func (this *MockCoordinate) WrapIfOutOf(limit size.Size) {
}

func (this *MockCoordinate) Equals(other coordinate.Coordinate) bool {
	args := this.Called(other)
	return args.Bool(0)
}

func (this *MockCoordinate) X() int {
	args := this.Called()
	return args.Int(0)
}

func (this *MockCoordinate) Y() int {
	args := this.Called()
	return args.Int(0)
}
