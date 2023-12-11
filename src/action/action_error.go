package action

import (
	"fmt"
	. "github.com/google/uuid"
)

type ErrorType int

const (
	RoverNotFound ErrorType = iota
	RoverNotUpdated
)

func (e ActionError) Type() ErrorType {
	if e.RoverNotFoundError != nil {
		return RoverNotFound
	}
	if e.UpdateRoverError != nil {
		return RoverNotUpdated
	}
	return -1
}

type RoverNotFoundError struct {
	ID  UUID
	Err error
}

func (e RoverNotFoundError) Error() string {
	return fmt.Sprintf("rover with ID %s not found", e.ID)
}

type RoverNotUpdatedError struct {
	ID  UUID
	Err error
}

func (e RoverNotUpdatedError) Error() string {
	return fmt.Sprintf("failed to update rover with ID %s", e.ID)
}

type ActionError struct {
	RoverNotFoundError *RoverNotFoundError
	UpdateRoverError   *RoverNotUpdatedError
}

func (e ActionError) Error() string {
	if e.RoverNotFoundError != nil {
		return e.RoverNotFoundError.Error()
	}
	if e.UpdateRoverError != nil {
		return e.UpdateRoverError.Error()
	}
	return "unknown error"
}

func BuildNotFound(id UUID, err error) *ActionError {
	return &ActionError{
		RoverNotFoundError: &RoverNotFoundError{
			ID:  id,
			Err: err,
		},
	}
}

func BuildNotUpdated(id UUID, err error) *ActionError {
	return &ActionError{
		UpdateRoverError: &RoverNotUpdatedError{
			ID:  id,
			Err: err,
		},
	}
}
