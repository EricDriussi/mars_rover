package command

import (
	. "mars_rover/src/domain/rover"
	. "strings"
)

type RoverMovementFunc func() error
type Command interface {
	MapToRoverMovementFunction(rover Rover) RoverMovementFunc
	String() string
}
type Commands []Command
type BasicCommand int

const (
	Forward BasicCommand = iota
	Backward
	Left
	Right
)

func (this BasicCommand) String() string {
	var commandToCharMap = map[BasicCommand]rune{
		Forward:  'f',
		Backward: 'b',
		Left:     'l',
		Right:    'r',
	}

	char, ok := commandToCharMap[this]
	if !ok {
		return "?"
	}
	return string(char)
}

func FromString(stringCommands string) Commands {
	commands := make(Commands, 0, len(stringCommands))
	for _, char := range ToLower(stringCommands) {
		commands = appendIfValid(char, commands)
	}
	return commands
}

func appendIfValid(char rune, commands Commands) Commands {
	var charToCommandMap = map[rune]BasicCommand{
		'f': Forward,
		'b': Backward,
		'l': Left,
		'r': Right,
	}

	command, ok := charToCommandMap[char]
	if ok {
		commands = append(commands, command)
	}
	return commands
}

func (this BasicCommand) MapToRoverMovementFunction(rover Rover) RoverMovementFunc {
	return map[BasicCommand]interface{}{
		Forward: RoverMovementFunc(func() error {
			return rover.MoveForward()
		}),
		Backward: RoverMovementFunc(func() error {
			return rover.MoveBackward()
		}),
		Left: RoverMovementFunc(func() error {
			rover.TurnLeft()
			return nil // TODO.LM: rotations can't error, this is just done to simplify both types of movement into a single Go type (RoverMovementFunc)
		}),
		Right: RoverMovementFunc(func() error {
			rover.TurnRight()
			return nil // TODO.LM: rotations can't error, this is just done to simplify both types of movement into a single Go type (RoverMovementFunc)
		}),
	}[this].(RoverMovementFunc)
}
