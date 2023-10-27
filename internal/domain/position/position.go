package position

import (
	"errors"
	relativePosition "mars_rover/internal/domain/relative_position"
	"mars_rover/internal/domain/size"
)

type Position struct {
	x, y int
}

func From(x, y int) (*Position, error) {
	if x < 0 || y < 0 {
		return nil, errors.New("no negative positions!")
	}
	return &Position{x, y}, nil
}

func (this Position) Equals(other Position) bool {
	return this.x == other.x && this.y == other.y
}

func (this Position) IsWithin(limit size.Size) bool {
	return this.x < limit.Width && this.y < limit.Height
}

func (this *Position) AddOrWrap(relativePosition relativePosition.RelativePosition, size size.Size) {
	this.add(relativePosition)
	this.wrapIfOutOfSize(relativePosition, size)
}

func (this *Position) add(relativePosition relativePosition.RelativePosition) {
	this.x += relativePosition.X()
	this.y += relativePosition.Y()
}

func (this *Position) wrapIfOutOfSize(relativePosition relativePosition.RelativePosition, size size.Size) {
	if this.x >= size.Width {
		this.x = 0
	} else if this.x < 0 {
		this.x = size.Width
	}

	if this.y >= size.Height {
		this.y = 0
	} else if this.y < 0 {
		this.y = size.Height
	}
}
