package action_test

import (
	"mars_rover/src/action"
	. "mars_rover/src/test_helpers/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomCreationDoesNotPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("The code panicked: %v", r)
		}
	}()
	repo := new(MockRepo)
	repo.On("AddPlanet").Return(nil)
	repo.On("AddRover").Return(nil)
	act := action.For(repo)

	for i := 0; i < 25; i++ {
		rover, err := act.Random()
		assert.Nil(t, err)
		assert.NotNil(t, rover)
	}
}
