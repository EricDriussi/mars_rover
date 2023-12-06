package create_test

import (
	"mars_rover/src/action/create"
	. "mars_rover/src/infra/persistence/test"
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
	action := create.For(repo)

	for i := 0; i < 25; i++ {
		rover, err := action.Random()
		assert.Nil(t, err)
		assert.NotNil(t, rover)
	}
}
