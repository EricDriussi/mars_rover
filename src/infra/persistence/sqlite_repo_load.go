package persistence

import (
	"encoding/json"
	"errors"
	. "github.com/google/uuid"
	. "mars_rover/src/domain"
	. "mars_rover/src/infra/persistence/entities"
	. "mars_rover/src/infra/persistence/mappers"

	_ "github.com/mattn/go-sqlite3"
)

func (r *SQLiteRepository) LoadGame(id UUID) (GameDTO, error) {

	optionalRover, err := r.getRover(id)
	if err != nil {
		return GameDTO{}, err
	}
	if !optionalRover.Present {
		return GameDTO{}, errors.New("rover not found")
	}

	planetEntity, err := r.getPlanet(optionalRover.Value.PlanetId)
	if err != nil {
		return GameDTO{}, err
	}

	domainPlanet, err := MapToDomainPlanet(planetEntity)
	if err != nil {
		return GameDTO{}, err
	}

	domainRover, err := MapToDomainRover(optionalRover.Value, domainPlanet)
	if err != nil {
		return GameDTO{}, err
	}

	return GameDTO{
		Planet: domainPlanet,
		Rover:  domainRover,
	}, nil
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
