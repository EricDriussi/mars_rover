package relativePosition

import (
	"mars_rover/internal/domain/coordinate"
	coordinate2d "mars_rover/internal/domain/coordinate/coordinate2D"
)

// rename and extract
type IRelativePosition interface {
	X() int
	Y() int
}

// should reflect 1-step-only movements
// Relative2DSingleStep?
type RelativePosition struct {
	coordinate coordinate.Coordinate
}

// should have tests to back up the name, only combinations of 0, 1 and -1 are allowed
// is this a From?
func New(x, y int) *RelativePosition {
	return &RelativePosition{coordinate2d.New(x, y)}
}

func (this RelativePosition) X() int {
	return this.coordinate.X()
}

func (this RelativePosition) Y() int {
	return this.coordinate.Y()
}
