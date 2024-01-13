package infra_test

import (
	. "mars_rover/src/domain/planet"
	. "mars_rover/src/domain/rover"
	. "mars_rover/src/infra/persistence"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddsPlanet(t *testing.T) {
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
			_, planet := testCase.setupFunc(t)

			_, repoErr := repo.AddPlanet(planet)

			assert.Nil(t, repoErr)
			foundPlanet, err := getLastPersistedPlanet(db)
			assert.Nil(t, err)
			assertPlanetsAreEqual(t, planet, foundPlanet)
		})
	}
}
