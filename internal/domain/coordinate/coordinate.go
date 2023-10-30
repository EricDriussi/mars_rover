package coordinate

import "mars_rover/internal/domain/size"

// TODO: add tests
type Coordinate struct {
	x, y int
}

func New(x, y int) *Coordinate {
	return &Coordinate{x, y}
}

func (this *Coordinate) Equals(other Coordinate) bool {
	return this.x == other.x && this.y == other.y
}

func (this *Coordinate) IsWithin(limit size.Size) bool {
	return this.x <= limit.Width && this.y <= limit.Height
}

func (this *Coordinate) WrapXIfOutOf(width int) {
	if this.x > width {
		this.x = 0
	} else if this.x < 0 {
		this.x = width
	}
}

func (this *Coordinate) WrapYIfOutOf(height int) {
	if this.y > height {
		this.y = 0
	} else if this.y < 0 {
		this.y = height
	}
}

func (this Coordinate) X() int {
	return this.x
}

func (this Coordinate) Y() int {
	return this.y
}
