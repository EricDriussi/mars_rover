package rockyPlanet

import (
	"errors"
	"mars_rover/internal/domain/obstacle"
	"mars_rover/internal/domain/size"
)

type RockyPlanet struct {
	color     string
	size      size.Size
	obstacles []obstacle.Obstacle
}

func Create(color string, size size.Size, obstacles []obstacle.Obstacle) (*RockyPlanet, error) {
	for _, obs := range obstacles {
		if obs.IsBeyond(size) {
			return nil, errors.New("an obstacle was set outside of the planet :c")
		}
	}

	return &RockyPlanet{color, size, obstacles}, nil
}

func (this *RockyPlanet) Size() size.Size {
	return this.size
}

func (this *RockyPlanet) Obstacles() []obstacle.Obstacle {
	return this.obstacles
}
