package move

import (
	. "github.com/google/uuid"
	. "mars_rover/src/action/error"
)

type MovementError struct {
	RoverNotFoundError *RoverNotFoundError
	UpdateRoverError   *RoverNotUpdatedError
}

func (e MovementError) Error() string {
	if e.RoverNotFoundError != nil {
		return e.RoverNotFoundError.Error()
	}
	if e.UpdateRoverError != nil {
		return e.UpdateRoverError.Error()
	}
	return "unknown error"
}

func (e MovementError) Type() ErrorType {
	if e.RoverNotFoundError != nil {
		return RoverNotFound
	}
	if e.UpdateRoverError != nil {
		return RoverNotUpdated
	}
	return -1
}

func BuildNotFoundErr(id UUID, err error) *MovementError {
	return &MovementError{
		RoverNotFoundError: &RoverNotFoundError{
			ID:  id,
			Err: err,
		},
	}
}

func BuildNotUpdatedErr(id UUID, err error) *MovementError {
	return &MovementError{
		UpdateRoverError: &RoverNotUpdatedError{
			ID:  id,
			Err: err,
		},
	}
}
