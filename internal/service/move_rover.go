package service

import (
	"errors"
	"fmt"
	"mars_rover/internal/domain/rover"
)

type MoveService struct {
	rover *rover.Rover
}

func For(rover *rover.Rover) *MoveService {
	return &MoveService{
		rover: rover,
	}
}

func (this *MoveService) MoveSequence(commands []string) {
	for _, cmd := range commands {
		this.mapCommandToMovement(cmd)
	}
}

func (this *MoveService) mapCommandToMovement(command string) error {
	commandActions := map[string]func(){
		"f": this.rover.MoveForward,
		"b": this.rover.MoveBackward,
		"l": this.rover.TurnLeft,
		"r": this.rover.TurnRight,
	}

	// TODO.LM: is this more readable than ⬇️ ?
	// if action := commandMapper[cmd]; action != nil {
	if action, ok := commandActions[command]; ok {
		action()
		return nil
	}
	return errors.New(fmt.Sprintf("invalid command, don't know what to do with %v", command))
}
