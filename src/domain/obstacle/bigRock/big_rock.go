package bigRock

import (
	"errors"
	. "mars_rover/src/domain/coordinate/absoluteCoordinate"
	"mars_rover/src/domain/coordinate/coordinates"
	. "mars_rover/src/domain/coordinate/coordinates"
	. "mars_rover/src/domain/size"
)

type BigRock struct {
	coordinates Coordinates
}

func In(occupiedCoordinates ...AbsoluteCoordinate) (*BigRock, error) {
	coords := coordinates.New(occupiedCoordinates...)
	if len(coords.List()) < 2 {
		return nil, errors.New("cannot create big rock with less than 2 coordinates")
	}
	return &BigRock{*coords}, nil
}

func (this *BigRock) Occupies(coordinate AbsoluteCoordinate) bool {
	return this.coordinates.Contain(coordinate)
}

func (this *BigRock) IsBeyond(size Size) bool {
	return this.coordinates.Overflow(size)
}

func (this *BigRock) Coordinates() []AbsoluteCoordinate {
	return this.coordinates.List()
}
