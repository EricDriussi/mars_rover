package strict_mover

import (
	. "github.com/google/uuid"
	. "mars_rover/src/action/move"
	. "mars_rover/src/action/move/command"
	. "mars_rover/src/domain"
	. "mars_rover/src/domain/rover"
)

type StrictMover struct {
	repo Repository
}

func With(repo Repository) *StrictMover {
	return &StrictMover{
		repo: repo,
	}
}

func (this *StrictMover) Move(roverId UUID, commands Commands) ([]MovementResult, *MovementError) {
	rover, err := this.repo.GetRover(roverId)
	if err != nil {
		return nil, BuildNotFoundErr()
	}

	movementResults := moveRover(rover, commands)

	err = this.repo.UpdateRover(rover)
	if err != nil {
		return nil, BuildNotUpdatedErr()
	}

	return movementResults, nil
}

type (
	Movement func() error
	Rotation func()
)

func moveRover(rover Rover, commands Commands) []MovementResult {
	commandToRoverFunctionMap := map[Command]interface{}{
		Forward:  Movement(rover.MoveForward),
		Backward: Movement(rover.MoveBackward),
		Left:     Rotation(rover.TurnLeft),
		Right:    Rotation(rover.TurnRight),
	}

	results := make([]MovementResult, 0, len(commands))
	for _, cmd := range commands {
		roverAction, ok := commandToRoverFunctionMap[cmd]
		if !ok {
			continue
		}
		err := execute(roverAction)
		movementIssue := err != nil
		results = append(results, buildMovementResult(rover, cmd, movementIssue))
		if movementIssue {
			break
		}
	}

	return results
}

func execute(roverAction interface{}) error {
	switch roverAction := roverAction.(type) {
	case Movement:
		return roverAction()
	case Rotation:
		roverAction()
		return nil
	}
	return nil
}

func buildMovementResult(rover Rover, cmd Command, movementIssue bool) MovementResult {
	return MovementResult{
		Cmd:           cmd,
		IssueDetected: movementIssue,
		Coord:         rover.Coordinate(),
		Dir:           rover.Direction(),
	}
}
