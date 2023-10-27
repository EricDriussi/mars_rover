package size

import "errors"

type Size struct {
	Width, Height int
}

func From(width, height int) (*Size, error) {
	if width <= 0 || height <= 0 {
		return nil, errors.New("invalid size!")
	}
	return &Size{width, height}, nil
}
