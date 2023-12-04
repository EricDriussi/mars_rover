package create_test

import (
	"mars_rover/src/action/create"
	. "mars_rover/src/infra/test"
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
	repo.On("SaveGame").Return(nil)

	for i := 0; i < 25; i++ {
		rover, err := create.Random(repo)
		assert.Nil(t, err)
		assert.NotNil(t, rover)
	}
}
