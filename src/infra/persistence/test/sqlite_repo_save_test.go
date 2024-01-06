package infra_test

import (
	. "mars_rover/src/domain/planet"
	. "mars_rover/src/domain/rover"
	. "mars_rover/src/infra/persistence"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddsPlanet(t *testing.T) {
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
			_, testPlanet := testCase.setupFunc(t)

			_, err := repo.AddPlanet(testPlanet)

			assert.Nil(t, err)
			foundPlanet, err := getLastPersistedPlanet(db)
			assertPlanetsAreEqual(t, testPlanet, foundPlanet)
		})
	}
}

func TestAddsRover(t *testing.T) {
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

			planetId, err := savePlanet(db, testPlanet)
			assert.Nil(t, err)
			err = repo.AddRover(testRover, planetId)
			assert.Nil(t, err)

			foundPlanet, err := getLastPersistedPlanet(db)
			assertPlanetsAreEqual(t, testPlanet, foundPlanet)
			foundRover, err := getLastPersistedRover(db, foundPlanet)
			assertRoversAreEqual(t, foundRover, testRover)
		})
	}
}
