package persistence

import (
	"encoding/json"
	. "mars_rover/src/domain"
	. "mars_rover/src/domain/planet"
	. "mars_rover/src/domain/rover"
	. "mars_rover/src/domain/rover/uuid"
	. "mars_rover/src/infra/persistence/entities"
	. "mars_rover/src/infra/persistence/mappers"

	_ "github.com/mattn/go-sqlite3"
)

func (this *SQLiteRepository) GetRover(roverId UUID) (Rover, *RepositoryError) {
	_, rover, err := this.getGame(roverId)
	if err != nil {
		return nil, err
	}
	return rover, nil
}

func (this *SQLiteRepository) GetGame(roverId UUID) (*Game, *RepositoryError) {
	planet, rover, err := this.getGame(roverId)
	if err != nil {
		return nil, err
	}
	return &Game{
		Rover:  rover,
		Planet: planet,
	}, nil
}

// TODO.LM: This is an awful coupling between rover and planet.
// I designed the domain so that a rover cannot be instantiated without a planet.
// Since there is no circumstance in which a planet needs to be retrieved without its rover, the planet domain entity
// has no ID of its own, depending on the rover entity's planetID
// This seems to make logical sense (to me), but produces this weird scenario in which both need to be retrieved together.
// Hopefully there is a better way to do this that does not force a planetID into the domain layer.
func (this *SQLiteRepository) getGame(roverId UUID) (Planet, Rover, *RepositoryError) {
	roverEntity, repoErr := this.getRover(roverId)
	if repoErr != nil {
		return nil, nil, repoErr
	}

	planetEntity, repoErr := this.getPlanet(roverEntity.PlanetId)
	if repoErr != nil {
		return nil, nil, repoErr
	}

	domainPlanet, err := MapToDomainPlanet(*planetEntity)
	if err != nil {
		return nil, nil, CouldNotMap(err)
	}

	domainRover, err := MapToDomainRover(*roverEntity, domainPlanet)
	if err != nil {
		return nil, nil, CouldNotMap(err)
	}

	return domainPlanet, domainRover, nil
}

func (this *SQLiteRepository) getRover(roverId UUID) (*RoverEntity, *RepositoryError) {
	row := this.db.QueryRow("SELECT * FROM "+RoversTable+" WHERE id = ?", roverId.String())

	var id string
	var rover string
	var typeOf string
	var planetId int
	err := row.Scan(&id, &rover, &typeOf, &planetId)
	if err != nil {
		return nil, NotFound()
	}
	var roverEntity RoverEntity
	err = json.Unmarshal([]byte(rover), &roverEntity)
	if err != nil {
		return nil, CouldNotMap(err)
	}
	roverEntity.Type = typeOf
	roverEntity.PlanetId = planetId
	return &roverEntity, nil
}

func (this *SQLiteRepository) getPlanet(planetId int) (*PlanetEntity, *RepositoryError) {
	row := this.db.QueryRow("SELECT * FROM "+PlanetsTable+" WHERE id = ?", planetId)

	var id string
	var planet string
	err := row.Scan(&id, &planet)
	if err != nil {
		return nil, NotFound()
	}
	var planetEntity PlanetEntity
	err = json.Unmarshal([]byte(planet), &planetEntity)
	if err != nil {
		return nil, CouldNotMap(err)
	}
	return &planetEntity, nil
}
