package domain

import (
	. "github.com/google/uuid"
	. "mars_rover/internal/domain/planet"
	. "mars_rover/internal/domain/rover"
	"mars_rover/internal/infra/entities"
)

type Repository interface {
	SaveRover(rover Rover) error
	SavePlanet(planet Planet) error
	LoadGame(id UUID) (entities.GameDTO, error)
}
