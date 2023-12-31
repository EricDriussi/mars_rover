package bigRock

import (
	. "mars_rover/src/domain/coordinate/absoluteCoordinate"
	. "mars_rover/src/domain/coordinate/coordinates"
	coord "mars_rover/src/domain/coordinate/coordinates"
	. "mars_rover/src/domain/size"
)

type BigRock struct {
	coordinates Coordinates
}

func In(coordinates []AbsoluteCoordinate) BigRock {
	//TODO: if len(coordinates) < 2 then small rock?
	return BigRock{*coord.New(coordinates)}
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
