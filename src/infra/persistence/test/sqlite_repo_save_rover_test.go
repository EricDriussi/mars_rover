package infra_test

import (
	. "mars_rover/src/domain/planet"
	. "mars_rover/src/domain/rover"
	. "mars_rover/src/infra/persistence"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddsRover(t *testing.T) {
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
			planetId, err := savePlanet(db, planet)
			assert.Nil(t, err)

			repoErr := repo.AddRover(rover, planetId)

			assert.Nil(t, repoErr)
			foundPlanet, err := getLastPersistedPlanet(db)
			assert.Nil(t, err)
			foundRover, err := getLastPersistedRover(db, foundPlanet)
			assert.Nil(t, err)
			assertRoversAreEqual(t, foundRover, rover)
		})
	}
}
