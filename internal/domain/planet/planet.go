package planet

import (
	"errors"
	"mars_rover/internal/domain/obstacle"
	"mars_rover/internal/domain/size"
)

type Planet interface {
	Size() size.Size
	Obstacles() []obstacle.Obstacle
}

type Mars struct {
	size      size.Size
	obstacles []obstacle.Obstacle
}

func Create(size size.Size, obstacles []obstacle.Obstacle) (*Mars, error) {
	for _, obs := range obstacles {
		if obs.IsBeyond(size) {
			return nil, errors.New("an obstacle was set outside of the planet :c")
		}
	}

	return &Mars{size, obstacles}, nil
}

func (this *Mars) Size() size.Size {
	return this.size
}

func (this *Mars) Obstacles() []obstacle.Obstacle {
	return this.obstacles
}
