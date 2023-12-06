package infra

import (
	. "github.com/google/uuid"
	. "github.com/stretchr/testify/mock"
	. "mars_rover/src/domain/planet"
	. "mars_rover/src/domain/rover"
)

type MockRepo struct {
	Mock
}

func (this MockRepo) UpdateRover(rover Rover) error {
	args := this.Called()
	return args.Error(0)
}

func (this MockRepo) AddRover(rover Rover, planetId int64) error {
	args := this.Called()
	return args.Error(0)
}

func (this MockRepo) AddPlanet(planet Planet) (int64, error) {
	args := this.Called()
	return 0, args.Error(0)
}

func (this MockRepo) GetRover(roverId UUID) (Rover, error) {
	args := this.Called()
	return nil, args.Error(0)
}

func (this MockRepo) GetPlanet(roverId UUID) (Planet, error) {
	args := this.Called()
	return nil, args.Error(0)
}
