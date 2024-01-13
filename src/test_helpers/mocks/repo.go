package mocks

import (
	. "github.com/stretchr/testify/mock"
	. "mars_rover/src/domain"
	. "mars_rover/src/domain/planet"
	. "mars_rover/src/domain/rover"
	. "mars_rover/src/domain/rover/id"
)

type MockRepo struct {
	Mock
}

func (this *MockRepo) UpdateRover(Rover) *RepositoryError {
	args := this.Called()
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(*RepositoryError)
}

func (this *MockRepo) AddRover(Rover, int) *RepositoryError {
	args := this.Called()
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(*RepositoryError)
}

func (this *MockRepo) AddPlanet(Planet) (int, *RepositoryError) {
	args := this.Called()
	if args.Get(1) == nil {
		return args.Get(0).(int), nil
	}
	return args.Get(0).(int), args.Get(1).(*RepositoryError)
}

func (this *MockRepo) GetRover(ID) (Rover, *RepositoryError) {
	args := this.Called()
	if args.Get(1) == nil {
		return args.Get(0).(Rover), nil
	}
	return args.Get(0).(Rover), args.Get(1).(*RepositoryError)
}

func (this *MockRepo) GetGame(ID) (*Game, *RepositoryError) {
	args := this.Called()
	if args.Get(1) == nil {
		return args.Get(0).(*Game), nil
	}
	return args.Get(0).(*Game), args.Get(1).(*RepositoryError)
}

func SuccessfulRepoFor(rover Rover) *MockRepo {
	repo := new(MockRepo)
	repo.On("GetRover", Anything).Return(rover, nil)
	repo.On("UpdateRover").Return(nil)
	repo.On("AddRover").Return(nil)
	repo.On("AddPlanet").Return(0, nil)
	repo.On("GetGame").Return(&Game{Rover: rover}, nil)
	return repo
}
