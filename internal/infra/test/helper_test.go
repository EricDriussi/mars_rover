package infra_test

import (
	"database/sql"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"mars_rover/internal/domain/coordinate/absoluteCoordinate"
	. "mars_rover/internal/domain/obstacle"
	"mars_rover/internal/domain/obstacle/smallRock"
	. "mars_rover/internal/domain/obstacle/smallRock"
	. "mars_rover/internal/domain/planet/rockyPlanet"
	. "mars_rover/internal/domain/rover/direction"
	. "mars_rover/internal/domain/rover/wrappingCollidingRover"
	s "mars_rover/internal/domain/size"
	. "mars_rover/internal/infra"
	"reflect"
	"testing"
)

func getAllPersistedRovers(t *testing.T, db *sql.DB) []RoverPersistenceEntity {
	return getAllPersistedEntities(t, db, WrappingRoversTable, reflect.TypeOf(RoverPersistenceEntity{})).([]RoverPersistenceEntity)
}

func getAllPersistedRockyPlanets(t *testing.T, db *sql.DB) []RockyPlanetPersistenceEntity {
	return getAllPersistedEntities(t, db, RockyPlanetsTable, reflect.TypeOf(RockyPlanetPersistenceEntity{})).([]RockyPlanetPersistenceEntity)
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

func aWrappingTestRover(planet RockyPlanet) WrappingCollidingRover {
	coordinate := absoluteCoordinate.From(0, 0)
	testRover, _ := LandFacing(North{}, *coordinate, &planet)
	return *testRover
}

func aSmallRockWithin(size int) SmallRock {
	coordinate := absoluteCoordinate.From(size-1, size-1)
	return smallRock.In(*coordinate)
}

func aRockyTestPlanetWith(size int, rock SmallRock) RockyPlanet {
	siz, _ := s.Square(size)
	planet, _ := Create("testColor", *siz, []Obstacle{&rock})
	return *planet
}
