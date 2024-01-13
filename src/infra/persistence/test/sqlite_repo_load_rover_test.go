package infra_test

import (
	. "mars_rover/src/domain/planet"
	. "mars_rover/src/domain/rover"
	"mars_rover/src/domain/rover/id"
	. "mars_rover/src/infra/persistence"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadsRoverWhenPresent(t *testing.T) {
	db, repo := InitMem()
	testCases := []struct {
		name      string
		setupFunc func(t *testing.T) (Rover, Planet)
	}{
		{
			name:      "wrapping rover",
			setupFunc: setupWrappingRover,
		},
		{
			name:      "god mod rover",
			setupFunc: setupGodModRover,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testRover, testPlanet := testCase.setupFunc(t)
			err := saveGame(db, testRover, testPlanet)
			assert.Nil(t, err)

			foundRover, repoErr := repo.GetRover(testRover.Id())

			assert.Nil(t, repoErr)
			assertRoversAreEqual(t, testRover, foundRover)
		})
	}
}

func TestDoesNotLoadRoverWhenNotPresent(t *testing.T) {
	db, repo := InitMem()
	testCases := []struct {
		name      string
		setupFunc func(t *testing.T) (Rover, Planet)
	}{
		{
			name:      "wrapping rover",
			setupFunc: setupWrappingRover,
		},
		{
			name:      "god mod rover",
			setupFunc: setupGodModRover,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			rover, planet := testCase.setupFunc(t)
			err := saveGame(db, rover, planet)
			assert.Nil(t, err)

			_, repoErr := repo.GetRover(id.New())

			assert.Error(t, repoErr)
		})
	}
}
