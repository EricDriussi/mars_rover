package dto

import (
	"fmt"
	. "mars_rover/src/action"
)

func FromActionResult(movementResult []MovementResult) MovementResponseDTO {
	var responseDTO MovementResponseDTO
	for _, result := range movementResult {
		issue := ""
		if result.IssueDetected {
			issue = fmt.Sprintf("unable to move on command '%v'.", result.Cmd.ToString())
		}
		responseDTO.Results = append(responseDTO.Results, SingleMovementDTO{
			Issue: issue,
			Coordinate: CoordinateDTO{
				X: result.Coord.X(),
				Y: result.Coord.Y(),
			},
			Direction: result.Dir.CardinalPoint(),
		})
	}
	return responseDTO
}
