package wrappingCollidingRover

import (
	"errors"
	. "mars_rover/internal/domain/coordinate/absoluteCoordinate"
	. "mars_rover/internal/domain/direction"
	. "mars_rover/internal/domain/planet"
	"mars_rover/internal/domain/rover/planetMap"
	. "mars_rover/internal/domain/rover/planetMap"
	"mars_rover/internal/domain/rover/wrappingCollidingRover/positionCalculator"
)

type WrappingCollidingRover struct {
	planetMap  Map
	coordinate AbsoluteCoordinate
	direction  Direction
}

func Land(coordinate AbsoluteCoordinate, planet Planet) (*WrappingCollidingRover, error) {
	mapOfPlanet := planetMap.Of(planet)
	if mapOfPlanet.CollidesWithObstacle(coordinate) {
		return nil, errors.New("cannot land on obstacle")
	}
	return &WrappingCollidingRover{
		planetMap:  *mapOfPlanet,
		coordinate: coordinate,
		direction:  North{},
	}, nil
}

func LandFacing(direction Direction, coordinate AbsoluteCoordinate, planet Planet) (*WrappingCollidingRover, error) {
	mapOfPlanet := planetMap.Of(planet)
	if mapOfPlanet.CollidesWithObstacle(coordinate) {
		return nil, errors.New("cannot land on obstacle")
	}
	return &WrappingCollidingRover{
		planetMap:  *mapOfPlanet,
		coordinate: coordinate,
		direction:  direction,
	}, nil
}

func (this *WrappingCollidingRover) MoveForward() error {
	futureCoordinate := positionCalculator.Forward(this.direction, this.coordinate, this.planetMap)
	willHitSomething := this.planetMap.CollidesWithObstacle(futureCoordinate)
	if willHitSomething {
		return errors.New("cannot move forward, something is in the way")
	}
	this.coordinate = futureCoordinate
	return nil
}

func (this *WrappingCollidingRover) MoveBackward() error {
	futureCoordinate := positionCalculator.Backward(this.direction, this.coordinate, this.planetMap)
	willHitSomething := this.planetMap.CollidesWithObstacle(futureCoordinate)
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

func (this *WrappingCollidingRover) Coordinate() AbsoluteCoordinate {
	return this.coordinate
}

func (this *WrappingCollidingRover) Direction() Direction {
	return this.direction
}

func (this *WrappingCollidingRover) Map() Map {
	return this.planetMap
}
