package infra_test

import (
	. "mars_rover/src/domain/planet"
	. "mars_rover/src/domain/rover"
	"mars_rover/src/domain/rover/uuid"
	. "mars_rover/src/infra/persistence"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadsRover(t *testing.T) {
	testCases := []struct {
		name      string
		setupFunc func(t *testing.T) (Rover, Planet)
	}{
		{
			name:      "wrapping rover",
			setupFunc: setupWrappingRoverOnRockyPlanet,
		},
		{
			name:      "god mod rover",
			setupFunc: setupGodModRoverOnRockyPlanet,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			db, repo := InitMem()
			testRover, testPlanet := testCase.setupFunc(t)
			err := saveGame(db, testRover, testPlanet)
			assert.Nil(t, err)

			foundRover, err := repo.GetRover(testRover.Id())

			assert.Nil(t, err)
			assertRoversAreEqual(t, testRover, foundRover)
		})
	}
}

func TestDoesNotLoadRover(t *testing.T) {
	testCases := []struct {
		name      string
		setupFunc func(t *testing.T) (Rover, Planet)
	}{
		{
			name:      "wrapping rover",
			setupFunc: setupWrappingRoverOnRockyPlanet,
		},
		{
			name:      "god mod rover",
			setupFunc: setupGodModRoverOnRockyPlanet,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			db, repo := InitMem()
			testRover, testPlanet := testCase.setupFunc(t)
			err := saveGame(db, testRover, testPlanet)
			assert.Nil(t, err)

			_, err = repo.GetRover(uuid.New())

			assert.NotNil(t, err)
		})
	}
}
