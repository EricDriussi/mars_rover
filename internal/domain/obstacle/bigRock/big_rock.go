package bigRock

import (
	abs "mars_rover/internal/domain/coordinate/absoluteCoordinate"
	coord "mars_rover/internal/domain/coordinate/coordinates"
	"mars_rover/internal/domain/obstacle"
	"mars_rover/internal/domain/size"
)

type BigRock struct {
	coordinates coord.Coordinates
}

func In(coordinates []abs.AbsoluteCoordinate) obstacle.Obstacle {
	return &BigRock{*coord.New(coordinates)}
}

func (this *BigRock) Occupies(coordinate abs.AbsoluteCoordinate) bool {
	return this.coordinates.Contain(coordinate)
}

func (this *BigRock) IsBeyond(size size.Size) bool {
	return this.coordinates.GoBeyond(size)
}

func (this *BigRock) Coordinates() []abs.AbsoluteCoordinate {
	return this.coordinates.List()
}
