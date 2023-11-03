package coordinate

import (
	"mars_rover/internal/domain/size"
)

type Coordinate struct {
	x, y int
}

func New(x, y int) *Coordinate {
	return &Coordinate{x, y}
}

func (this *Coordinate) WrapIfOutOf(limit size.Size) {
	this.wrapXIfOutOf(limit.Width)
	this.wrapYIfOutOf(limit.Height)
}

func (this *Coordinate) Equals(other *Coordinate) bool {
	return this.x == other.X() && this.y == other.Y()
}

func (this Coordinate) X() int {
	return this.x
}

func (this Coordinate) Y() int {
	return this.y
}

func (this *Coordinate) wrapXIfOutOf(width int) {
	if this.x > width {
		this.x = 0
	} else if this.x < 0 {
		this.x = width
	}
}

func (this *Coordinate) wrapYIfOutOf(height int) {
	if this.y > height {
		this.y = 0
	} else if this.y < 0 {
		this.y = height
	}
}
