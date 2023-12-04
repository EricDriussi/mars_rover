package infra_test

import (
	"database/sql"
	"github.com/google/uuid"
	. "mars_rover/src/domain/planet"
	. "mars_rover/src/domain/rover"
	. "mars_rover/src/infra/persistence"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadsGame(t *testing.T) {
	testCases := []struct {
		name      string
		setupFunc func() (Rover, Planet)
	}{
		{
			name:      "wrapping rover on rocky planet",
			setupFunc: setupWrappingRoverOnRockyPlanet,
		},
		{
			name:      "god mod rover on rocky planet",
			setupFunc: setupGodModRoverOnRockyPlanet,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			db, repo := InitMem()
			defer func(db *sql.DB) {
				err := db.Close()
				if err != nil {
					panic("err closing db connection")
				}
			}(db)
			testRover, testPlanet := testCase.setupFunc()

			err := repo.SaveGame(testRover, testPlanet)
			assert.Nil(t, err)

			gameDTO, err := repo.LoadGame(testRover.Id())
			assert.Nil(t, err)

			assertRoversAreEqual(t, testRover, gameDTO.Rover)
			assertPlanetsAreEqual(t, testPlanet, gameDTO.Planet)
		})
	}
}

func TestDoesNotLoadGame(t *testing.T) {
	testCases := []struct {
		name      string
		setupFunc func() (Rover, Planet)
	}{
		{
			name:      "wrapping rover on rocky planet",
			setupFunc: setupWrappingRoverOnRockyPlanet,
		},
		{
			name:      "god mod rover on rocky planet",
			setupFunc: setupGodModRoverOnRockyPlanet,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			db, repo := InitMem()
			defer func(db *sql.DB) {
				err := db.Close()
				if err != nil {
					panic("err closing db connection")
				}
			}(db)
			testRover, testPlanet := testCase.setupFunc()

			err := repo.SaveGame(testRover, testPlanet)
			assert.Nil(t, err)

			_, err = repo.LoadGame(uuid.New())
			assert.NotNil(t, err)
		})
	}
}
