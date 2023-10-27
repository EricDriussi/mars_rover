package planetMap

import (
	"mars_rover/internal/domain/obstacle"
	"mars_rover/internal/domain/planet"
	"mars_rover/internal/domain/position"
	"mars_rover/internal/domain/size"
)

type PlanetMap struct {
	planetSize      size.Size
	planetObstacles []obstacle.Obstacle
}

func Of(planet planet.Planet) *PlanetMap {
	return &PlanetMap{planet.Size, planet.Obstacles}
}

func (this PlanetMap) CheckCollision(position position.Position) bool {
	for _, obstacle := range this.planetObstacles {
		return position.Equals(*obstacle.Position)
	}
	return false
}

func (this PlanetMap) WouldGoOutOfBounds(position position.Position) bool {
	return !position.IsWithin(this.planetSize)
}

func (this PlanetMap) Size() size.Size {
	return this.planetSize
}
