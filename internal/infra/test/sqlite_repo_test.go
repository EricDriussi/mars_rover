package infra_test

import (
	"database/sql"
	"encoding/json"
	"mars_rover/internal/domain/coordinate/absoluteCoordinate"
	"mars_rover/internal/domain/location"
	"mars_rover/internal/domain/location/direction"
	"mars_rover/internal/domain/obstacle"
	"mars_rover/internal/domain/obstacle/small_rock"
	"mars_rover/internal/domain/planet"
	rockyPlanet "mars_rover/internal/domain/planet/rocky_planet"
	"mars_rover/internal/domain/rover"
	"mars_rover/internal/domain/size"
	"mars_rover/internal/infra"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSaveRover(t *testing.T) {
	db := infra.InitMem()
	defer db.Close()

	planet := aTestPlanet()
	testRover := aTestRover(planet)

	repo := infra.NewSQLite(db)
	err := repo.SaveRover(testRover)
	assert.Nil(t, err)

	expectedNumberOfRovers := 1

	actualNumberOfRovers := getAllPersistenceRovers(t, db)
	assert.Equal(t, expectedNumberOfRovers, len(actualNumberOfRovers))

	savedPersistanceRover := actualNumberOfRovers[0]

	var foundRover rover.Rover
	foundRover, err = mapToDomainRover(savedPersistanceRover, planet)
	assert.Nil(t, err)
	assert.Equal(t, testRover, foundRover)
}

func getAllPersistenceRovers(t *testing.T, db *sql.DB) []infra.RoverPersistenceEntity {
	var listOfRovers []infra.RoverPersistenceEntity
	var rovers []string
	rows, err := db.Query("SELECT rover FROM rovers")
	defer rows.Close()
	assert.Nil(t, err)

	for rows.Next() {
		var rover string
		err := rows.Scan(&rover)
		assert.Nil(t, err)
		rovers = append(rovers, rover)
	}

	for _, roverString := range rovers {
		var roverData infra.RoverPersistenceEntity
		err := json.Unmarshal([]byte(roverString), &roverData)
		assert.Nil(t, err)
		listOfRovers = append(listOfRovers, roverData)
	}
	return listOfRovers
}

func aTestRover(planet planet.Planet) rover.Rover {
	loc, _ := location.From(*absoluteCoordinate.From(0, 0), &direction.North{})
	testRover, _ := rover.Land(*loc, planet)
	return testRover
}

func aTestPlanet() planet.Planet {
	planet, _ := rockyPlanet.Create(size.Size{Width: 5, Height: 5}, []obstacle.Obstacle{small_rock.In(*absoluteCoordinate.From(1, 1))})
	return planet
}
