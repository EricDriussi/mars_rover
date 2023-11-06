package service

import (
	"errors"
	"fmt"
	"mars_rover/internal/domain/rover"
	"strings"
)

type MoveService struct {
	rover rover.Rover
}

func For(rover rover.Rover) *MoveService {
	return &MoveService{
		rover: rover,
	}
}

func (this *MoveService) MoveSequence(commands string) []error {
	var errs []error
	for _, cmd := range strings.ToLower(commands) {
		err := this.mapCommandToMovement(string(cmd))
		// TODO: Persist rover state (location? rover?)
		if err != nil {
			errs = append(errs, errors.New(fmt.Sprintf("%v, skipping command %v", err, cmd)))
		}
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
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
			return action()
		case Rotation:
			action()
			return nil
		}
	}
	return errors.New("invalid command")
}
