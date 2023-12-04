package planetMap

import (
	. "mars_rover/src/domain/coordinate/absoluteCoordinate"
	. "mars_rover/src/domain/obstacle/obstacles"
	. "mars_rover/src/domain/planet"
	. "mars_rover/src/domain/size"
)

type Map struct {
	size      Size
	obstacles Obstacles
}

func OfPlanet(planet Planet) *Map {
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
	// Planet size starts at 1, rover movement starts at 0
	return absoluteCoordinate.X() >= this.size.Width() || absoluteCoordinate.Y() >= this.size.Height() || absoluteCoordinate.X() < 0 || absoluteCoordinate.Y() < 0
}
