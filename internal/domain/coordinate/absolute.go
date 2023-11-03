package coordinate

import (
	"mars_rover/internal/domain/size"
)

type AbsoluteCoordinate struct {
	x, y int
}

func NewAbsolute(x, y int) *AbsoluteCoordinate {
	return &AbsoluteCoordinate{x, y}
}

// TODO: add tests
func SumOf(coordinateOne, coordinateTwo Coordinate) *AbsoluteCoordinate {
	return &AbsoluteCoordinate{coordinateOne.X() + coordinateTwo.X(), coordinateOne.Y() + coordinateTwo.Y()}
}

func (this *AbsoluteCoordinate) WrapIfOutOf(limit size.Size) {
	this.wrapXIfOutOf(limit.Width)
	this.wrapYIfOutOf(limit.Height)
}

func (this *AbsoluteCoordinate) Equals(other *AbsoluteCoordinate) bool {
	return this.x == other.X() && this.y == other.Y()
}

func (this AbsoluteCoordinate) X() int {
	return this.x
}

func (this AbsoluteCoordinate) Y() int {
	return this.y
}

func (this *AbsoluteCoordinate) wrapXIfOutOf(width int) {
	if this.x > width {
		this.x = 0
	} else if this.x < 0 {
		this.x = width
	}
}

func (this *AbsoluteCoordinate) wrapYIfOutOf(height int) {
	if this.y > height {
		this.y = 0
	} else if this.y < 0 {
		this.y = height
	}
}
