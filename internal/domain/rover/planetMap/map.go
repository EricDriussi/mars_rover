package planetMap

import (
	"mars_rover/internal/domain/coordinate/absoluteCoordinate"
	"mars_rover/internal/domain/obstacle/obstacles"
	"mars_rover/internal/domain/planet"
	"mars_rover/internal/domain/size"
)

type Map struct {
	size      size.Size
	obstacles obstacles.Obstacles
}

func Of(planet planet.Planet) *Map {
	return &Map{planet.Size(), planet.Obstacles()}
}

func (this *Map) Size() size.Size {
	return this.size
}

func (this *Map) Obstacles() obstacles.Obstacles {
	return this.obstacles
}

func (this *Map) CollidesWithObstacle(coord absoluteCoordinate.AbsoluteCoordinate) bool {
	return this.obstacles.Occupy(coord)
}

func (this *Map) IsOutOfBounds(coord absoluteCoordinate.AbsoluteCoordinate) bool {
	return coord.X() > this.size.Width() || coord.Y() > this.size.Height() || coord.X() < 0 || coord.Y() < 0
}
