package location

import (
	"errors"
	"mars_rover/internal/domain/coordinate"
	relativePosition "mars_rover/internal/domain/relative_position"
	"mars_rover/internal/domain/size"
)

type Location struct {
	coordinate coordinate.Coordinate
}

func From(x, y int) (*Location, error) {
	if x < 0 || y < 0 {
		return nil, errors.New("no negative locations!")
	}
	return &Location{coordinate.Coordinate{X: x, Y: y}}, nil
}

func (this Location) Equals(other Location) bool {
	return this.coordinate.X == other.coordinate.X && this.coordinate.Y == other.coordinate.Y
}

func (this Location) IsWithin(limit size.Size) bool {
	return this.coordinate.X <= limit.Width && this.coordinate.Y <= limit.Height
}

func (this *Location) WillBeAt(relativePosition relativePosition.RelativePosition, size size.Size) Location {
	futureCoordinate := &coordinate.Coordinate{
		X: this.coordinate.X + relativePosition.X(),
		Y: this.coordinate.Y + relativePosition.Y(),
	}
	futureCoordinate.WrapXIfOutOf(size.Width)
	futureCoordinate.WrapYIfOutOf(size.Height)
	return Location{*futureCoordinate}
}
