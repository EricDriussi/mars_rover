package emptyPlanet

import (
	"mars_rover/internal/domain/obstacle/obstacles"
	"mars_rover/internal/domain/size"
)

type EmptyPlanet struct {
	color string
	size  size.Size
}

func Create(color string, size size.Size) (*EmptyPlanet, error) {
	return &EmptyPlanet{color, size}, nil
}

func (this *EmptyPlanet) Size() size.Size {
	return this.size
}

func (this *EmptyPlanet) Obstacles() obstacles.Obstacles {
	return obstacles.Obstacles{}
}
