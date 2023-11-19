package infra_test

import (
	"database/sql"
	. "mars_rover/internal/infra"
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

	testRover, testPlanet := setupWrappingRoverOnRockyPlanet()
	repo := NewSQLite(db)
	err := repo.SaveRover(testRover)
	assert.Nil(t, err)

	foundRovers := getAllPersistedRovers(t, db, testPlanet)
	assert.Len(t, foundRovers, 1)
	foundRover := foundRovers[0]
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

	testPlanet := setupRockyPlanet()
	repo := NewSQLite(db)
	err := repo.SavePlanet(testPlanet)
	assert.Nil(t, err)

	foundPlanets := getAllPersistedPlanets(t, db)
	assert.Len(t, foundPlanets, 1)
	foundPlanet := foundPlanets[0]
	assert.Equal(t, testPlanet.Color(), foundPlanet.Color())
	assert.Equal(t, testPlanet.Obstacles(), foundPlanet.Obstacles())
	assert.Equal(t, testPlanet.Size(), foundPlanet.Size())
}
