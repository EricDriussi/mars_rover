package infra_test

import (
	"database/sql"
	"encoding/json"
	"errors"
	. "mars_rover/src/domain/planet"
	. "mars_rover/src/domain/rover"
	. "mars_rover/src/domain/rover/godModRover"
	. "mars_rover/src/infra/persistence"
	. "mars_rover/src/infra/persistence/entities"
	. "mars_rover/src/infra/persistence/mappers"
)

func saveGame(db *sql.DB, rover Rover, planet Planet) error {
	planetAsBytes, err := json.Marshal(MapToPersistencePlanet(planet))
	if err != nil {
		return err
	}

	planetInsertResult, err := db.Exec("INSERT INTO "+PlanetsTable+" (planet) VALUES (?)",
		string(planetAsBytes))
	if err != nil {
		return err
	}
	planetId, err := planetInsertResult.LastInsertId()
	if err != nil {
		return err
	}

	roverAsBytes, err := json.Marshal(MapToPersistenceRover(rover))
	if err != nil {
		return err
	}

	_, err = db.Exec("INSERT INTO "+RoversTable+" (id, rover, type, planet_id) VALUES (?, ?, ?, ?)",
		rover.Id().String(),
		string(roverAsBytes),
		typeOf(rover),
		planetId,
	)
	return err
}

func typeOf(rover Rover) string {
	if _, ok := rover.(*GodModRover); ok {
		return "godmod"
	}
	return "normal"
}

func savePlanet(db *sql.DB, planet Planet) (int, error) {
	planetAsBytes, err := json.Marshal(MapToPersistencePlanet(planet))
	if err != nil {
		return 0, err
	}

	planetInsertResult, err := db.Exec("INSERT INTO "+PlanetsTable+" (planet) VALUES (?)",
		string(planetAsBytes))
	if err != nil {
		return 0, err
	}
	num, err := planetInsertResult.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(num), nil
}

func getLastPersistedRover(db *sql.DB, planet Planet) (Rover, error) {
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
	if len(persistedRovers) > 1 {
		return nil, errors.New("more than one rover found")
	}
	foundRover, err := MapToDomainRover(persistedRovers[0], planet)
	if err != nil {
		return nil, err
	}
	return foundRover, nil
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
		var typeOf string
		var planetId int
		err := rows.Scan(&id, &rover, &typeOf, &planetId)
		if err != nil {
			return nil, err
		}
		var roverData RoverEntity
		err = json.Unmarshal([]byte(rover), &roverData)
		if err != nil {
			return nil, err
		}
		roverData.Type = typeOf
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
