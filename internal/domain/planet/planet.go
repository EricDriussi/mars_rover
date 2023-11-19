package planet

import (
	. "mars_rover/internal/domain/obstacle/obstacles"
	. "mars_rover/internal/domain/size"
)

type Planet interface {
	Color() string
	Size() Size
	Obstacles() Obstacles
}
