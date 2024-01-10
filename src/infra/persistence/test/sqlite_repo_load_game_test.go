package infra_test

import (
	. "mars_rover/src/domain/planet"
	. "mars_rover/src/domain/rover"
	"mars_rover/src/domain/rover/uuid"
	. "mars_rover/src/infra/persistence"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadsGameWhenPresent(t *testing.T) {
	testCases := []struct {
		name      string
		setupFunc func(t *testing.T) (Rover, Planet)
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
			testRover, testPlanet := testCase.setupFunc(t)
			err := saveGame(db, testRover, testPlanet)
			assert.Nil(t, err)

			game, repoErr := repo.GetGame(testRover.Id())

			assert.Nil(t, repoErr)
			assertPlanetsAreEqual(t, testPlanet, game.Planet)
			assertRoversAreEqual(t, testRover, game.Rover)
		})
	}
}

func TestDoesNotLoadGameWhenNotPresent(t *testing.T) {
	testCases := []struct {
		name      string
		setupFunc func(t *testing.T) (Rover, Planet)
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
			testRover, testPlanet := testCase.setupFunc(t)
			err := saveGame(db, testRover, testPlanet)
			assert.Nil(t, err)

			_, repoErr := repo.GetGame(uuid.New())

			assert.Error(t, repoErr)
		})
	}
}
