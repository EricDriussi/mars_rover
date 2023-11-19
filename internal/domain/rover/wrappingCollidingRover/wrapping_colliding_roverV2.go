package wrappingCollidingRover

import (
	"errors"
	. "mars_rover/internal/domain/coordinate/absoluteCoordinate"
	. "mars_rover/internal/domain/planet"
	. "mars_rover/internal/domain/rover/direction"
	"mars_rover/internal/domain/rover/planetMap"
	. "mars_rover/internal/domain/rover/planetMap"
	"mars_rover/internal/domain/rover/wrappingCollidingRover/gps"
	. "mars_rover/internal/domain/rover/wrappingCollidingRover/gps"
)

type WrappingCollidingRoverV2 struct {
	planetMap  Map
	coordinate AbsoluteCoordinate
	direction  Direction
	gps        GPS
}

func LandV2(coordinate AbsoluteCoordinate, planet Planet) (*WrappingCollidingRoverV2, error) {
	mapOfPlanet := planetMap.Of(planet)
	if mapOfPlanet.HasObstacleIn(coordinate) {
		return nil, errors.New("cannot land on obstacle")
	}
	newRover := &WrappingCollidingRoverV2{
		planetMap:  *mapOfPlanet,
		coordinate: coordinate,
		direction:  North{},
	}
	newRover.gps = gps.Bind(newRover)
	return newRover, nil
}

// TODO.LM: should be LandFacing{North, East, South, West}
func LandFacingV2(direction Direction, coordinate AbsoluteCoordinate, planet Planet) (*WrappingCollidingRoverV2, error) {
	mapOfPlanet := planetMap.Of(planet)
	if mapOfPlanet.HasObstacleIn(coordinate) {
		return nil, errors.New("cannot land on obstacle")
	}
	newRover := &WrappingCollidingRoverV2{
		planetMap:  *mapOfPlanet,
		coordinate: coordinate,
		direction:  direction,
	}
	newRover.gps = gps.Bind(newRover)
	return newRover, nil
}

func (this *WrappingCollidingRoverV2) MoveForward() error {
	futureCoordinate := this.gps.Ahead()
	willHitSomething := this.planetMap.HasObstacleIn(futureCoordinate)
	if willHitSomething {
		return errors.New("cannot move forward, something is in the way")
	}
	this.coordinate = futureCoordinate
	return nil
}

func (this *WrappingCollidingRoverV2) MoveBackward() error {
	futureCoordinate := this.gps.Behind()
	willHitSomething := this.planetMap.HasObstacleIn(futureCoordinate)
	if willHitSomething {
		return errors.New("cannot move backward, something is in the way")
	}
	this.coordinate = futureCoordinate
	return nil
}

func (this *WrappingCollidingRoverV2) TurnLeft() {
	this.direction = this.direction.DirectionOnTheLeft()
}

func (this *WrappingCollidingRoverV2) TurnRight() {
	this.direction = this.direction.DirectionOnTheRight()
}

func (this *WrappingCollidingRoverV2) Coordinate() AbsoluteCoordinate {
	return this.coordinate
}

func (this *WrappingCollidingRoverV2) Direction() Direction {
	return this.direction
}

func (this *WrappingCollidingRoverV2) Map() Map {
	return this.planetMap
}
