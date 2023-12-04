package infra_test

import (
	"database/sql"
	. "mars_rover/src/domain/planet"
	. "mars_rover/src/domain/rover"
	. "mars_rover/src/infra/persistence"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSavesRover(t *testing.T) {
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

			foundPlanet, err := getLastPersistedPlanet(db)
			assertPlanetsAreEqual(t, testPlanet, foundPlanet)
			foundRover, err := getLastPersistedRover(db, foundPlanet)
			assertRoversAreEqual(t, foundRover, testRover)
		})
	}
}
