package planet

import (
	. "mars_rover/internal/domain/obstacle/obstacles"
	. "mars_rover/internal/domain/size"
)

type Planet interface {
	Size() Size
	Obstacles() Obstacles
}
