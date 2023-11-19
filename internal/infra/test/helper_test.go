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
	"testing"
)

func getAllPersistedRovers(t *testing.T, db *sql.DB) []RoverPersistenceEntity {
	var listOfRovers []RoverPersistenceEntity
	var rovers []string
	rows, err := db.Query("SELECT rover FROM wrapping_rovers")
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)
	assert.Nil(t, err)

	for rows.Next() {
		var rover string
		err := rows.Scan(&rover)
		assert.Nil(t, err)
		rovers = append(rovers, rover)
	}

	for _, roverString := range rovers {
		var roverData RoverPersistenceEntity
		err := json.Unmarshal([]byte(roverString), &roverData)
		assert.Nil(t, err)
		listOfRovers = append(listOfRovers, roverData)
	}
	return listOfRovers
}

func getAllPersistedRockyPlanets(t *testing.T, db *sql.DB) []RockyPlanetPersistenceEntity {
	var listOfRockyPlanets []RockyPlanetPersistenceEntity
	var rockyPlanets []string
	rows, err := db.Query("SELECT planet FROM rocky_planets")
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)
	assert.Nil(t, err)

	for rows.Next() {
		var planet string
		err := rows.Scan(&planet)
		assert.Nil(t, err)
		rockyPlanets = append(rockyPlanets, planet)
	}

	for _, rockyPlanetString := range rockyPlanets {
		var rockyPlanetData RockyPlanetPersistenceEntity
		err := json.Unmarshal([]byte(rockyPlanetString), &rockyPlanetData)
		assert.Nil(t, err)
		listOfRockyPlanets = append(listOfRockyPlanets, rockyPlanetData)
	}
	return listOfRockyPlanets
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
