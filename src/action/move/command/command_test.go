package command_test

import (
	"github.com/stretchr/testify/assert"
	. "mars_rover/src/action/move/command"
	"testing"
)

func TestCommandsAreBuiltFromString(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected Commands
	}{
		{
			name:     "containing valid commands",
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
			assert.Equal(t, testCase.expected, FromString(testCase.input))
		})
	}
}

func TestCommandsArePrintedAsStrings(t *testing.T) {
	testCases := []struct {
		name     string
		command  Command
		expected string
	}{
		{
			name:     "forward",
			command:  Forward,
			expected: "f",
		},
		{
			name:     "backward",
			command:  Backward,
			expected: "b",
		},
		{
			name:     "left",
			command:  Left,
			expected: "l",
		},
		{
			name:     "right",
			command:  Right,
			expected: "r",
		},
		{
			name:     "unknown",
			command:  Command(999),
			expected: "?",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expected, testCase.command.String())
		})
	}
}
