package controllers

import (
	"fmt"
	"github.com/google/uuid"
	. "mars_rover/src/action"
	"mars_rover/src/action/command"
	. "mars_rover/src/infra/apiServer/dto"
)

type MoveRequest struct {
	Commands string `json:"commands"`
	Id       string `json:"id"`
}

func MoveRover(action Action, request MoveRequest) (MovementResponseDTO, error) {
	roverId, err := uuid.Parse(request.Id)
	// TODO.LM: this validation would not happen if the uuid was wrapped
	// It's already being validated in the calling controller
	if err != nil {
		return MovementResponseDTO{}, err
	}
	// TODO: test that if len(applicationCommands) != len(request.Commands) some commands where not valid and skipped
	applicationCommands := command.FromString(request.Commands)
	movementResults, err := action.MoveSequence(roverId, applicationCommands)
	// TODO: this would send a 500 if the rover was not found, but it should be a 404
	if err != nil {
		return MovementResponseDTO{}, err
	}

	out := mapResultsToDTO(movementResults)

	return out, nil
}

func mapResultsToDTO(results []MovementResult) MovementResponseDTO {
	var out MovementResponseDTO
	for _, result := range results {
		issue := ""
		if result.IssueDetected {
			issue = fmt.Sprintf("unable to move on command '%v'.", result.Cmd.ToString())
		}
		out.Results = append(out.Results, SingleMovementDTO{
			Issue: issue,
			Coordinate: CoordinateDTO{
				X: result.Coord.X(),
				Y: result.Coord.Y(),
			},
			Direction: result.Dir.CardinalPoint(),
		})
	}
	return out
}
