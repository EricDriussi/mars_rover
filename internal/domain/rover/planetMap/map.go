package planetMap

import (
	. "mars_rover/internal/domain/coordinate/absoluteCoordinate"
	. "mars_rover/internal/domain/obstacle/obstacles"
	. "mars_rover/internal/domain/planet"
	. "mars_rover/internal/domain/size"
)

type Map struct {
	size      Size
	obstacles Obstacles
}

func Of(planet Planet) *Map {
	return &Map{planet.Size(), planet.Obstacles()}
}

func (this *Map) Width() int {
	return this.size.Width()
}

func (this *Map) Height() int {
	return this.size.Height()
}

func (this *Map) Obstacles() Obstacles {
	return this.obstacles
}

func (this *Map) HasObstacleIn(absoluteCoordinate AbsoluteCoordinate) bool {
	return this.obstacles.Occupy(absoluteCoordinate)
}

func (this *Map) IsOutOfBounds(absoluteCoordinate AbsoluteCoordinate) bool {
	return absoluteCoordinate.X() > this.size.Width() || absoluteCoordinate.Y() > this.size.Height() || absoluteCoordinate.X() < 0 || absoluteCoordinate.Y() < 0
}
