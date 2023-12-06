package infra_test

import (
	. "mars_rover/src/domain/planet"
	. "mars_rover/src/domain/rover"
	. "mars_rover/src/infra/persistence"
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
			db, repo := InitMem()

			testRover, testPlanet := testCase.setupFunc()

			// TODO: don't use save in update test
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
