package domain

import (
	. "mars_rover/src/domain/planet"
	. "mars_rover/src/domain/rover"
	. "mars_rover/src/domain/rover/id"
)

type Repository interface {
	UpdateRover(rover Rover) *RepositoryError
	AddRover(rover Rover, planetId int) *RepositoryError
	AddPlanet(planet Planet) (int, *RepositoryError)
	GetRover(roverId ID) (Rover, *RepositoryError)
	GetGame(roverId ID) (*Game, *RepositoryError)
}

// TODO.LM: This should have a proper constructor and private fields
// It's also not nice that this is used by the repo as well as the actions
// But I'm running out of time!!!!
type Game struct {
	Rover  Rover
	Planet Planet
}
