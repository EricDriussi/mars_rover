package planetWithObstacles

import (
	"errors"
	. "mars_rover/src/domain/obstacle/obstacles"
	. "mars_rover/src/domain/size"
)

type PlanetWithObstacles struct {
	color     string
	size      Size
	obstacles Obstacles
}

func Create(color string, size Size, obstacles Obstacles) (*PlanetWithObstacles, error) {
	if obstacles.Amount() < 1 {
		return nil, errors.New("cannot create without obstacles")
	}
	if size.Area() < 2 {
		return nil, errors.New("size too small") // rover + 1 obstacle would not fit
	}
	if obstacles.IsAnyBeyond(size) {
		return nil, errors.New("an obstacle was set outside of the planet :c")
	}

	return &PlanetWithObstacles{color, size, obstacles}, nil
}

func (this *PlanetWithObstacles) Size() Size {
	return this.size
}

func (this *PlanetWithObstacles) Obstacles() Obstacles {
	return this.obstacles
}

func (this *PlanetWithObstacles) Color() string {
	return this.color
}
