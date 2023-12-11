package action

import (
	"errors"
	"fmt"
	. "github.com/google/uuid"
	"mars_rover/src/action/command"
	. "mars_rover/src/domain/rover"
	"strings"
)

func (this *LaxAction) MoveSequence(roverId UUID, commands command.Commands) ([]MovementResult, *ActionError) {
	rover, err := this.repo.GetRover(roverId)
	if err != nil {
		return nil, BuildNotFound(roverId, err)
	}

	movementResults := moveRover(rover, commands)

	err = this.repo.UpdateRover(rover)
	if err != nil {
		return nil, BuildNotUpdated(roverId, err)
	}

	return movementResults, nil
}

func moveRover(rover Rover, commands command.Commands) []MovementResult {
	commandToRoverFunctionMap := map[command.Command]interface{}{
		command.Forward:  Movement(rover.MoveForward),
		command.Backward: Movement(rover.MoveBackward),
		command.Left:     Rotation(rover.TurnLeft),
		command.Right:    Rotation(rover.TurnRight),
	}

	results := make([]MovementResult, 0, len(commands))
	var err error
	for _, cmd := range commands {
		action, doesMap := commandToRoverFunctionMap[cmd]
		if doesMap {
			switch action := action.(type) {
			case Movement:
				err = action()
			case Rotation:
				action()
				err = nil
			}
		}
		res := MovementResult{
			Cmd:           cmd,
			IssueDetected: err != nil,
			Coord:         rover.Coordinate(),
			Dir:           rover.Direction(),
		}
		results = append(results, res)
	}

	return results
}

func (this *LaxAction) MoveSequenceAborting(rover Rover, commands string) (Rover, error) {
	if rover == nil {
		return nil, errors.New("unexpected error, got nil rover")
	}
	for _, cmd := range strings.ToLower(commands) {
		err := mapCommandToMovement(rover, string(cmd))
		if err != nil {
			return rover, fmt.Errorf("aborting Command '%v': %v", string(cmd), err)
		}
	}
	err := this.repo.UpdateRover(rover)
	if err != nil {
		return nil, fmt.Errorf("unexpected error, couldn't save rover: %v", err)
	}
	return rover, nil
}

type (
	Movement func() error
	Rotation func()
)

// TODO: this should go somewhere in the API controller
func mapCommandToMovement(rover Rover, command string) error {
	commandActions := map[string]interface{}{
		"f": Movement(rover.MoveForward),
		"b": Movement(rover.MoveBackward),
		"l": Rotation(rover.TurnLeft),
		"r": Rotation(rover.TurnRight),
	}
	// if action := commandActions[Command]; action != nil {}
	// TODO.LM: not sure if ⬆️ is more readable than ⬇️ ¯\_(ツ)_/¯
	// Go people like ⬆️ but I think that ⬇️ is easier to read if you come from other langs
	if action, ok := commandActions[command]; ok {
		switch action := action.(type) {
		case Movement:
			return action()
		case Rotation:
			action()
			return nil
		}
	}
	// TODO: this error is not a collision, don't treat it as such
	return errors.New("invalid command")
}
