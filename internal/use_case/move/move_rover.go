package move

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	. "mars_rover/internal/domain"
	. "mars_rover/internal/domain/rover"
	"strings"
)

type UseCase struct {
	repo Repository
}

func For(repo Repository) *UseCase {
	return &UseCase{
		repo: repo,
	}
}

func (this *UseCase) MoveSequence(id string, commands string) []error {
	var errs []error
	uid, err := uuid.Parse(id)
	if err != nil {
		errs = append(errs, errors.New("invalid id format"))
		return errs
	}
	rover, err := this.repo.GetRover(uid)
	if err != nil {
		errs = append(errs, errors.New(fmt.Sprintf("Repository error: %v", err)))
		return errs
	}
	for _, cmd := range strings.ToLower(commands) {
		err := mapCommandToMovement(rover, string(cmd))
		if err != nil {
			errs = append(errs, errors.New(fmt.Sprintf("%v, skipping command %v", err, string(cmd))))
		}
	}
	err = this.repo.UpdateRover(rover)
	if err != nil {
		errs = append(errs, err)
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func (this *UseCase) MoveSequenceAborting(id string, commands string) error {
	uid, err := uuid.Parse(id)
	if err != nil {
		return errors.New("invalid id format")
	}
	rover, err := this.repo.GetRover(uid)
	if err != nil {
		return errors.New(fmt.Sprintf("Repository error: %v", err))
	}
	for _, cmd := range strings.ToLower(commands) {
		err := mapCommandToMovement(rover, string(cmd))
		if err != nil {
			return errors.New(fmt.Sprintf("aborting command '%v': %v", string(cmd), err))
		}
	}
	return this.repo.UpdateRover(rover)
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
