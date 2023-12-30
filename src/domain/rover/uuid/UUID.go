package uuid

import (
	"github.com/google/uuid"
)

type UUID struct {
	uuid.UUID
}

func (this UUID) String() string {
	return this.UUID.String()
}

func New() UUID {
	return UUID{uuid.New()}
}

func Parse(id string) (UUID, error) {
	parsedId, err := uuid.Parse(id)
	if err != nil {
		return UUID{}, err
	}
	return UUID{parsedId}, nil
}
