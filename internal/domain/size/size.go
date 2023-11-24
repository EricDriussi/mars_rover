package size

import "errors"

type Size struct {
	width, height int
}

func Square(side int) (*Size, error) {
	if side <= 0 {
		return nil, errors.New("invalid size!")
	}
	return &Size{side, side}, nil
}

func Shape(width, height int) (*Size, error) {
	if width <= 0 || height <= 0 {
		return nil, errors.New("invalid size!")
	}
	return &Size{width, height}, nil
}

func (this *Size) Width() int {
	return this.width
}

func (this *Size) Height() int {
	return this.height

}
