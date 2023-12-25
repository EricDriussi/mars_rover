package size

import (
	"errors"
	"math"
)

var sqrtMaxIntAsInt int

// Magic Go function that gets called on import
// spares redundant calculations at runtime
func init() {
	maxIntAsFloat := float64(math.MaxInt)
	sqrtMaxInt := math.Round(math.Sqrt(maxIntAsFloat))
	sqrtMaxIntAsInt = int(sqrtMaxInt)
}

type Size struct {
	width, height int
}

func Square(side int) (*Size, error) {
	if side <= 0 { // size 0 is no size
		return nil, errors.New("invalid size!")
	}
	if side > sqrtMaxIntAsInt { // a bigger side would overflow when calculating the area
		return &Size{sqrtMaxIntAsInt, sqrtMaxIntAsInt}, nil
	}
	return &Size{side, side}, nil
}

func (this *Size) Width() int {
	return this.width
}

func (this *Size) Height() int {
	return this.height
}

func (this *Size) Area() int {
	return this.width * this.height
}
