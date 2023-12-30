package resilient_mover

import (
	. "github.com/google/uuid"
	. "mars_rover/src/action/move"
	. "mars_rover/src/action/move/command"
	. "mars_rover/src/domain"
	. "mars_rover/src/domain/rover"
)

type ResilientMover struct {
	repo Repository
}

func With(repo Repository) *ResilientMover {
	return &ResilientMover{
		repo: repo,
	}
}

func (this *ResilientMover) Move(roverId UUID, commands Commands) ([]MovementResult, *MovementError) {
	rover, err := this.repo.GetRover(roverId)
	if err != nil {
		return nil, NotFoundErr()
	}

	movementResults := moveRover(rover, commands)

	err = this.repo.UpdateRover(rover)
	if err != nil {
		return nil, NotUpdatedErr()
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
		if err != nil {
			results = append(results, buildMovementFail(rover, cmd))
		} else {
			results = append(results, buildMovementSuccess(rover, cmd))
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

func buildMovementSuccess(rover Rover, cmd Command) MovementResult {
	return buildMovementResult(rover, cmd, false)
}

func buildMovementFail(rover Rover, cmd Command) MovementResult {
	return buildMovementResult(rover, cmd, true)
}

func buildMovementResult(rover Rover, cmd Command, hadIssue bool) MovementResult {
	return MovementResult{
		Cmd:           cmd,
		IssueDetected: hadIssue,
		Coord:         rover.Coordinate(),
		Dir:           rover.Direction(),
	}
}
