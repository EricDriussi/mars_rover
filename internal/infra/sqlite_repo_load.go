package infra

import (
	"encoding/json"
	. "github.com/google/uuid"
	. "mars_rover/internal/domain"
	. "mars_rover/internal/infra/entities"
	. "mars_rover/internal/infra/mappers"

	_ "github.com/mattn/go-sqlite3"
)

func (r *SQLiteRepository) LoadGame(id UUID) (GameDTO, error) {
	roverEntity, err := r.getRover(id)
	if err != nil {
		return GameDTO{}, err
	}

	planetEntity, err := r.getPlanet(roverEntity.PlanetId)
	if err != nil {
		return GameDTO{}, err
	}

	domainPlanet, err := MapToDomainPlanet(planetEntity)
	if err != nil {
		return GameDTO{}, err
	}

	domainRover, err := MapToDomainRover(roverEntity, domainPlanet)
	if err != nil {
		return GameDTO{}, err
	}

	return GameDTO{
		Planet: domainPlanet,
		Rover:  domainRover,
	}, nil
}

func (r *SQLiteRepository) getRover(roverId UUID) (RoverEntity, error) {
	row := r.db.QueryRow("SELECT * FROM "+RoversTable+" WHERE id = ?", roverId.String())

	var id string
	var rover string
	var godMod bool
	var planetId int
	err := row.Scan(&id, &rover, &godMod, &planetId)
	if err != nil {
		return RoverEntity{}, err
	}
	var roverEntity RoverEntity
	err = json.Unmarshal([]byte(rover), &roverEntity)
	if err != nil {
		return RoverEntity{}, err
	}
	roverEntity.GodMod = godMod
	roverEntity.PlanetId = planetId
	return roverEntity, nil
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
