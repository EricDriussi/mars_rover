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

// TODO: should not take an empty list, that would be an EmptyPlanet,
// hide behind a factory?
// TODO.LM: color is here just as an example of why you might want to
// decouple map from planet, but it's not used for anything
func Create(color string, size Size, obstacles []Obstacle) (*RockyPlanet, error) {
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
