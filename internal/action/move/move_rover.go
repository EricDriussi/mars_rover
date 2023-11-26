package move

import (
	"errors"
	"fmt"
	. "mars_rover/internal/domain"
	. "mars_rover/internal/domain/rover"
	"reflect"
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
	roverBeforeMovement, err := copyOfRover(rover)
	if err != nil {
		return nil, []error{errors.New("unexpected error, got nil rover")}
	}
	var movementErrors []error
	for _, cmd := range strings.ToLower(commands) {
		err := mapCommandToMovement(rover, string(cmd))
		if err != nil {
			movementErrors = append(movementErrors, errors.New(fmt.Sprintf("%v, skipping command %v", err, string(cmd))))
		}
	}
	err = this.repo.UpdateRover(rover)
	if err != nil {
		return roverBeforeMovement, []error{errors.New(fmt.Sprintf("unexpected error, couldn't save rover: %v", err))}
	}
	return rover, movementErrors
}

func (this *Action) MoveSequenceAborting(rover Rover, commands string) (Rover, error) {
	roverBeforeMovement, err := copyOfRover(rover)
	if err != nil {
		return nil, errors.New("unexpected error, got nil rover")
	}
	for _, cmd := range strings.ToLower(commands) {
		err := mapCommandToMovement(rover, string(cmd))
		if err != nil {
			return rover, errors.New(fmt.Sprintf("aborting command '%v': %v", string(cmd), err))
		}
	}
	err = this.repo.UpdateRover(rover)
	if err != nil {
		return roverBeforeMovement, errors.New(fmt.Sprintf("unexpected error, couldn't save rover: %v", err))
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

func copyOfRover(original Rover) (Rover, error) {
	if original == nil {
		return nil, errors.New("cannot create copy of nil rover")
	}

	roverType := reflect.TypeOf(original).Elem()
	copyRover := reflect.New(roverType).Interface().(Rover)

	copyRoverValue := reflect.ValueOf(copyRover).Elem()
	originalRoverValue := reflect.ValueOf(original).Elem()

	for i := 0; i < roverType.NumField(); i++ {
		copyRoverValue.Field(i).Set(originalRoverValue.Field(i))
	}

	return copyRover, nil
}
