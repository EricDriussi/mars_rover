package infra_test

import (
	"database/sql"
	"mars_rover/internal/domain/planet"
	"mars_rover/internal/domain/rover"
	. "mars_rover/internal/infra"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSaveGodModRoverGame(t *testing.T) {
	db := InitMem()
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)
	testRover, testPlanet := setupGodModRoverOnRockyPlanet()
	repo := NewSQLite(db)

	err := repo.SaveGame(testRover, testPlanet)
	assert.Nil(t, err)

	foundRover, err := getLastPersistedRover(db, testPlanet)
	assertRoversAreEqual(t, foundRover, testRover)
	foundPlanet, err := getLastPersistedPlanet(db)
	assertPlanetsAreEqual(t, testPlanet, foundPlanet)
}

func TestSaveWrappingRoverGame(t *testing.T) {
	db := InitMem()
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)
	testRover, testPlanet := setupWrappingRoverOnRockyPlanet()
	repo := NewSQLite(db)

	err := repo.SaveGame(testRover, testPlanet)
	assert.Nil(t, err)

	foundRover, err := getLastPersistedRover(db, testPlanet)
	assertRoversAreEqual(t, foundRover, testRover)
	foundPlanet, err := getLastPersistedPlanet(db)
	assertPlanetsAreEqual(t, testPlanet, foundPlanet)
}

func TestUpdateWrappingRover(t *testing.T) {
	db := InitMem()
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)
	testRover, testPlanet := setupWrappingRoverOnRockyPlanet()
	repo := NewSQLite(db)

	err := repo.SaveGame(testRover, testPlanet)
	assert.Nil(t, err)
	testRover.TurnRight()
	err = repo.UpdateRover(testRover)
	assert.Nil(t, err)

	foundRover, err := getLastPersistedRover(db, testPlanet)
	assertRoversAreEqual(t, foundRover, testRover)
}

func TestUpdateGodModRover(t *testing.T) {
	db := InitMem()
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)
	testRover, testPlanet := setupGodModRoverOnRockyPlanet()
	repo := NewSQLite(db)

	err := repo.SaveGame(testRover, testPlanet)
	assert.Nil(t, err)
	testRover.TurnRight()
	err = repo.UpdateRover(testRover)
	assert.Nil(t, err)

	foundRover, err := getLastPersistedRover(db, testPlanet)
	assertRoversAreEqual(t, foundRover, testRover)
}

func assertPlanetsAreEqual(t *testing.T, testPlanet planet.Planet, foundPlanet planet.Planet) {
	assert.Equal(t, testPlanet.Color(), foundPlanet.Color())
	assert.Equal(t, testPlanet.Obstacles(), foundPlanet.Obstacles())
	assert.Equal(t, testPlanet.Size(), foundPlanet.Size())
}

func assertRoversAreEqual(t *testing.T, foundRover rover.Rover, testRover rover.Rover) {
	assert.Equal(t, testRover.Coordinate(), foundRover.Coordinate())
	assert.Equal(t, testRover.Direction().CardinalPoint(), foundRover.Direction().CardinalPoint())
	assert.Equal(t, testRover.Map(), foundRover.Map())
}
