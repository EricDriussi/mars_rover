package location

import (
	"errors"
	coord "mars_rover/internal/domain/coordinate"
	"mars_rover/internal/domain/location/direction"
	"mars_rover/internal/domain/size"
)

type Location struct {
	coordinate  coord.AbsoluteCoordinate
	orientation direction.Direction
}

// TODO.LM: Should "direction" be exposed? Should "From" take (coord, "N")?
// Should "From" take a "directionDTO" enum and create the corresponding direction with a factory?
// Should direction be inside location directory?
func From(coordinate coord.AbsoluteCoordinate, direction direction.Direction) (*Location, error) {
	if coordinate.X() < 0 || coordinate.Y() < 0 {
		return nil, errors.New("no negative coordinates!")
	}
	return &Location{coordinate, direction}, nil
}

func (this *Location) Orientation() string {
	return this.orientation.CardinalPoint()
}

func (this *Location) Position() coord.AbsoluteCoordinate {
	return this.coordinate
}

func (this *Location) UpdateWithDirectionOnTheLeft() {
	this.orientation = this.orientation.DirectionOnTheLeft()
}

func (this *Location) UpdateWithDirectionOnTheRight() {
	this.orientation = this.orientation.DirectionOnTheRight()
}

// TODO: should not exist
func (this *Location) Equals(other Location) bool {
	return this.coordinate.Equals(&other.coordinate)
}

func (this *Location) UpdateCoordinate(coordinate coord.AbsoluteCoordinate) {
	this.coordinate = coordinate
}

func (this *Location) AheadWillBeAt(size size.Size) coord.AbsoluteCoordinate {
	futureCoordinate := *coord.SumOf(this.coordinate, this.orientation.RelativePositionAhead())
	futureCoordinate.WrapIfOutOf(size)
	return futureCoordinate
}

func (this *Location) BehindWillBeAt(size size.Size) coord.AbsoluteCoordinate {
	futureCoordinate := *coord.SumOf(this.coordinate, this.orientation.RelativePositionBehind())
	futureCoordinate.WrapIfOutOf(size)
	return futureCoordinate
}
