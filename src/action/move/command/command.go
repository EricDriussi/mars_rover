package command

import (
	. "strings"
)

type Commands []Command
type Command int

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
