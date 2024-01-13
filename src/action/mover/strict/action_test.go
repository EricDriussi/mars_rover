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
	testRover := mocks.LandedRover(*absoluteCoordinate.Build(1, 1))
	repo := mocks.SuccessfulRepoFor(testRover)
	command := new(MockCommand)
	command.On("MapToRoverMovementFunction", testRover).Return(SuccessfulRoverFunc())
	commands := Commands{command}

	act := strict_mover.With(repo)
	movementResults, err := act.Move(id.New(), commands)

	assert.Nil(t, err)
	AssertEncounteredNoIssues(t, movementResults)
}

func TestMovementResultsContainAnIssueIfRoverReportsAnError(t *testing.T) {
	testRover := mocks.LandedRover(*absoluteCoordinate.Build(1, 1))
	repo := mocks.SuccessfulRepoFor(testRover)
	command := new(MockCommand)
	command.On("MapToRoverMovementFunction", testRover).Return(FailedRoverFunc())
	commands := Commands{command}

	act := strict_mover.With(repo)
	movementResults, err := act.Move(id.New(), commands)

	assert.Nil(t, err)
	AssertEncounteredAnIssue(t, movementResults)
}

func TestOnlyCallsRoverForGivenCommands(t *testing.T) {
	testRover := mocks.LandedRover(*absoluteCoordinate.Build(1, 1))
	mocks.MakeAlwaysSuccessful(testRover)
	repo := mocks.SuccessfulRepoFor(testRover)
	firstCommand := new(MockCommand)
	firstCommand.On("MapToRoverMovementFunction", testRover).Return(RoverFunc(testRover.MoveBackward))
	secondCommand := new(MockCommand)
	secondCommand.On("MapToRoverMovementFunction", testRover).Return(RoverFunc(testRover.MoveForward))
	commands := Commands{firstCommand, secondCommand}

	act := strict_mover.With(repo)
	_, err := act.Move(id.New(), commands)

	assert.Nil(t, err)
	testRover.AssertCalled(t, "MoveBackward")
	testRover.AssertCalled(t, "MoveForward")
	testRover.AssertNotCalled(t, "TurnLeft")
	testRover.AssertNotCalled(t, "TurnRight")
}

func TestReportsResultsBasedOnGivenCommandsOrder(t *testing.T) {
	testRover := mocks.LandedRover(*absoluteCoordinate.Build(1, 1))
	mocks.MakeAlwaysSuccessful(testRover)
	repo := mocks.SuccessfulRepoFor(testRover)
	firstCommand := new(MockCommand)
	firstCommand.On("MapToRoverMovementFunction", testRover).Return(RoverFunc(testRover.MoveBackward))
	secondCommand := new(MockCommand)
	secondCommand.On("MapToRoverMovementFunction", testRover).Return(RoverFunc(testRover.MoveForward))
	commands := Commands{firstCommand, secondCommand}

	act := strict_mover.With(repo)
	movementResults, err := act.Move(id.New(), commands)

	assert.Nil(t, err)
	AssertContainsOrderedCommands(t, movementResults, commands)
	AssertEncounteredNoIssues(t, movementResults)
}

func TestStopsCallingRoverForGivenCommandsOnceOneFails(t *testing.T) {
	testRover := mocks.LandedRover(*absoluteCoordinate.Build(1, 1))
	testRover.On("MoveBackward").Return(nil)
	testRover.On("MoveForward").Return(errors.New("movement blocked"))
	repo := mocks.SuccessfulRepoFor(testRover)
	firstCommand := new(MockCommand)
	firstCommand.On("MapToRoverMovementFunction", testRover).Return(RoverFunc(testRover.MoveBackward))
	failedCommand := new(MockCommand)
	failedCommand.On("MapToRoverMovementFunction", testRover).Return(RoverFunc(testRover.MoveForward))
	thirdCommand := new(MockCommand)
	thirdCommand.On("MapToRoverMovementFunction", testRover).Return(RoverFunc(testRover.MoveBackward))
	commands := Commands{firstCommand, failedCommand, thirdCommand}

	act := strict_mover.With(repo)
	_, err := act.Move(id.New(), commands)

	assert.Nil(t, err)
	testRover.AssertNumberOfCalls(t, "MoveBackward", 1)
	testRover.AssertCalled(t, "MoveForward")
}

func TestReportsRepoErrorWhenGettingRover(t *testing.T) {
	repo := new(MockRepo)
	repo.On("GetRover", Anything).Return(new(MockRover), aRepoError())
	commands := Commands{new(MockCommand)}

	act := strict_mover.With(repo)
	movementResults, err := act.Move(id.New(), commands)

	assert.Empty(t, movementResults)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "not found")
}

func TestReportsRepoErrorWhenUpdatingRover(t *testing.T) {
	testRover := mocks.LandedRover(*absoluteCoordinate.Build(1, 1))
	repo := new(MockRepo)
	repo.On("GetRover", Anything).Return(testRover, nil)
	repo.On("UpdateRover").Return(aRepoError())
	irrelevantCommand := new(MockCommand)
	irrelevantCommand.On("MapToRoverMovementFunction", testRover).Return(SuccessfulRoverFunc())
	commands := Commands{irrelevantCommand}

	act := strict_mover.With(repo)
	movementResults, err := act.Move(id.New(), commands)

	assert.Empty(t, movementResults)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "failed to update")
}

func aRepoError() *RepositoryError {
	return PersistenceMalfunction(errors.New("whatever"))
}
