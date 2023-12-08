package controllers

import (
	"encoding/json"
	"github.com/google/uuid"
	. "mars_rover/src/action/move"
	. "mars_rover/src/infra/apiServer/dto"
)

type MoveRequest struct {
	Commands string `json:"commands"`
	Id       string `json:"id"`
}

func MoveRover(action Action, request MoveRequest) ([]byte, error) {
	roverId, err := uuid.Parse(request.Id)
	if err != nil {
		return nil, err
	}
	movementResult := action.MoveSequence(roverId, request.Commands)
	if movementResult.Error != nil {
		return nil, err
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

	response := MovementResponseDTO{
		// TODO: returning the rover is not enough, should return a list of coordinates-directions since one command might fail but the rover can keep moving
		// Decide in front end if paint all positions or just the last one
		Rover: roverToReturn,
		// TODO: these are not "Errors", they are collisions
		Errors: movementResult.MovementErrors.AsStringArray(),
		// TODO: what about non-movement errors?
	}

	return json.Marshal(response)
}
