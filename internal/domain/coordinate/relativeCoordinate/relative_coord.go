package relativeCoordinate

type RelativeCoordinate struct {
	x, y int
}

func New(x, y int) *RelativeCoordinate {
	step := 1
	if isOrthogonal(x, y) {
		return &RelativeCoordinate{x * step, y * step}
	}
	// TODO.LM: should this return an error?
	return &RelativeCoordinate{0, 0}
}

func (this *RelativeCoordinate) X() int {
	return this.x
}

func (this *RelativeCoordinate) Y() int {
	return this.y
}

func isOrthogonal(x, y int) bool {
	if x == 0 && isAdjacent(y) {
		return true
	}
	if y == 0 && isAdjacent(x) {
		return true
	}
	return false
}

func isAdjacent(num int) bool {
	return num == 1 || num == -1
}
