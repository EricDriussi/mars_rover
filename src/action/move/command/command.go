package command

import (
	. "mars_rover/src/domain/rover"
	. "strings"
)

type Commands []Command
type Command int
type RoverMovementFunc func() error

const (
	Forward Command = iota
	Backward
	Left
	Right
)

func (this Command) String() string {
	var commandToCharMap = map[Command]rune{
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
	var charToCommandMap = map[rune]Command{
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

func (this Command) MapToRoverMovementFunction(rover Rover) RoverMovementFunc {
	return map[Command]interface{}{
		Forward:  RoverMovementFunc(rover.MoveForward),
		Backward: RoverMovementFunc(rover.MoveBackward),
		Left: RoverMovementFunc(func() error {
			rover.TurnLeft()
			return nil
		}),
		Right: RoverMovementFunc(func() error {
			rover.TurnRight()
			return nil
		}),
	}[this].(RoverMovementFunc)
}
