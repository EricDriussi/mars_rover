package strict_mover_test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	. "github.com/stretchr/testify/mock"
	. "mars_rover/src/action/mover/command"
	"mars_rover/src/action/mover/strict"
	. "mars_rover/src/domain"
	"mars_rover/src/domain/coordinate/absoluteCoordinate"
	"mars_rover/src/domain/rover/id"
	. "mars_rover/src/test_helpers"
	"mars_rover/src/test_helpers/mocks"
	. "mars_rover/src/test_helpers/mocks"
	"testing"
)

func TestMovementResultsContainNoIssueIfRoverReportsNoError(t *testing.T) {
	rover := mocks.LandedRover(*absoluteCoordinate.Build(1, 1))
	repo := mocks.SuccessfulRepoFor(rover)
	command := new(MockCommand)
	command.On("MapToRoverMovementFunction", rover).Return(SuccessfulRoverFunc())
	commands := Commands{command}
	moveAction := strict_mover.With(repo)

	movementResults, err := moveAction.Move(id.New(), commands)

	assert.Nil(t, err)
	AssertEncounteredNoIssues(t, movementResults)
}

func TestMovementResultsContainAnIssueIfRoverReportsAnError(t *testing.T) {
	rover := mocks.LandedRover(*absoluteCoordinate.Build(1, 1))
	repo := mocks.SuccessfulRepoFor(rover)
	command := new(MockCommand)
	command.On("MapToRoverMovementFunction", rover).Return(FailedRoverFunc())
	commands := Commands{command}
	moveAction := strict_mover.With(repo)

	movementResults, err := moveAction.Move(id.New(), commands)

	assert.Nil(t, err)
	AssertEncounteredAnIssue(t, movementResults)
}

func TestOnlyCallsRoverForGivenCommands(t *testing.T) {
	rover := mocks.LandedRover(*absoluteCoordinate.Build(1, 1))
	mocks.MakeAlwaysSuccessful(rover)
	repo := mocks.SuccessfulRepoFor(rover)
	firstCommand := new(MockCommand)
	firstCommand.On("MapToRoverMovementFunction", rover).Return(RoverFunc(rover.MoveBackward))
	secondCommand := new(MockCommand)
	secondCommand.On("MapToRoverMovementFunction", rover).Return(RoverFunc(rover.MoveForward))
	commands := Commands{firstCommand, secondCommand}
	moveAction := strict_mover.With(repo)

	_, err := moveAction.Move(id.New(), commands)

	assert.Nil(t, err)
	rover.AssertCalled(t, "MoveBackward")
	rover.AssertCalled(t, "MoveForward")
	rover.AssertNotCalled(t, "TurnLeft")
	rover.AssertNotCalled(t, "TurnRight")
}

func TestReportsResultsBasedOnGivenCommandsOrder(t *testing.T) {
	rover := mocks.LandedRover(*absoluteCoordinate.Build(1, 1))
	mocks.MakeAlwaysSuccessful(rover)
	repo := mocks.SuccessfulRepoFor(rover)
	firstCommand := new(MockCommand)
	firstCommand.On("MapToRoverMovementFunction", rover).Return(RoverFunc(rover.MoveBackward))
	secondCommand := new(MockCommand)
	secondCommand.On("MapToRoverMovementFunction", rover).Return(RoverFunc(rover.MoveForward))
	commands := Commands{firstCommand, secondCommand}
	moveAction := strict_mover.With(repo)

	movementResults, err := moveAction.Move(id.New(), commands)

	assert.Nil(t, err)
	AssertContainsOrderedCommands(t, movementResults, commands)
	AssertEncounteredNoIssues(t, movementResults)
}

func TestStopsCallingRoverForGivenCommandsOnceOneFails(t *testing.T) {
	rover := mocks.LandedRover(*absoluteCoordinate.Build(1, 1))
	rover.On("MoveBackward").Return(nil)
	rover.On("MoveForward").Return(errors.New("movement blocked"))
	repo := mocks.SuccessfulRepoFor(rover)
	firstCommand := new(MockCommand)
	firstCommand.On("MapToRoverMovementFunction", rover).Return(RoverFunc(rover.MoveBackward))
	failedCommand := new(MockCommand)
	failedCommand.On("MapToRoverMovementFunction", rover).Return(RoverFunc(rover.MoveForward))
	thirdCommand := new(MockCommand)
	thirdCommand.On("MapToRoverMovementFunction", rover).Return(RoverFunc(rover.MoveBackward))
	commands := Commands{firstCommand, failedCommand, thirdCommand}
	moveAction := strict_mover.With(repo)

	_, err := moveAction.Move(id.New(), commands)

	assert.Nil(t, err)
	rover.AssertNumberOfCalls(t, "MoveBackward", 1)
	rover.AssertCalled(t, "MoveForward")
}

func TestReportsRepoErrorWhenGettingRover(t *testing.T) {
	repo := new(MockRepo)
	repo.On("GetRover", Anything).Return(new(MockRover), aRepoError())
	commands := Commands{new(MockCommand)}
	moveAction := strict_mover.With(repo)

	movementResults, err := moveAction.Move(id.New(), commands)

	assert.Empty(t, movementResults)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not found")
}

func TestReportsRepoErrorWhenUpdatingRover(t *testing.T) {
	rover := mocks.LandedRover(*absoluteCoordinate.Build(1, 1))
	repo := new(MockRepo)
	repo.On("GetRover", Anything).Return(rover, nil)
	repo.On("UpdateRover").Return(aRepoError())
	irrelevantCommand := new(MockCommand)
	irrelevantCommand.On("MapToRoverMovementFunction", rover).Return(SuccessfulRoverFunc())
	commands := Commands{irrelevantCommand}
	moveAction := strict_mover.With(repo)

	movementResults, err := moveAction.Move(id.New(), commands)

	assert.Empty(t, movementResults)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to update")
}

func aRepoError() *RepositoryError {
	return PersistenceMalfunction(errors.New("whatever"))
}
