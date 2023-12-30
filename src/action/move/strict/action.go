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
		return nil, NotFoundErr()
	}

	movementResults := moveRover(rover, commands)

	err = this.repo.UpdateRover(rover)
	if err != nil {
		return nil, NotUpdatedErr()
	}

	return movementResults, nil
}

func moveRover(rover Rover, commands Commands) []MovementResult {
	results := make([]MovementResult, 0)
	for _, cmd := range commands {
		result := execute(rover, cmd)
		results = append(results, result)
		if result.IssueDetected {
			break
		}
	}
	return results
}

func execute(rover Rover, cmd Command) MovementResult {
	roverFunction := cmd.MapToRoverMovementFunction(rover)
	err := roverFunction()
	if err != nil {
		return BuildFailedMovementResult(rover, cmd)
	} else {
		return BuildSuccessfulMovementResult(rover, cmd)
	}
}
