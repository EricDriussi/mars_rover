package bounded_random_creator_test

import (
	random_creator "mars_rover/src/action/createRandom/bounded"
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
	act := random_creator.With(repo)

	for i := 0; i < 25; i++ {
		rover, err := act.Create()
		assert.Nil(t, err)
		assert.NotNil(t, rover)
	}
}
