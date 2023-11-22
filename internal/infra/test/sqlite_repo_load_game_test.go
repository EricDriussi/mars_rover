package infra_test

import (
	"database/sql"
	"github.com/google/uuid"
	. "mars_rover/internal/infra"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadGodModRoverGame(t *testing.T) {
	db := InitMem()
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic("err closing db connection")
		}
	}(db)
	testRover, testPlanet := setupGodModRoverOnRockyPlanet()
	repo := NewSQLite(db)
	err := repo.SaveGame(testRover, testPlanet)
	assert.Nil(t, err)

	gameDTO, err := repo.LoadGame(testRover.Id())
	assert.Nil(t, err)

	assertRoversAreEqual(t, testRover, gameDTO.Rover)
	assertPlanetsAreEqual(t, testPlanet, gameDTO.Planet)
}

func TestNotLoadGodModRoverGameWhenIdIsIncorrect(t *testing.T) {
	db := InitMem()
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic("err closing db connection")
		}
	}(db)
	testRover, testPlanet := setupGodModRoverOnRockyPlanet()
	repo := NewSQLite(db)
	err := repo.SaveGame(testRover, testPlanet)
	assert.Nil(t, err)

	_, err = repo.LoadGame(uuid.New())
	assert.NotNil(t, err)
}

func TestLoadWrappingRoverGame(t *testing.T) {
	db := InitMem()
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic("err closing db connection")
		}
	}(db)
	testRover, testPlanet := setupWrappingRoverOnRockyPlanet()
	repo := NewSQLite(db)
	err := repo.SaveGame(testRover, testPlanet)
	assert.Nil(t, err)

	gameDTO, err := repo.LoadGame(testRover.Id())
	assert.Nil(t, err)

	assertRoversAreEqual(t, testRover, gameDTO.Rover)
	assertPlanetsAreEqual(t, testPlanet, gameDTO.Planet)
}

func TestNotLoadWrappingRoverGameWhenIdIsIncorrect(t *testing.T) {
	db := InitMem()
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic("err closing db connection")
		}
	}(db)
	testRover, testPlanet := setupWrappingRoverOnRockyPlanet()
	repo := NewSQLite(db)
	err := repo.SaveGame(testRover, testPlanet)
	assert.Nil(t, err)

	_, err = repo.LoadGame(uuid.New())
	assert.NotNil(t, err)
}
