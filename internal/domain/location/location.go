package location

import (
	"errors"
	coord "mars_rover/internal/domain/coordinate"
	relativePosition "mars_rover/internal/domain/relative_position"
	"mars_rover/internal/domain/size"
)

type Location struct {
	coordinate coord.Coordinate
}

func From(coordinate coord.Coordinate) (*Location, error) {
	if coordinate.X() < 0 || coordinate.Y() < 0 {
		return nil, errors.New("no negative coordinates!")
	}
	return &Location{coordinate}, nil
}

func (this *Location) Equals(other Location) bool {
	return this.coordinate.Equals(other.coordinate)
}

func (this *Location) UpdateCoordinate(coordinate coord.Coordinate) {
	this.coordinate = coordinate
}

func (this *Location) WillBeAt(relativePosition relativePosition.RelativePosition, size size.Size) coord.Coordinate {
	futureCoordinate := coord.New(
		this.coordinate.X()+relativePosition.X(),
		this.coordinate.Y()+relativePosition.Y(),
	)
	futureCoordinate.WrapXIfOutOf(size.Width)
	futureCoordinate.WrapYIfOutOf(size.Height)
	return *futureCoordinate
}
