package persistence

import (
	"encoding/json"
	"errors"
	. "github.com/google/uuid"
	. "mars_rover/src/domain/planet"
	. "mars_rover/src/domain/rover"
	. "mars_rover/src/infra/persistence/entities"
	. "mars_rover/src/infra/persistence/mappers"

	_ "github.com/mattn/go-sqlite3"
)

func (r *SQLiteRepository) GetRover(roverId UUID) (Rover, error) {
	optionalRover, err := r.getRover(roverId)
	if err != nil {
		return nil, err
	}
	if !optionalRover.Present {
		return nil, errors.New("rover not found")
	}

	domainPlanet, err := r.GetPlanet(roverId)
	if err != nil {
		return nil, err
	}

	domainRover, err := MapToDomainRover(optionalRover.Value, domainPlanet)
	if err != nil {
		return nil, err
	}

	return domainRover, nil
}

func (r *SQLiteRepository) GetPlanet(roverId UUID) (Planet, error) {
	optionalRover, err := r.getRover(roverId)
	if err != nil {
		return nil, err
	}
	if !optionalRover.Present {
		return nil, errors.New("rover not found")
	}

	planetEntity, err := r.getPlanet(optionalRover.Value.PlanetId)
	if err != nil {
		return nil, err
	}

	domainPlanet, err := MapToDomainPlanet(planetEntity)
	if err != nil {
		return nil, err
	}

	return domainPlanet, nil
}

// TODO: why the OptionalRover?
func (r *SQLiteRepository) getRover(roverId UUID) (OptionalRover, error) {
	row := r.db.QueryRow("SELECT * FROM "+RoversTable+" WHERE id = ?", roverId.String())

	var id string
	var rover string
	var godMod bool
	var planetId int
	err := row.Scan(&id, &rover, &godMod, &planetId)
	if err != nil {
		return OptionalRover{
			Value:   RoverEntity{},
			Present: false,
		}, nil
	}
	var roverEntity RoverEntity
	err = json.Unmarshal([]byte(rover), &roverEntity)
	if err != nil {
		return OptionalRover{
			Value:   RoverEntity{},
			Present: false,
		}, err
	}
	roverEntity.GodMod = godMod
	roverEntity.PlanetId = planetId
	return OptionalRover{
		Value:   roverEntity,
		Present: true,
	}, nil
}

func (r *SQLiteRepository) getPlanet(planetId int) (PlanetEntity, error) {
	row := r.db.QueryRow("SELECT * FROM "+PlanetsTable+" WHERE id = ?", planetId)

	var id string
	var planet string
	err := row.Scan(&id, &planet)
	if err != nil {
		return PlanetEntity{}, err
	}
	var planetEntity PlanetEntity
	err = json.Unmarshal([]byte(planet), &planetEntity)
	if err != nil {
		return PlanetEntity{}, err
	}
	return planetEntity, nil
}
