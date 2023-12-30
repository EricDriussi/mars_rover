package domain

import (
	. "mars_rover/src/domain/planet"
	. "mars_rover/src/domain/rover"
	. "mars_rover/src/domain/rover/uuid"
)

type Repository interface {
	UpdateRover(rover Rover) error
	AddRover(rover Rover, planetId int64) error
	AddPlanet(planet Planet) (int64, error)
	GetRover(roverId UUID) (Rover, error)
	GetPlanet(roverId UUID) (Planet, error)
}
