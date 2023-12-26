package move

import (
	"fmt"
	. "github.com/google/uuid"
)

type errorType int

const (
	roverNotFound errorType = iota
	roverNotUpdated
)

type MovementError struct {
	ID      UUID
	errType errorType
}

func (e MovementError) Error() string {
	if e.errType == roverNotFound {
		return fmt.Sprintf("rover with ID %s not found", e.ID)
	}
	if e.errType == roverNotUpdated {
		return fmt.Sprintf("failed to update rover with ID %s", e.ID)
	}
	return "unknown movement error"
}

func (e MovementError) IsNotFound() bool {
	return e.errType == roverNotFound
}

func (e MovementError) IsNotUpdated() bool {
	return e.errType == roverNotUpdated
}

func BuildNotFoundErr() *MovementError {
	return &MovementError{
		errType: roverNotFound,
	}
}

func BuildNotUpdatedErr() *MovementError {
	return &MovementError{
		errType: roverNotUpdated,
	}
}
