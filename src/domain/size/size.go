package size

import (
	"errors"
	"math"
)

var sqrtOfMaxInt int

// Magic Go function that gets called on import
// spares redundant calculations at runtime
func init() {
	maxIntAsFloat := float64(math.MaxInt)
	sqrtOfMaxIntAsFloat := math.Round(math.Sqrt(maxIntAsFloat))
	sqrtOfMaxInt = int(sqrtOfMaxIntAsFloat)
}

type Size struct {
	width, height int
}

func Square(side int) (*Size, error) {
	if side < 1 { // size 0 is no size
		return nil, errors.New("invalid size!")
	}
	if side > sqrtOfMaxInt { // a bigger side would overflow when calculating the area
		return &Size{sqrtOfMaxInt, sqrtOfMaxInt}, nil
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
