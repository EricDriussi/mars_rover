package strict_mover

import (
	. "mars_rover/src/action/mover"
	. "mars_rover/src/action/mover/command"
	. "mars_rover/src/domain"
	. "mars_rover/src/domain/rover"
	. "mars_rover/src/domain/rover/id"
)

type StrictMover struct {
	repo Repository
}

func With(repo Repository) *StrictMover {
	return &StrictMover{
		repo: repo,
	}
}

func (this *StrictMover) Move(roverId ID, commands Commands) ([]MovementResult, *MovementError) {
	rover, getErr := this.repo.GetRover(roverId)
	if getErr != nil {
		return nil, NotFoundErr()
	}

	movementResults := moveRover(rover, commands)

	updateErr := this.repo.UpdateRover(rover)
	if updateErr != nil {
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
