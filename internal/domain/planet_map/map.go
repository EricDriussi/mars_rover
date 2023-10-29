package planetMap

import (
	"mars_rover/internal/domain/coordinate"
	"mars_rover/internal/domain/obstacle"
	"mars_rover/internal/domain/planet"
	"mars_rover/internal/domain/size"
)

type Map struct {
	size      size.Size
	obstacles []obstacle.Obstacle
}

func Of(planet planet.Planet) *Map {
	return &Map{planet.Size, planet.Obstacles}
}

func (this Map) CheckCollision(coordinates coordinate.Coordinate) bool {
	for _, obstacle := range this.obstacles {
		return coordinates.Equals(*obstacle.Position)
	}
	return false
}

func (this Map) WouldGoOutOfBounds(coordinates coordinate.Coordinate) bool {
	return !coordinates.IsWithin(this.size)
}

func (this Map) Size() size.Size {
	return this.size
}
