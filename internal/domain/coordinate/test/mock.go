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

func (this *MockCoordinate) IsOutsideOf(limit size.Size) bool {
	args := this.Called(limit)
	return args.Bool(0)
}

func (this *MockCoordinate) Equals(other coordinate.Coordinate) bool {
	args := this.Called(other)
	return args.Bool(0)
}

func (this *MockCoordinate) X() int {
	return this.x
}

func (this *MockCoordinate) Y() int {
	return this.y
}
