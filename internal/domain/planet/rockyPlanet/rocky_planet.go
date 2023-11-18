package rockyPlanet

import (
	"errors"
	. "mars_rover/internal/domain/obstacle"
	. "mars_rover/internal/domain/obstacle/obstacles"
	obs "mars_rover/internal/domain/obstacle/obstacles"
	. "mars_rover/internal/domain/size"
)

type RockyPlanet struct {
	color     string
	size      Size
	obstacles Obstacles
}

func Create(color string, size Size, obstacles []Obstacle) (*RockyPlanet, error) {
	obstacleList := obs.New(obstacles)
	if obstacleList.IsAnyBeyond(size) {
		return nil, errors.New("an obstacle was set outside of the planet :c")
	}

	return &RockyPlanet{color, size, *obstacleList}, nil
}

func (this *RockyPlanet) Size() Size {
	return this.size
}

func (this *RockyPlanet) Obstacles() Obstacles {
	return this.obstacles
}
