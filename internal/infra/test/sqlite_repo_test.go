package infra_test

import (
	"database/sql"
	. "mars_rover/internal/domain/rover"
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

	testPlanet := aTestPlanet()
	testRover := aTestRover(testPlanet)

	repo := NewSQLite(db)
	err := repo.SaveRover(testRover)
	assert.Nil(t, err)

	actualNumberOfRovers := getAllPersistedRovers(t, db)
	assert.Len(t, actualNumberOfRovers, 1)

	savedPersistenceRover := actualNumberOfRovers[0]

	var foundRover Rover
	foundRover, err = MapToDomainRover(savedPersistenceRover, testPlanet)
	assert.Nil(t, err)
	assert.Equal(t, testRover.Coordinate(), foundRover.Coordinate())
	assert.Equal(t, testRover.Direction().CardinalPoint(), foundRover.Direction().CardinalPoint())
	assert.Equal(t, testRover.Map(), foundRover.Map())
}
