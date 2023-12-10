package action_test

import (
	"github.com/stretchr/testify/assert"
	"mars_rover/src/action"
	. "mars_rover/src/action"
	"testing"
)

func TestFrom(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected Commands
	}{
		{
			name:     "parses valid commands",
			input:    "fblr",
			expected: Commands{Forward, Backward, Left, Right},
		},
		{
			name:     "skips invalid commands",
			input:    "fxb",
			expected: Commands{Forward, Backward},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expected, action.ParseFrom(testCase.input))
		})
	}
}
