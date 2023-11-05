package location

import (
	"errors"
	coord "mars_rover/internal/domain/coordinate"
	"mars_rover/internal/domain/location/direction"
	"mars_rover/internal/domain/size"
)

type Location struct {
	coordinate       coord.AbsoluteCoordinate
	futureCoordinate coord.AbsoluteCoordinate
	direction        direction.Direction
}

// TODO.LM: Should "direction" be exposed? Should "From" take (coord, "N")?
// Should "From" take a "directionDTO" enum and create the corresponding direction with a factory?
func From(coordinate coord.AbsoluteCoordinate, direction direction.Direction) (*Location, error) {
	if coordinate.X() < 0 || coordinate.Y() < 0 {
		return nil, errors.New("no negative coordinates!")
	}
	return &Location{coordinate, coordinate, direction}, nil
}

func (this *Location) WillBeAt() coord.AbsoluteCoordinate {
	return this.futureCoordinate
}

func (this *Location) UpdatePosition() {
	this.coordinate = this.futureCoordinate
}

func (this *Location) Reset() {
	this.futureCoordinate = this.coordinate
}

func (this *Location) Position() coord.AbsoluteCoordinate {
	return this.coordinate
}

func (this *Location) FaceLeft() {
	this.direction = this.direction.DirectionOnTheLeft()
}

func (this *Location) FaceRight() {
	this.direction = this.direction.DirectionOnTheRight()
}

func (this *Location) CalculatePositionAhead() {
	this.futureCoordinate = *coord.SumOf(this.coordinate, this.direction.RelativePositionAhead())
}

func (this *Location) CalculatePositionBehind() {
	this.futureCoordinate = *coord.SumOf(this.coordinate, this.direction.RelativePositionBehind())
}

func (this *Location) WrapAround(limit size.Size) {
	this.futureCoordinate = *coord.NewAbsolute(
		this.wrap(this.futureCoordinate.X(), limit.Width),
		this.wrap(this.futureCoordinate.Y(), limit.Height),
	)
}

func (this *Location) wrap(point, limit int) int {
	if point > limit {
		return 0
	} else if point < 0 {
		return limit
	}
	return point
}
