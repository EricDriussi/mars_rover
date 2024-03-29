package command_test

import (
	"github.com/stretchr/testify/assert"
	. "mars_rover/src/action/mover/command"
	"mars_rover/src/test_helpers/mocks"
	. "mars_rover/src/test_helpers/mocks"
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
		command  BasicCommand
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
			command:  BasicCommand(999),
			expected: "?",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expected, testCase.command.String())
		})
	}
}

func TestCommandsAreMappedToRoverMovementFunction(t *testing.T) {
	rover := new(MockRover)
	mocks.MakeAlwaysSuccessful(rover)
	testCases := []struct {
		name                  string
		roverFunction         RoverMovementFunc
		expectedRoverFunction string
	}{
		{
			name:                  "forward",
			roverFunction:         Forward.MapToRoverMovementFunction(rover),
			expectedRoverFunction: "MoveForward",
		},
		{
			name:                  "backward",
			roverFunction:         Backward.MapToRoverMovementFunction(rover),
			expectedRoverFunction: "MoveBackward",
		},
		{
			name:                  "left",
			roverFunction:         Left.MapToRoverMovementFunction(rover),
			expectedRoverFunction: "TurnLeft",
		},
		{
			name:                  "right",
			roverFunction:         Right.MapToRoverMovementFunction(rover),
			expectedRoverFunction: "TurnRight",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			err := testCase.roverFunction()
			assert.Nil(t, err)

			rover.AssertCalled(t, testCase.expectedRoverFunction)
		})
	}
}
