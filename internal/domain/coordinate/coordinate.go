package coordinate

type Coordinate interface {
	X() int
	Y() int
}

func SumOf(coordinateOne AbsoluteCoordinate, coordinateTwo RelativeCoordinate) *AbsoluteCoordinate {
	return &AbsoluteCoordinate{coordinateOne.X() + coordinateTwo.X(), coordinateOne.Y() + coordinateTwo.Y()}
}
