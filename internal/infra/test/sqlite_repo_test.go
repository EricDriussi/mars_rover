package infra_test

import (
	"database/sql"
	. "mars_rover/internal/domain/planet"
	. "mars_rover/internal/domain/rover"
	. "mars_rover/internal/infra"
	. "mars_rover/internal/infra/mappers"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSaveRover(t *testing.T) {
	db := InitMem()
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	size := 5
	testRock := aSmallRockWithin(size)
	testPlanet := aRockyTestPlanetWith(size, testRock)
	testRover := aWrappingTestRover(testPlanet)

	repo := NewSQLite(db)
	err := repo.SaveWrappingRover(testRover)
	assert.Nil(t, err)

	persistedRovers := getAllPersistedRovers(t, db)
	assert.Len(t, persistedRovers, 1)
	savedPersistenceRover := persistedRovers[0]
	var foundRover Rover
	foundRover, err = MapToDomainRover(savedPersistenceRover, &testPlanet)
	assert.Nil(t, err)

	assert.Equal(t, testRover.Coordinate(), foundRover.Coordinate())
	assert.Equal(t, testRover.Direction().CardinalPoint(), foundRover.Direction().CardinalPoint())
	assert.Equal(t, testRover.Map(), foundRover.Map())
}

func TestSavePlanet(t *testing.T) {
	db := InitMem()
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	size := 5
	testRock := aSmallRockWithin(size)
	testPlanet := aRockyTestPlanetWith(size, testRock)

	repo := NewSQLite(db)
	err := repo.SaveRockyPlanet(testPlanet)
	assert.Nil(t, err)

	persistedRockyPlanets := getAllPersistedRockyPlanets(t, db)
	assert.Len(t, persistedRockyPlanets, 1)
	savedPersistenceRover := persistedRockyPlanets[0]
	var foundPlanet Planet
	foundPlanet, err = MapToDomainPlanet(savedPersistenceRover)
	assert.Nil(t, err)

	assert.Equal(t, testPlanet.Color(), foundPlanet.Color())
	assert.Equal(t, testPlanet.Obstacles(), foundPlanet.Obstacles())
	assert.Equal(t, testPlanet.Size(), foundPlanet.Size())
}
