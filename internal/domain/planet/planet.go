package planet

import (
	obs "mars_rover/internal/domain/obstacle/obstacles"
	"mars_rover/internal/domain/size"
)

type Planet interface {
	Size() size.Size
	Obstacles() obs.Obstacles
}
