package domain

import (
	. "github.com/google/uuid"
	. "mars_rover/src/domain/planet"
	. "mars_rover/src/domain/rover"
)

type Repository interface {
	UpdateRover(rover Rover) error
	SaveGame(rover Rover, planet Planet) error
	LoadGame(id UUID) (GameDTO, error)
}

type GameDTO struct {
	Planet Planet
	Rover  Rover
}
