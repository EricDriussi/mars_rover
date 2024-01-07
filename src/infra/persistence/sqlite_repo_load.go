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
	optionalRover, err := this.getRover(roverId)
	if err != nil {
		return nil, PersistenceMalfunction(err)
	}
	if !optionalRover.Present {
		return nil, NotFound()
	}

	domainPlanet, getErr := this.GetPlanet(roverId)
	if getErr != nil {
		return nil, NotFound()
	}

	domainRover, err := MapToDomainRover(optionalRover.Value, domainPlanet)
	if err != nil {
		return nil, CouldNotMap(err)
	}

	return domainRover, nil
}

func (this *SQLiteRepository) GetPlanet(roverId UUID) (Planet, *RepositoryError) {
	optionalRover, err := this.getRover(roverId)
	if err != nil {
		return nil, PersistenceMalfunction(err)
	}
	if !optionalRover.Present {
		return nil, NotFound()
	}

	planetEntity, err := this.getPlanet(optionalRover.Value.PlanetId)
	if err != nil {
		return nil, NotFound()
	}

	domainPlanet, err := MapToDomainPlanet(planetEntity)
	if err != nil {
		return nil, CouldNotMap(err)
	}

	return domainPlanet, nil
}

// TODO: why the OptionalRover?
func (this *SQLiteRepository) getRover(roverId UUID) (OptionalRover, error) {
	row := this.db.QueryRow("SELECT * FROM "+RoversTable+" WHERE id = ?", roverId.String())

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

func (this *SQLiteRepository) getPlanet(planetId int) (PlanetEntity, error) {
	row := this.db.QueryRow("SELECT * FROM "+PlanetsTable+" WHERE id = ?", planetId)

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
