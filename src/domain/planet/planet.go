package planet

import (
	. "mars_rover/src/domain/obstacle/obstacles"
	"mars_rover/src/domain/planet/emptyPlanet"
	"mars_rover/src/domain/planet/planetWithObstacles"
	. "mars_rover/src/domain/size"
)

type Planet interface {
	// TODO.LM: Color is not used for much, it's just here to make apparent the decoupling between the planet and the map
	// The map doesn't know or care about the planet's color and the planet could have more attributes like this
	Color() string
	Size() Size
	Obstacles() Obstacles
}

func CreatePlanet(color string, size Size, obstacles Obstacles) (Planet, error) {
	if obstacles.Amount() == 0 {
		return emptyPlanet.Create(color, size)
	}
	return planetWithObstacles.Create(color, size, obstacles)
}
