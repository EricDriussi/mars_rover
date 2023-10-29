package location

import (
	"errors"
	"mars_rover/internal/domain/coordinate"
	relativePosition "mars_rover/internal/domain/relative_position"
	"mars_rover/internal/domain/size"
)

type Location struct {
	coordinate coordinate.Coordinate
}

func From(x, y int) (*Location, error) {
	if x < 0 || y < 0 {
		return nil, errors.New("no negative locations!")
	}
	return &Location{coordinate.Coordinate{X: x, Y: y}}, nil
}

func (this Location) Equals(other Location) bool {
	return this.coordinate.X == other.coordinate.X && this.coordinate.Y == other.coordinate.Y
}

func (this Location) IsWithin(limit size.Size) bool {
	return this.coordinate.X < limit.Width && this.coordinate.Y < limit.Height
}

func (this *Location) WillBeAt(relativePosition relativePosition.RelativePosition, size size.Size) Location {
	x := this.coordinate.X + relativePosition.X()
	y := this.coordinate.Y + relativePosition.Y()
	tmp := &coordinate.Coordinate{X: x, Y: y}
	if tmp.X >= size.Width {
		tmp.X = 0
	} else if tmp.X < 0 {
		tmp.X = size.Width
	}

	if tmp.Y >= size.Height {
		tmp.Y = 0
	} else if tmp.Y < 0 {
		tmp.Y = size.Height
	}
	return Location{*tmp}
}

func (this *Location) AddOrWrap(relativePosition relativePosition.RelativePosition, size size.Size) {
	this.add(relativePosition)
	this.wrapIfOutOfBounds(size)
}

func (this *Location) add(relativePosition relativePosition.RelativePosition) {
	this.coordinate.X += relativePosition.X()
	this.coordinate.Y += relativePosition.Y()
}

func (this *Location) wrapIfOutOfBounds(size size.Size) {
	this.wrapXIfGreaterThan(size.Width)
	this.wrapYIfGreaterThan(size.Height)
}

func (this *Location) wrapXIfGreaterThan(width int) {
	if this.coordinate.X >= width {
		this.coordinate.X = 0
	} else if this.coordinate.X < 0 {
		this.coordinate.X = width
	}
}

func (this *Location) wrapYIfGreaterThan(height int) {
	if this.coordinate.Y >= height {
		this.coordinate.Y = 0
	} else if this.coordinate.Y < 0 {
		this.coordinate.Y = height
	}
}
