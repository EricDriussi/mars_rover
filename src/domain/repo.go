package domain

import (
	. "github.com/google/uuid"
	. "mars_rover/src/domain/planet"
	. "mars_rover/src/domain/rover"
)

type Repository interface {
	UpdateRover(rover Rover) error
	LoadGame(id UUID) (GameDTO, error) // TODO: divert tests and remove
	AddRover(rover Rover, planetId int64) error
	AddPlanet(planet Planet) (int64, error)
	GetRover(roverId UUID) (Rover, error)
	GetPlanet(roverId UUID) (Planet, error)
}

type GameDTO struct {
	Planet Planet
	Rover  Rover
}
