package infra_test

import (
	"encoding/json"
	"mars_rover/internal/domain/coordinate"
	"mars_rover/internal/domain/location"
	"mars_rover/internal/domain/location/direction"
	"mars_rover/internal/domain/obstacle"
	"mars_rover/internal/domain/obstacle/small_rock"
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

	repo := infra.NewSQLite(db)
	loc, _ := location.From(*coordinate.NewAbsolute(0, 0), &direction.North{})
	planet, _ := rockyPlanet.Create(size.Size{Width: 5, Height: 5}, []obstacle.Obstacle{small_rock.In(*coordinate.NewAbsolute(1, 1))})
	testRover, _ := rover.Land(*loc, planet)

	err := repo.SaveRover(testRover)
	assert.Nil(t, err)

	expectedNumberOfRovers := 1
	var actualNumberOfRovers int
	err = db.QueryRow("SELECT COUNT(*) FROM rovers").Scan(&actualNumberOfRovers)
	assert.Nil(t, err)
	assert.Equal(t, expectedNumberOfRovers, actualNumberOfRovers)

	var savedRover string
	err = db.QueryRow("SELECT rover FROM rovers LIMIT ?",
		expectedNumberOfRovers).Scan(&savedRover)
	assert.Nil(t, err)

	var roverData infra.RoverPersistenceEntity
	err = json.Unmarshal([]byte(savedRover), &roverData)
	assert.Nil(t, err)
	var rover rover.Rover
	rover, err = infra.ConvertToRover(roverData, planet)
	assert.Nil(t, err)
	assert.Equal(t, testRover, rover)
}
