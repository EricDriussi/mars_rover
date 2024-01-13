package infra_test

import (
	. "mars_rover/src/domain/planet"
	. "mars_rover/src/domain/rover"
	. "mars_rover/src/infra/persistence"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdatesRover(t *testing.T) {
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

			directionBeforeMovement := rover.Direction()
			rover.TurnRight()
			repoErr := repo.UpdateRover(rover)

			assert.Nil(t, repoErr)
			foundRover, err := getLastPersistedRover(db, planet)
			assert.Nil(t, err)
			assertRoversAreEqual(t, foundRover, rover)
			assert.NotEqual(t, directionBeforeMovement.CardinalPoint(), foundRover.Direction().CardinalPoint())
		})
	}
}
