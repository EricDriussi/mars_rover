package obstacle

import (
	coord "mars_rover/internal/domain/coordinate"
	"mars_rover/internal/domain/size"
)

type Obstacle interface {
	IsBeyond(size.Size) bool
	Coordinate() coord.Coordinate
}
