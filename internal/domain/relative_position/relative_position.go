package relativePosition

type RelativePosition struct {
	x, y int
}

func New(x, y int) *RelativePosition {
	return &RelativePosition{x, y}
}

func (this RelativePosition) X() int {
	return this.x
}

func (this RelativePosition) Y() int {
	return this.y
}
