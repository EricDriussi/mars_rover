package absolute_position

type AbsolutePosition interface {
	X() int
	Y() int
	Equals(other AbsolutePosition) bool
}
