package id

import (
	"github.com/google/uuid"
)

// TODO.LM: just a wrapper for Google's UUID
type ID struct {
	uuid.UUID
}

func (this ID) String() string {
	return this.UUID.String()
}

func New() ID {
	return ID{uuid.New()}
}

func Parse(id string) (ID, error) {
	parsedId, err := uuid.Parse(id)
	if err != nil {
		return ID{}, err
	}
	return ID{parsedId}, nil
}
