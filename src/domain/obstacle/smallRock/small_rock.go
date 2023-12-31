package smallRock

import (
	. "mars_rover/src/domain/coordinate/absoluteCoordinate"
	. "mars_rover/src/domain/size"
)

type SmallRock struct {
	coordinate AbsoluteCoordinate
}

func In(coordinate AbsoluteCoordinate) SmallRock {
	return SmallRock{coordinate}
}

func (this *SmallRock) Occupies(coordinate AbsoluteCoordinate) bool {
	return this.coordinate.Equals(coordinate)
}

func (this *SmallRock) IsBeyond(size Size) bool {
	return this.coordinate.X() > size.Width() || this.coordinate.Y() > size.Height()
}

func (this *SmallRock) Coordinates() []AbsoluteCoordinate {
	return []AbsoluteCoordinate{this.coordinate}
}
