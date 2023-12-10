package action

import (
	"fmt"
	"strings"
)

type Commands []command
type command int

const invalidCommand command = -1
const (
	Forward command = iota
	Backward
	Left
	Right
)

var charToCommandMap = map[rune]command{
	'f': Forward,
	'b': Backward,
	'l': Left,
	'r': Right,
}

func from(char rune) (command, error) {
	command, doesMap := charToCommandMap[char]
	if !doesMap {
		return invalidCommand, fmt.Errorf("invalid command: %c", char)
	}
	return command, nil
}

func ParseFrom(str string) Commands {
	commands := make([]command, 0, len(str))
	for _, char := range strings.ToLower(str) {
		command, err := from(char)
		if err != nil {
			continue
		}
		commands = append(commands, command)
	}
	return commands
}
