package emptyPlanet

import (
	. "mars_rover/internal/domain/obstacle/obstacles"
	. "mars_rover/internal/domain/size"
)

type EmptyPlanet struct {
	color string
	size  Size
}

func Create(color string, size Size) (*EmptyPlanet, error) {
	return &EmptyPlanet{color, size}, nil
}

func (this *EmptyPlanet) Size() Size {
	return this.size
}

func (this *EmptyPlanet) Obstacles() Obstacles {
	return Obstacles{}
}
