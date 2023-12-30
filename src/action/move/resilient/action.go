package resilient_mover

import (
	. "mars_rover/src/action/move"
	. "mars_rover/src/action/move/command"
	. "mars_rover/src/domain"
	. "mars_rover/src/domain/rover"
	. "mars_rover/src/domain/rover/uuid"
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

func moveRover(rover Rover, commands Commands) []MovementResult {
	results := make([]MovementResult, 0)
	for _, cmd := range commands {
		result := execute(rover, cmd)
		results = append(results, result)
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
