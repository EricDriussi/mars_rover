package small_rock

import (
	coord "mars_rover/internal/domain/coordinate/absoluteCoordinate"
	"mars_rover/internal/domain/obstacle"
	"mars_rover/internal/domain/size"
)

type SmallRock struct {
	coordinate coord.AbsoluteCoordinate
}

func In(coordinate coord.AbsoluteCoordinate) obstacle.Obstacle {
	return &SmallRock{coordinate}
}

func (this SmallRock) Occupies(coordinate coord.AbsoluteCoordinate) bool {
	return this.coordinate.Equals(&coordinate)
}

func (this SmallRock) IsBeyond(size size.Size) bool {
	return this.coordinate.X() > size.Width || this.coordinate.Y() > size.Height
}

func (this SmallRock) Coordinates() coord.AbsoluteCoordinate {
	return this.coordinate
}
