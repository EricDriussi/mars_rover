package infra_test

import (
	"database/sql"
	"encoding/json"
	"github.com/stretchr/testify/assert"
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
	"testing"
)

func getAllPersistedRovers(t *testing.T, db *sql.DB, planet Planet) []Rover {
	rows, err := db.Query("SELECT * FROM " + RoversTable)
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)
	assert.Nil(t, err)

	persistedRovers, err := unmarshalRoverEntity(rows)
	assert.Nil(t, err)
	foundRovers, err := MapToDomainRovers(persistedRovers, planet)
	assert.Nil(t, err)
	return foundRovers
}

func unmarshalRoverEntity(rows *sql.Rows) ([]RoverEntity, error) {
	var listOfRovers []RoverEntity
	for rows.Next() {
		var rover string
		var id int
		var godMod bool
		err := rows.Scan(&id, &rover, &godMod)
		if err != nil {
			return nil, err

		}
		var roverData RoverEntity
		err = json.Unmarshal([]byte(rover), &roverData)
		if err != nil {
			return nil, err
		}
		roverData.GodMod = godMod
		listOfRovers = append(listOfRovers, roverData)
	}
	return listOfRovers, nil
}

func getAllPersistedPlanets(t *testing.T, db *sql.DB) []Planet {
	persistedPlanets := getAllPersistedEntities(t, db, PlanetsTable, reflect.TypeOf(PlanetEntity{})).([]PlanetEntity)
	foundPlanets, err := MapToDomainPlanets(persistedPlanets)
	assert.Nil(t, err)
	return foundPlanets
}

func getAllPersistedEntities(t *testing.T, db *sql.DB, tableName string, entityType reflect.Type) interface{} {
	var listOfEntities reflect.Value
	listOfEntities = reflect.MakeSlice(reflect.SliceOf(entityType), 0, 0)
	var entities []string
	rows, err := db.Query("SELECT * FROM " + tableName)
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)
	assert.Nil(t, err)

	for rows.Next() {
		var entity string
		var id int
		err := rows.Scan(&id, &entity)
		assert.Nil(t, err)
		entities = append(entities, entity)
	}

	for _, entityString := range entities {
		entityValue := reflect.New(entityType)
		err := json.Unmarshal([]byte(entityString), entityValue.Interface())
		assert.Nil(t, err)
		listOfEntities = reflect.Append(listOfEntities, entityValue.Elem())
	}
	return listOfEntities.Interface()
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
