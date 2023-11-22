package infra_test

import (
	"database/sql"
	. "mars_rover/internal/infra"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdateWrappingRover(t *testing.T) {
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
			panic("err closing db connection")
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
