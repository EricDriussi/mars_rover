package planet

import (
	. "mars_rover/src/domain/obstacle/obstacles"
	. "mars_rover/src/domain/size"
)

type Planet interface {
	Color() string
	Size() Size
	Obstacles() Obstacles
}
