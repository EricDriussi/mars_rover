package rockyPlanet

import (
	"errors"
	"mars_rover/internal/domain/obstacle"
	obs "mars_rover/internal/domain/obstacle/obstacles"
	"mars_rover/internal/domain/size"
)

type RockyPlanet struct {
	color     string
	size      size.Size
	obstacles obs.Obstacles
}

func Create(color string, size size.Size, obstacles []obstacle.Obstacle) (*RockyPlanet, error) {
	obstacleList := obs.New(obstacles)
	if obstacleList.IsAnyBeyond(size) {
		return nil, errors.New("an obstacle was set outside of the planet :c")
	}

	return &RockyPlanet{color, size, *obstacleList}, nil
}

func (this *RockyPlanet) Size() size.Size {
	return this.size
}

func (this *RockyPlanet) Obstacles() obs.Obstacles {
	return this.obstacles
}
