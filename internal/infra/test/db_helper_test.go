package infra_test

import (
	"database/sql"
	"encoding/json"
	"errors"
	. "mars_rover/internal/domain/planet"
	. "mars_rover/internal/domain/rover"
	. "mars_rover/internal/infra"
	. "mars_rover/internal/infra/entities"
	. "mars_rover/internal/infra/mappers"
)

func getLastPersistedRover(db *sql.DB) (Rover, error) {
	rows, err := db.Query("SELECT * FROM " + RoversTable)
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			panic("err closing db connection")
		}
	}(rows)
	if err != nil {
		return nil, err
	}

	persistedRovers, err := unmarshalRoverEntities(rows)
	if err != nil {
		return nil, err
	}
	foundRovers, err := MapToDomainRovers(persistedRovers)
	if err != nil {
		return nil, err
	}
	if len(foundRovers) > 1 {
		return nil, errors.New("more than one rover found")
	}
	return foundRovers[0], nil
}

func getLastPersistedPlanet(db *sql.DB) (Planet, error) {
	rows, err := db.Query("SELECT * FROM " + PlanetsTable)
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			panic("err closing db connection")
		}
	}(rows)
	if err != nil {
		return nil, err
	}

	persistedPlanets, err := unmarshalPlanetEntities(rows)

	if err != nil {
		return nil, err
	}
	foundPlanets, err := MapToDomainPlanets(persistedPlanets)
	if err != nil {
		return nil, err
	}
	if len(foundPlanets) > 1 {
		return nil, errors.New("more than one planet found")
	}
	return foundPlanets[0], nil
}

func unmarshalRoverEntities(rows *sql.Rows) ([]RoverEntity, error) {
	var listOfRovers []RoverEntity
	for rows.Next() {
		var id string
		var rover string
		var godMod bool
		var planetId int
		err := rows.Scan(&id, &rover, &godMod, &planetId)
		if err != nil {
			return nil, err

		}
		var roverData RoverEntity
		err = json.Unmarshal([]byte(rover), &roverData)
		if err != nil {
			return nil, err
		}
		roverData.GodMod = godMod
		roverData.PlanetId = planetId
		listOfRovers = append(listOfRovers, roverData)
	}
	return listOfRovers, nil
}

func unmarshalPlanetEntities(rows *sql.Rows) ([]PlanetEntity, error) {
	var planetEntities []PlanetEntity
	for rows.Next() {
		var id string
		var planet string
		err := rows.Scan(&id, &planet)
		if err != nil {
			return nil, err

		}
		var planetData PlanetEntity
		err = json.Unmarshal([]byte(planet), &planetData)
		if err != nil {
			return nil, err
		}
		planetEntities = append(planetEntities, planetData)
	}
	return planetEntities, nil
}
