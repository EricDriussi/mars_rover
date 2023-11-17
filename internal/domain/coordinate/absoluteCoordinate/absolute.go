package absoluteCoordinate

type AbsoluteCoordinate struct {
	x, y int
}

func From(x, y int) *AbsoluteCoordinate {
	return &AbsoluteCoordinate{x, y}
}

func (this *AbsoluteCoordinate) Equals(other *AbsoluteCoordinate) bool {
	return this.x == other.X() && this.y == other.Y()
}

func (this *AbsoluteCoordinate) X() int {
	return this.x
}

func (this *AbsoluteCoordinate) Y() int {
	return this.y
}
