package planet

import (
	. "mars_rover/src/domain/obstacle/obstacles"
	"mars_rover/src/domain/planet/emptyPlanet"
	"mars_rover/src/domain/planet/rockyPlanet"
	. "mars_rover/src/domain/size"
)

type Planet interface {
	Color() string
	Size() Size
	Obstacles() Obstacles
}

func CreatePlanet(color string, size Size, obstacles Obstacles) (Planet, error) {
	if len(obstacles.List()) == 0 {
		return emptyPlanet.Create(color, size)
	}
	return rockyPlanet.Create(color, size, obstacles.List())
}
