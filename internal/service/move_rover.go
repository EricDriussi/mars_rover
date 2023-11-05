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
	// TODO: should this skip invalid commands? stop if an invalid command is found?
	for _, cmd := range commands {
		this.mapCommandToMovement(cmd)
	}
}

type (
	Movement func() error
	Rotation func()
)

func (this *MoveService) mapCommandToMovement(command string) error {
	commandActions := map[string]interface{}{
		"f": Movement(this.rover.MoveForward),
		"b": Movement(this.rover.MoveBackward),
		"l": Rotation(this.rover.TurnLeft),
		"r": Rotation(this.rover.TurnRight),
	}
	if action, ok := commandActions[command]; ok {
		// TODO.LM: is this more readable than ⬇️ ?
		// if action := commandActions[command]; action != nil {}
		switch action := action.(type) {
		case Movement:
			err := action()
			if err != nil {
				return err
			}
		case Rotation:
			action()
			return nil
		}
	}
	return errors.New(fmt.Sprintf("invalid command, don't know what to do with %v", command))
}
