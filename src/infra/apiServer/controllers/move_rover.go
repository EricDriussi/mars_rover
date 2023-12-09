package controllers

import (
	"github.com/google/uuid"
	. "mars_rover/src/action"
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
	movementResult := action.MoveSequence(roverId, request.Commands)
	if movementResult.Error != nil {
		return MovementResponseDTO{}, movementResult.Error
	}

	updatedRover := movementResult.Rover
	coordinate := updatedRover.Coordinate()
	roverToReturn := RoverDTO{
		Id: updatedRover.Id().String(),
		Coordinate: CoordinateDTO{
			X: coordinate.X(),
			Y: coordinate.Y(),
		},
		Direction: updatedRover.Direction().CardinalPoint(),
	}

	return MovementResponseDTO{
		// TODO: returning the rover is not enough, should return a list of coordinates-directions since one command might fail but the rover can keep moving
		// Decide in front end if paint all positions or just the last one
		Rover: roverToReturn,
		// TODO: these are not "Errors", they are collisions
		Errors: movementResult.MovementErrors.AsStringArray(),
		// TODO: what about non-movement errors?
	}, nil
}
