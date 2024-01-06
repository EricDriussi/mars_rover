package bigRock

import (
	"errors"
	. "mars_rover/src/domain/coordinate/absoluteCoordinate"
	. "mars_rover/src/domain/coordinate/coordinates"
	. "mars_rover/src/domain/size"
)

// TODO.LM: this should be set by config
const (
	MinSize = 2
	MaxSize = 5
)

type BigRock struct {
	coordinates Coordinates
}

func In(coordinates Coordinates) (*BigRock, error) {
	if coordinates.Amount() < MinSize {
		return nil, errors.New("cannot create big rock with less than 2 coordinates")
	}
	if coordinates.Amount() > MaxSize {
		return nil, errors.New("cannot create big rock with more than 5 coordinates")
	}
	if !coordinates.AreContiguous() {
		return nil, errors.New("cannot create big rock with non-contiguous coordinates")
	}
	return &BigRock{coordinates}, nil
}

func (this *BigRock) Occupies(coordinate AbsoluteCoordinate) bool {
	return this.coordinates.Contain(coordinate)
}

func (this *BigRock) IsBeyond(size Size) bool {
	return this.coordinates.Overflow(size)
}

func (this *BigRock) Coordinates() Coordinates {
	return this.coordinates
}
