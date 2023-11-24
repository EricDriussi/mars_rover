package infra

import (
	. "github.com/google/uuid"
	. "mars_rover/internal/domain/rover"
	. "mars_rover/internal/infra/mappers"
)

func (r *SQLiteRepository) GetRover(id UUID) (Rover, error) {
	roverEntity, err := r.getRover(id)
	if err != nil {
		return nil, err
	}
	if !roverEntity.Present {
		return nil, nil
	}

	domainRover, err := MapToDomainRover(roverEntity.Value)
	if err != nil {
		return nil, err
	}

	return domainRover, nil
}
