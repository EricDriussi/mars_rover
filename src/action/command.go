package action

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

var charToCommandMap = map[rune]Command{
	'f': Forward,
	'b': Backward,
	'l': Left,
	'r': Right,
}

func from(char rune) (Command, error) {
	command, doesMap := charToCommandMap[char]
	if !doesMap {
		return invalidCommand, fmt.Errorf("invalid Command: %c", char)
	}
	return command, nil
}

func ParseFrom(str string) Commands {
	commands := make([]Command, 0, len(str))
	for _, char := range strings.ToLower(str) {
		command, err := from(char)
		if err != nil {
			continue
		}
		commands = append(commands, command)
	}
	return commands
}

var commandToCharMap = map[Command]rune{
	Forward:  'f',
	Backward: 'b',
	Left:     'l',
	Right:    'r',
}

func (c Command) toString() string {
	char, doesMap := commandToCharMap[c]
	if !doesMap {
		return "invalid"
	}
	return string(char)
}
