package relativeCoordinate

const step = 1 // this should probably be a config var

type RelativeCoordinate struct {
	x, y int
}

func Up() *RelativeCoordinate {
	return build(0, 1)
}

func Right() *RelativeCoordinate {
	return build(1, 0)
}

func Down() *RelativeCoordinate {
	return build(0, -1)
}

func Left() *RelativeCoordinate {
	return build(-1, 0)
}

func build(x, y int) *RelativeCoordinate {
	return &RelativeCoordinate{x * step, y * step}
}

func (this *RelativeCoordinate) X() int {
	return this.x
}

func (this *RelativeCoordinate) Y() int {
	return this.y
}
