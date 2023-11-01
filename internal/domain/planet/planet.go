package planet

import (
	"mars_rover/internal/domain/obstacle"
	"mars_rover/internal/domain/size"
)

type Planet interface {
	Size() size.Size
	Obstacles() []obstacle.Obstacle
}
