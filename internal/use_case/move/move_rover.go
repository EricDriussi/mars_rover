package move

import (
	"errors"
	"fmt"
	. "mars_rover/internal/domain"
	. "mars_rover/internal/domain/rover"
	"strings"
)

type UseCase struct {
	rover Rover
	repo  Repository
}

func For(rover Rover, repo Repository) *UseCase {
	return &UseCase{
		rover: rover,
		repo:  repo,
	}
}

func (this *UseCase) MoveSequence(commands string) []error {
	var errs []error
	for _, cmd := range strings.ToLower(commands) {
		err := this.mapCommandToMovement(string(cmd))
		if err != nil {
			errs = append(errs, errors.New(fmt.Sprintf("%v, skipping command %v", err, cmd)))
		}
	}
	err := this.repo.UpdateRover(this.rover)
	if err != nil {
		errs = append(errs, err)
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func (this *UseCase) MoveSequenceAborting(commands string) error {
	for _, cmd := range strings.ToLower(commands) {
		err := this.mapCommandToMovement(string(cmd))
		if err != nil {
			return errors.New(fmt.Sprintf("aborting command '%v': %v", cmd, err))
		}
	}
	return this.repo.UpdateRover(this.rover)
}

type (
	Movement func() error
	Rotation func()
)

func (this *UseCase) mapCommandToMovement(command string) error {
	commandActions := map[string]interface{}{
		"f": Movement(this.rover.MoveForward),
		"b": Movement(this.rover.MoveBackward),
		"l": Rotation(this.rover.TurnLeft),
		"r": Rotation(this.rover.TurnRight),
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
