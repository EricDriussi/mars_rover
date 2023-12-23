package error

import (
	"fmt"
	. "github.com/google/uuid"
)

type ErrorType int

const (
	RoverNotFound ErrorType = iota
	RoverNotUpdated
)

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
