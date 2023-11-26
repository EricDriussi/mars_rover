package wrappingCollidingRover

import (
	"errors"
	"github.com/google/uuid"
	. "github.com/google/uuid"
	. "mars_rover/internal/domain/coordinate/absoluteCoordinate"
	. "mars_rover/internal/domain/planet"
	. "mars_rover/internal/domain/rover/direction"
	"mars_rover/internal/domain/rover/planetMap"
	. "mars_rover/internal/domain/rover/planetMap"
	"mars_rover/internal/domain/rover/wrappingCollidingRover/positionCalculator"
)

// TODO: use V2
type WrappingCollidingRover struct {
	id         UUID
	planetMap  Map
	coordinate AbsoluteCoordinate
	direction  Direction
}

func Land(coordinate AbsoluteCoordinate, planet Planet) (*WrappingCollidingRover, error) {
	mapOfPlanet := planetMap.OfPlanet(planet)
	if mapOfPlanet.HasObstacleIn(coordinate) {
		return nil, errors.New("cannot land on obstacle")
	}
	if mapOfPlanet.IsOutOfBounds(coordinate) {
		return nil, errors.New("cannot land out of planet")
	}
	newRover := &WrappingCollidingRover{
		id:         uuid.New(),
		planetMap:  *mapOfPlanet,
		coordinate: coordinate,
		direction:  North{},
	}
	return newRover, nil
}

// TODO.LM: should be LandFacing{North, East, South, West}
func LandFacing(direction Direction, coordinate AbsoluteCoordinate, planet Planet) (*WrappingCollidingRover, error) {
	mapOfPlanet := planetMap.OfPlanet(planet)
	if mapOfPlanet.HasObstacleIn(coordinate) {
		return nil, errors.New("cannot land on obstacle")
	}
	if mapOfPlanet.IsOutOfBounds(coordinate) {
		return nil, errors.New("cannot land out of planet")
	}
	newRover := &WrappingCollidingRover{
		id:         uuid.New(),
		planetMap:  *mapOfPlanet,
		coordinate: coordinate,
		direction:  direction,
	}
	return newRover, nil
}

func (this *WrappingCollidingRover) MoveForward() error {
	futureCoordinate := positionCalculator.Forward(this.direction, this.coordinate, this.planetMap)
	willHitSomething := this.planetMap.HasObstacleIn(futureCoordinate)
	if willHitSomething {
		return errors.New("cannot move forward, something is in the way")
	}
	this.coordinate = futureCoordinate
	return nil
}

func (this *WrappingCollidingRover) MoveBackward() error {
	futureCoordinate := positionCalculator.Backward(this.direction, this.coordinate, this.planetMap)
	willHitSomething := this.planetMap.HasObstacleIn(futureCoordinate)
	if willHitSomething {
		return errors.New("cannot move backward, something is in the way")
	}
	this.coordinate = futureCoordinate
	return nil
}

func (this *WrappingCollidingRover) TurnLeft() {
	this.direction = this.direction.DirectionOnTheLeft()
}

func (this *WrappingCollidingRover) TurnRight() {
	this.direction = this.direction.DirectionOnTheRight()
}

func (this *WrappingCollidingRover) Id() UUID {
	return this.id
}

func (this *WrappingCollidingRover) Coordinate() AbsoluteCoordinate {
	return this.coordinate
}

func (this *WrappingCollidingRover) Direction() Direction {
	return this.direction
}

func (this *WrappingCollidingRover) Map() Map {
	return this.planetMap
}
