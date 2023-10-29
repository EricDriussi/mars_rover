package coordinate

type Coordinate struct {
	X, Y int
}

func (this *Coordinate) WrapXIfOutOf(width int) {
	if this.X > width {
		this.X = 0
	} else if this.X < 0 {
		this.X = width
	}
}

func (this *Coordinate) WrapYIfOutOf(height int) {
	if this.Y > height {
		this.Y = 0
	} else if this.Y < 0 {
		this.Y = height
	}
}
