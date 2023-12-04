package emptyPlanet

import (
	. "mars_rover/src/domain/obstacle/obstacles"
	. "mars_rover/src/domain/size"
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

func (this *EmptyPlanet) Color() string {
	return this.color
}
