package coordinate2d

import (
	"mars_rover/internal/domain/coordinate"
	"mars_rover/internal/domain/size"
)

type Coordinate2D struct {
	x, y int
}

func New(x, y int) *Coordinate2D {
	return &Coordinate2D{x, y}
}

func (this *Coordinate2D) WrapIfOutOf(limit size.Size) {
	this.wrapXIfOutOf(limit.Width)
	this.wrapYIfOutOf(limit.Height)
}

func (this *Coordinate2D) Equals(other coordinate.Coordinate) bool {
	return this.x == other.X() && this.y == other.Y()
}

func (this Coordinate2D) X() int {
	return this.x
}

func (this Coordinate2D) Y() int {
	return this.y
}

func (this *Coordinate2D) wrapXIfOutOf(width int) {
	if this.x > width {
		this.x = 0
	} else if this.x < 0 {
		this.x = width
	}
}

func (this *Coordinate2D) wrapYIfOutOf(height int) {
	if this.y > height {
		this.y = 0
	} else if this.y < 0 {
		this.y = height
	}
}
