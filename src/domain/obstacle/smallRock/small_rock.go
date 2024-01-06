package smallRock

import (
	. "mars_rover/src/domain/coordinate/absoluteCoordinate"
	"mars_rover/src/domain/coordinate/coordinates"
	. "mars_rover/src/domain/coordinate/coordinates"
	. "mars_rover/src/domain/size"
)

type SmallRock struct {
	coordinates Coordinates
}

func In(coordinate AbsoluteCoordinate) (*SmallRock, error) {
	coords, err := coordinates.New(coordinate)
	if err != nil {
		return nil, err
	}
	return &SmallRock{*coords}, nil
}

func (this *SmallRock) Occupies(coordinate AbsoluteCoordinate) bool {
	return this.coordinates.Contain(coordinate)
}

func (this *SmallRock) IsBeyond(size Size) bool {
	singleCoordinate := this.coordinates.First()
	return singleCoordinate.X() > size.Width() || singleCoordinate.Y() > size.Height()
}

func (this *SmallRock) Coordinates() Coordinates {
	return this.coordinates
}
