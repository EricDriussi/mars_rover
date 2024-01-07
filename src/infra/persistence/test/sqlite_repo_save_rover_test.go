package infra_test

import (
	. "mars_rover/src/domain/planet"
	. "mars_rover/src/domain/rover"
	. "mars_rover/src/infra/persistence"
	"testing"

	"github.com/stretchr/testify/assert"
)

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

			repoErr := repo.AddRover(testRover, planetId)

			assert.Nil(t, repoErr)
			foundPlanet, err := getLastPersistedPlanet(db)
			assert.Nil(t, err)
			foundRover, err := getLastPersistedRover(db, foundPlanet)
			assert.Nil(t, err)
			assertRoversAreEqual(t, foundRover, testRover)
		})
	}
}
