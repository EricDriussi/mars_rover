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

			_, addErr := repo.AddPlanet(testPlanet)

			assert.Nil(t, addErr)
			foundPlanet, err := getLastPersistedPlanet(db)
			assert.Nil(t, err)
			assertPlanetsAreEqual(t, testPlanet, foundPlanet)
		})
	}
}

func TestDoesNotAddPlanetWhenAlreadyPresent(t *testing.T) {
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

			_, addErr := repo.AddPlanet(testPlanet)
			assert.Nil(t, addErr)
			_, addErr = repo.AddPlanet(testPlanet)

			assert.Error(t, addErr)
			assert.True(t, addErr.IsAlreadyExists())
			num, err := getNumberOfPlanets(db)
			assert.Nil(t, err)
			assert.Equal(t, 1, num)
		})
	}
}

func TestReturnPlanetIdWhenAlreadyPresent(t *testing.T) {
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
			_, repo := InitMem()
			_, testPlanet := testCase.setupFunc(t)

			id, err := repo.AddPlanet(testPlanet)
			assert.Nil(t, err)
			sameId, _ := repo.AddPlanet(testPlanet)

			assert.Equal(t, id, sameId)
		})
	}
}
