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
// Should direction be inside location directory?
func From(coordinate coord.AbsoluteCoordinate, direction direction.Direction) (*Location, error) {
	if coordinate.X() < 0 || coordinate.Y() < 0 {
		return nil, errors.New("no negative coordinates!")
	}
	return &Location{coordinate, coordinate, direction}, nil
}

// TODO: is this needed?
func (this *Location) Direction() string {
	return this.direction.CardinalPoint()
}

// TODO: should not exist
func (this *Location) Equals(other Location) bool {
	return this.coordinate.Equals(&other.coordinate)
}

func (this *Location) WillBeAt() coord.AbsoluteCoordinate {
	return this.futureCoordinate
}

func (this *Location) CommitMovement() {
	this.coordinate = this.futureCoordinate
}

func (this *Location) RollbackMovement() {
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

func (this *Location) StartMovementAhead() {
	this.futureCoordinate = *coord.SumOf(this.coordinate, this.direction.RelativePositionAhead())
}

func (this *Location) StartMovementBehind() {
	this.futureCoordinate = *coord.SumOf(this.coordinate, this.direction.RelativePositionBehind())
}

// TODO: add tests with mocks
func (this *Location) WrapAround(limit size.Size) {
	this.futureCoordinate = *coord.NewAbsolute(
		this.wrapX(limit.Width),
		this.wrapY(limit.Height),
	)
}

func (this *Location) wrapX(width int) int {
	if this.futureCoordinate.X() > width {
		return 0
	} else if this.futureCoordinate.X() < 0 {
		return width
	}
	return this.futureCoordinate.X()
}

func (this *Location) wrapY(height int) int {
	if this.futureCoordinate.Y() > height {
		return 0
	} else if this.futureCoordinate.Y() < 0 {
		return height
	}
	return this.futureCoordinate.Y()
}
