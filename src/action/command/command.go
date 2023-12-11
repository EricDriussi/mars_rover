package command

import (
	"fmt"
	"strings"
)

type Commands []Command
type Command int

const invalidCommand Command = -1
const (
	Forward Command = iota
	Backward
	Left
	Right
)

func (this Command) ToString() string {
	var commandToCharMap = map[Command]rune{
		Forward:  'f',
		Backward: 'b',
		Left:     'l',
		Right:    'r',
	}

	char, doesMap := commandToCharMap[this]
	if !doesMap {
		return "cannot stringify invalid Command"
	}
	return string(char)
}

func FromString(stringCommands string) Commands {
	commands := make([]Command, 0, len(stringCommands))
	for _, char := range strings.ToLower(stringCommands) {
		command, err := commandFrom(char)
		if err != nil {
			continue
		}
		commands = append(commands, command)
	}
	return commands
}

func commandFrom(char rune) (Command, error) {
	var charToCommandMap = map[rune]Command{
		'f': Forward,
		'b': Backward,
		'l': Left,
		'r': Right,
	}

	command, doesMap := charToCommandMap[char]
	if !doesMap {
		return invalidCommand, fmt.Errorf("invalid Command: %c", char)
	}
	return command, nil
}
