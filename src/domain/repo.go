package domain

import (
	. "mars_rover/src/domain/planet"
	. "mars_rover/src/domain/rover"
	. "mars_rover/src/domain/rover/uuid"
)

type Repository interface {
	UpdateRover(rover Rover) *RepositoryError
	AddRover(rover Rover, planetId int) *RepositoryError
	AddPlanet(planet Planet) (int, *RepositoryError)
	GetRover(roverId UUID) (Rover, *RepositoryError)
	GetGame(roverId UUID) (*Game, *RepositoryError)
}

type Game struct {
	Rover  Rover
	Planet Planet
}
