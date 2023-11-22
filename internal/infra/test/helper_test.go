package infra_test

import (
	"database/sql"
	"encoding/json"
	"errors"
	"mars_rover/internal/domain/coordinate/absoluteCoordinate"
	. "mars_rover/internal/domain/obstacle"
	"mars_rover/internal/domain/obstacle/bigRock"
	"mars_rover/internal/domain/obstacle/smallRock"
	. "mars_rover/internal/domain/planet"
	"mars_rover/internal/domain/planet/rockyPlanet"
	. "mars_rover/internal/domain/rover"
	"mars_rover/internal/domain/rover/godModRover"
	"mars_rover/internal/domain/rover/wrappingCollidingRover"
	s "mars_rover/internal/domain/size"
	. "mars_rover/internal/infra"
	. "mars_rover/internal/infra/entities"
	. "mars_rover/internal/infra/mappers"
	"reflect"
)

func getLastPersistedRover(db *sql.DB, planet Planet) (Rover, error) {
	rows, err := db.Query("SELECT * FROM " + RoversTable)
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)
	if err != nil {
		return nil, err
	}

	persistedRovers, err := unmarshalRoverEntity(rows)
	if err != nil {
		return nil, err
	}
	foundRovers, err := MapToDomainRovers(persistedRovers, planet)
	if err != nil {
		return nil, err
	}
	if len(foundRovers) > 1 {
		return nil, errors.New("more than one rover found")
	}
	return foundRovers[0], nil
}

func unmarshalRoverEntity(rows *sql.Rows) ([]RoverEntity, error) {
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

func getLastPersistedPlanet(db *sql.DB) (Planet, error) {
	persistedEntities, err := getAllPersistedEntities(db, PlanetsTable, reflect.TypeOf(PlanetEntity{}))
	if err != nil {
		return nil, err
	}
	foundPlanets, err := MapToDomainPlanets(persistedEntities.([]PlanetEntity))
	if err != nil {
		return nil, err
	}
	if len(foundPlanets) > 1 {
		return nil, errors.New("more than one planet found")
	}
	return foundPlanets[0], nil
}

func getAllPersistedEntities(db *sql.DB, tableName string, entityType reflect.Type) (interface{}, error) {
	var listOfEntities reflect.Value
	listOfEntities = reflect.MakeSlice(reflect.SliceOf(entityType), 0, 0)
	var entities []string
	rows, err := db.Query("SELECT * FROM " + tableName)
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var entity string
		var id int
		err := rows.Scan(&id, &entity)
		if err != nil {
			return nil, err
		}
		entities = append(entities, entity)
	}

	for _, entityString := range entities {
		entityValue := reflect.New(entityType)
		err := json.Unmarshal([]byte(entityString), entityValue.Interface())
		if err != nil {
			return nil, err
		}
		listOfEntities = reflect.Append(listOfEntities, entityValue.Elem())
	}
	return listOfEntities.Interface(), nil
}

func setupWrappingRoverOnRockyPlanet() (Rover, Planet) {
	rovCoord := absoluteCoordinate.From(0, 0)
	testPlanet := setupRockyPlanet()
	testRover, _ := wrappingCollidingRover.Land(*rovCoord, testPlanet)
	return testRover, testPlanet
}

func setupGodModRoverOnRockyPlanet() (Rover, Planet) {
	rovCoord := absoluteCoordinate.From(1, 1)
	testPlanet := setupRockyPlanet()
	testRover := godModRover.Land(*rovCoord, testPlanet)
	return testRover, testPlanet
}

func setupRockyPlanet() Planet {
	size, _ := s.Square(10)
	smallCoord := absoluteCoordinate.From(1, 1)
	testSmallRock := smallRock.In(*smallCoord)
	bigCoord1 := absoluteCoordinate.From(2, 2)
	bigCoord2 := absoluteCoordinate.From(2, 3)
	testBigRock := bigRock.In([]absoluteCoordinate.AbsoluteCoordinate{*bigCoord1, *bigCoord2})
	testPlanet, _ := rockyPlanet.Create("testColor", *size, []Obstacle{&testSmallRock, &testBigRock})
	return testPlanet
}
