package move

import (
	"errors"
	"fmt"
	. "mars_rover/internal/domain"
	. "mars_rover/internal/domain/rover"
	"strings"
)

type Action struct {
	repo Repository
}

func For(repo Repository) *Action {
	return &Action{
		repo: repo,
	}
}

func (this *Action) MoveSequence(rover Rover, commands string) (Rover, []error) {
	if rover == nil {
		return nil, []error{errors.New("unexpected error, got nil rover")}
	}
	var movementErrors []error
	for _, cmd := range strings.ToLower(commands) {
		err := mapCommandToMovement(rover, string(cmd))
		if err != nil {
			movementErrors = append(movementErrors, errors.New(fmt.Sprintf("%v, skipping command %v", err, string(cmd))))
		}
	}
	err := this.repo.UpdateRover(rover)
	if err != nil {
		return nil, []error{errors.New(fmt.Sprintf("unexpected error, couldn't save rover: %v", err))}
	}
	return rover, movementErrors
}

func (this *Action) MoveSequenceAborting(rover Rover, commands string) (Rover, error) {
	if rover == nil {
		return nil, errors.New("unexpected error, got nil rover")
	}
	for _, cmd := range strings.ToLower(commands) {
		err := mapCommandToMovement(rover, string(cmd))
		if err != nil {
			return rover, errors.New(fmt.Sprintf("aborting command '%v': %v", string(cmd), err))
		}
	}
	err := this.repo.UpdateRover(rover)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("unexpected error, couldn't save rover: %v", err))
	}
	return rover, nil
}

type (
	Movement func() error
	Rotation func()
)

func mapCommandToMovement(rover Rover, command string) error {
	commandActions := map[string]interface{}{
		"f": Movement(rover.MoveForward),
		"b": Movement(rover.MoveBackward),
		"l": Rotation(rover.TurnLeft),
		"r": Rotation(rover.TurnRight),
	}
	// if action := commandActions[command]; action != nil {}
	// TODO.LM: is ⬆️ more readable than ⬇️ ?
	if action, ok := commandActions[command]; ok {
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
