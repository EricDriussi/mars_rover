package wrappingCollidingRover

import (
	"errors"
	. "mars_rover/src/domain/coordinate/absoluteCoordinate"
	. "mars_rover/src/domain/planet"
	. "mars_rover/src/domain/rover/direction"
	"mars_rover/src/domain/rover/planetMap"
	. "mars_rover/src/domain/rover/planetMap"
	. "mars_rover/src/domain/rover/uuid"
	"mars_rover/src/domain/rover/wrappingCollidingRover/gps"
	. "mars_rover/src/domain/rover/wrappingCollidingRover/gps"
)

type WrappingCollidingRover struct {
	id         UUID
	planetMap  Map
	coordinate AbsoluteCoordinate
	direction  Direction
	gps        GPS
}

func LandFacing(id UUID, direction Direction, coordinate AbsoluteCoordinate, planet Planet) (*WrappingCollidingRover, error) {
	mapOfPlanet := planetMap.OfPlanet(planet)
	if mapOfPlanet.HasObstacleIn(coordinate) {
		return nil, errors.New("cannot land on obstacle")
	}
	if mapOfPlanet.IsOutOfBounds(coordinate) {
		return nil, errors.New("cannot land out of planet")
	}
	newRover := &WrappingCollidingRover{
		id:         id,
		planetMap:  *mapOfPlanet,
		coordinate: coordinate,
		direction:  direction,
	}
	newRover.gps = gps.Bind(newRover)
	return newRover, nil
}

func (this *WrappingCollidingRover) MoveForward() error {
	futureCoordinate := this.gps.Ahead()
	willHitSomething := this.planetMap.HasObstacleIn(futureCoordinate)
	if willHitSomething {
		return errors.New("cannot move forward, something is in the way")
	}
	this.coordinate = futureCoordinate
	return nil
}

func (this *WrappingCollidingRover) MoveBackward() error {
	futureCoordinate := this.gps.Behind()
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
