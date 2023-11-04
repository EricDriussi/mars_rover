package coordinate

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

func (this *AbsoluteCoordinate) Equals(other *AbsoluteCoordinate) bool {
	return this.x == other.X() && this.y == other.Y()
}

func (this AbsoluteCoordinate) X() int {
	return this.x
}

func (this AbsoluteCoordinate) Y() int {
	return this.y
}
