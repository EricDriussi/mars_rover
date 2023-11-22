package infra_test

import (
	"database/sql"
	. "mars_rover/internal/domain/planet"
	. "mars_rover/internal/domain/rover"
	. "mars_rover/internal/infra"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdatesRover(t *testing.T) {
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
			db := InitMem()
			defer func(db *sql.DB) {
				err := db.Close()
				if err != nil {
					panic("err closing db connection")
				}
			}(db)
			testRover, testPlanet := testCase.setupFunc()
			repo := NewSQLite(db)

			err := repo.SaveGame(testRover, testPlanet)
			assert.Nil(t, err)
			testRover.TurnRight()
			err = repo.UpdateRover(testRover)
			assert.Nil(t, err)

			foundRover, err := getLastPersistedRover(db, testPlanet)
			assertRoversAreEqual(t, foundRover, testRover)
		})
	}
}
