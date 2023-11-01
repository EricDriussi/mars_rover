package location

import (
	"errors"
	coord "mars_rover/internal/domain/coordinate"
	"mars_rover/internal/domain/location/direction"
	"mars_rover/internal/domain/size"
)

type Location struct {
	coordinate  coord.Coordinate
	orientation direction.Direction
}

// TODO.LM: Should "direction" be exposed? Should "From" take (coord, "N")?
// Should "From" take a "directionDTO" enum and create the corresponding direction with a factory?
// Should direction be inside location directory?
func From(coordinate coord.Coordinate, direction direction.Direction) (*Location, error) {
	if coordinate.X() < 0 || coordinate.Y() < 0 {
		return nil, errors.New("no negative coordinates!")
	}
	return &Location{coordinate, direction}, nil
}

func (this *Location) Orientation() string {
	return this.orientation.CardinalPoint()
}

func (this *Location) Position() coord.Coordinate {
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
	return this.coordinate.Equals(other.coordinate)
}

func (this *Location) UpdateCoordinate(coordinate coord.Coordinate) {
	this.coordinate = coordinate
}

func (this *Location) AheadWillBeAt(size size.Size) coord.Coordinate {
	futureCoordinate := coord.New(
		this.coordinate.X()+this.orientation.RelativePositionAhead().X(),
		this.coordinate.Y()+this.orientation.RelativePositionAhead().Y(),
	)
	futureCoordinate.WrapIfOutOf(size)
	return futureCoordinate
}

func (this *Location) BehindWillBeAt(size size.Size) coord.Coordinate {
	futureCoordinate := coord.New(
		this.coordinate.X()+this.orientation.RelativePositionBehind().X(),
		this.coordinate.Y()+this.orientation.RelativePositionBehind().Y(),
	)
	futureCoordinate.WrapIfOutOf(size)
	return futureCoordinate
}
