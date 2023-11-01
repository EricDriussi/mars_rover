package planet

import (
	"errors"
	"mars_rover/internal/domain/obstacle"
	"mars_rover/internal/domain/size"
)

type Planet struct {
	Size      size.Size
	Obstacles []obstacle.Obstacle
}

func Create(size size.Size, obstacles []obstacle.Obstacle) (*Planet, error) {
	for _, obs := range obstacles {
		if obs.IsBeyond(size) {
			return nil, errors.New("an obstacle was set outside of the planet :c")
		}
	}

	return &Planet{size, obstacles}, nil
}
