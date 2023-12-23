package strict_mover

import (
	. "github.com/google/uuid"
	. "mars_rover/src/action"
	. "mars_rover/src/action/move"
	"mars_rover/src/action/move/command"
	. "mars_rover/src/domain"
	. "mars_rover/src/domain/rover"
)

type StrictMover struct {
	repo Repository
}

func For(repo Repository) *StrictMover {
	return &StrictMover{
		repo: repo,
	}
}

func (this *StrictMover) Move(roverId UUID, commands command.Commands) ([]MovementResult, *MovementError) {
	rover, err := this.repo.GetRover(roverId)
	if err != nil {
		return nil, BuildNotFoundErr(roverId, err)
	}

	movementResults := moveRover(rover, commands)

	err = this.repo.UpdateRover(rover)
	if err != nil {
		return nil, BuildNotUpdatedErr(roverId, err)
	}

	return movementResults, nil
}

type (
	Movement func() error
	Rotation func()
)

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
				break
			case Rotation:
				action()
				break
			}
		}
		movementIssue := err != nil
		res := MovementResult{
			Cmd:           cmd,
			IssueDetected: movementIssue,
			Coord:         rover.Coordinate(),
			Dir:           rover.Direction(),
		}
		results = append(results, res)
		if movementIssue {
			break
		}
	}

	return results
}
