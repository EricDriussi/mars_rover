package planetWithObstacles

import (
	"errors"
	. "mars_rover/src/domain/obstacle"
	. "mars_rover/src/domain/obstacle/obstacles"
	obs "mars_rover/src/domain/obstacle/obstacles"
	. "mars_rover/src/domain/size"
)

type PlanetWithObstacles struct {
	color     string
	size      Size
	obstacles Obstacles
}

func Create(color string, size Size, obstacles []Obstacle) (*PlanetWithObstacles, error) {
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

	return &PlanetWithObstacles{color, size, *obstacleList}, nil
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
