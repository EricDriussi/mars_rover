package infra

import (
	. "github.com/google/uuid"
	. "github.com/stretchr/testify/mock"
	. "mars_rover/internal/domain"
	. "mars_rover/internal/domain/planet"
	. "mars_rover/internal/domain/rover"
)

type MockRepo struct {
	Mock
}

func (this MockRepo) UpdateRover(rover Rover) error {
	args := this.Called()
	return args.Error(0)
}

func (this MockRepo) SaveGame(rover Rover, planet Planet) error {
	args := this.Called()
	return args.Error(0)
}

func (this MockRepo) LoadGame(id UUID) (GameDTO, error) {
	args := this.Called()
	return GameDTO{}, args.Error(0)
}

func (this MockRepo) GetRover(id UUID) (Rover, error) {
	args := this.Called()
	return args.Get(0).(Rover), args.Error(0)
}
