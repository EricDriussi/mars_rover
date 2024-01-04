package absoluteCoordinate

type AbsoluteCoordinate struct {
	x, y int
}

func Build(x, y int) *AbsoluteCoordinate {
	return &AbsoluteCoordinate{x, y}
}

func (this *AbsoluteCoordinate) Equals(other AbsoluteCoordinate) bool {
	return this.X() == other.X() && this.Y() == other.Y()
}

func (this *AbsoluteCoordinate) X() int {
	return this.x
}

func (this *AbsoluteCoordinate) Y() int {
	return this.y
}

func (this *AbsoluteCoordinate) IsAdjacentTo(coordinate AbsoluteCoordinate) bool {
	sameX := this.X() == coordinate.X()
	sameY := this.Y() == coordinate.Y()
	adjacentY := this.Y() == coordinate.Y()+1 || this.Y() == coordinate.Y()-1
	adjacentX := this.X() == coordinate.X()+1 || this.X() == coordinate.X()-1
	return sameX && adjacentY || sameY && adjacentX
}
