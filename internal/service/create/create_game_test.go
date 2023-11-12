package create_test

import (
	"mars_rover/internal/service/create"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomCreationDoesNotPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("The code panicked: %v", r)
		}
	}()

	for i := 0; i < 25; i++ {
		rover := create.Random()
		assert.NotNil(t, rover)
	}
}
