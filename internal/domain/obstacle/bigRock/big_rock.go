package bigRock

import (
	. "mars_rover/internal/domain/coordinate/absoluteCoordinate"
	. "mars_rover/internal/domain/coordinate/coordinates"
	coord "mars_rover/internal/domain/coordinate/coordinates"
	. "mars_rover/internal/domain/obstacle"
	. "mars_rover/internal/domain/size"
)

type BigRock struct {
	coordinates Coordinates
}

func In(coordinates []AbsoluteCoordinate) Obstacle {
	return &BigRock{*coord.New(coordinates)}
}

func (this *BigRock) Occupies(coordinate AbsoluteCoordinate) bool {
	return this.coordinates.Contain(coordinate)
}

func (this *BigRock) IsBeyond(size Size) bool {
	return this.coordinates.GoBeyond(size)
}

func (this *BigRock) Coordinates() []AbsoluteCoordinate {
	return this.coordinates.List()
}
