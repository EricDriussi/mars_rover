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

			_, repoErr := repo.AddPlanet(testPlanet)

			assert.Nil(t, repoErr)
			foundPlanet, err := getLastPersistedPlanet(db)
			assert.Nil(t, err)
			assertPlanetsAreEqual(t, testPlanet, foundPlanet)
		})
	}
}
