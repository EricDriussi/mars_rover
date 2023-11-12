package domain

import (
	"mars_rover/internal/domain/rover"
)

type Repository interface {
	saveRover(rover rover.Rover) error
}
