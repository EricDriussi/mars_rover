package domain

import (
	"fmt"
)

type errorType int

const (
	persistenceMalfunction errorType = iota
	alreadyExists
	couldNotUpdate
	couldNotMap
	couldNotAdd
	notFound
)

type RepositoryError struct {
	errType errorType
	errMsg  string
}

func (e RepositoryError) Error() string {
	if e.errType == persistenceMalfunction {
		return fmt.Sprintf("persistence malfunction: %s", e.errMsg)
	}
	if e.errType == alreadyExists {
		return e.errMsg
	}
	if e.errType == couldNotUpdate {
		return fmt.Sprintf("could not update resource: %s", e.errMsg)
	}
	if e.errType == notFound {
		return e.errMsg
	}
	if e.errType == couldNotMap {
		return fmt.Sprintf("could not map resource: %s", e.errMsg)
	}
	if e.errType == couldNotAdd {
		return fmt.Sprintf("could not add resource: %s", e.errMsg)
	}
	return "unknown error"
}

func (e RepositoryError) IsAlreadyExists() bool {
	return e.errType == alreadyExists
}

func (e RepositoryError) IsCouldNotUpdate() bool {
	return e.errType == couldNotUpdate
}

func PersistenceMalfunction(err error) *RepositoryError {
	return &RepositoryError{
		errType: persistenceMalfunction,
		errMsg:  err.Error(),
	}
}

func AlreadyExists() *RepositoryError {
	return &RepositoryError{
		errType: alreadyExists,
		errMsg:  "resource already exists",
	}
}

func CouldNotUpdate(err error) *RepositoryError {
	return &RepositoryError{
		errType: persistenceMalfunction,
		errMsg:  err.Error(),
	}
}

func NotFound() *RepositoryError {
	return &RepositoryError{
		errType: notFound,
		errMsg:  "resource not found",
	}
}

func CouldNotMap(err error) *RepositoryError {
	return &RepositoryError{
		errType: couldNotMap,
		errMsg:  err.Error(),
	}
}

func CouldNotAdd(err error) *RepositoryError {
	return &RepositoryError{
		errType: couldNotAdd,
		errMsg:  err.Error(),
	}
}
