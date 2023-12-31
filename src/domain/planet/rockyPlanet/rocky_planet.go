package rockyPlanet

import (
	"errors"
	. "mars_rover/src/domain/obstacle"
	. "mars_rover/src/domain/obstacle/obstacles"
	obs "mars_rover/src/domain/obstacle/obstacles"
	. "mars_rover/src/domain/size"
)

type RockyPlanet struct {
	color     string
	size      Size
	obstacles Obstacles
}

func Create(color string, size Size, obstacles []Obstacle) (*RockyPlanet, error) {
	if len(obstacles) < 1 {
		return nil, errors.New("cannot create rocky planet without obstacles")
	}
	if size.Area() < 2 {
		return nil, errors.New("size too small") // rover + 1 obstacle would not fit
	}
	obstacleList := obs.FromList(obstacles)
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

func (this *RockyPlanet) Color() string {
	return this.color
}
