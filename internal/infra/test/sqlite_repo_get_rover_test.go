package infra_test

import (
	"database/sql"
	"github.com/google/uuid"
	. "mars_rover/internal/domain/planet"
	. "mars_rover/internal/domain/rover"
	. "mars_rover/internal/infra"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetsRover(t *testing.T) {
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

			foundRover, err := repo.GetRover(testRover.Id())
			assert.NotNil(t, foundRover)
			assert.Nil(t, err)

			assertRoversAreEqual(t, testRover, foundRover)
		})
	}
}

func TestDoesNotGetRover(t *testing.T) {
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

			foundRover, err := repo.GetRover(uuid.New())
			assert.Nil(t, foundRover)
			assert.Nil(t, err)
		})
	}
}
