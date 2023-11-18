package planetMap

import (
	"mars_rover/internal/domain/coordinate/absoluteCoordinate"
	obs "mars_rover/internal/domain/obstacle"
	"mars_rover/internal/domain/planet"
	"mars_rover/internal/domain/size"
)

type Map struct {
	size      size.Size
	obstacles []obs.Obstacle
}

func Of(planet planet.Planet) *Map {
	return &Map{planet.Size(), planet.Obstacles()}
}

func (this *Map) Size() size.Size {
	return this.size
}

func (this *Map) Obstacles() []obs.Obstacle {
	return this.obstacles
}

func (this *Map) CollidesWithObstacle(coord absoluteCoordinate.AbsoluteCoordinate) bool {
	for _, obstacle := range this.obstacles {
		if obstacle.Occupies(coord) {
			return true
		}
	}
	return false
}

func (this *Map) IsOutOfBounds(coord absoluteCoordinate.AbsoluteCoordinate) bool {
	return coord.X() > this.size.Width() || coord.Y() > this.size.Height() || coord.X() < 0 || coord.Y() < 0
}
