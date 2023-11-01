package test

import (
	"mars_rover/internal/domain/coordinate"
	"mars_rover/internal/domain/obstacle"
	"mars_rover/internal/domain/obstacle/test"
	"mars_rover/internal/domain/size"

	"github.com/stretchr/testify/mock"
)

type MockPlanet struct {
	mock.Mock
	coord coordinate.Coordinate
}

func (this *MockPlanet) Size() size.Size {
	return size.Size{Width: 1, Height: 1}
}

func (this *MockPlanet) Obstacles() []obstacle.Obstacle {
	return []obstacle.Obstacle{new(test.MockObstacle)}
}
