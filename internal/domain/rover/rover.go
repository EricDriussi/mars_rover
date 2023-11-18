package rover

import (
	"errors"
	coord "mars_rover/internal/domain/coordinate"
	"mars_rover/internal/domain/coordinate/absoluteCoordinate"
	. "mars_rover/internal/domain/coordinate/absoluteCoordinate"
	. "mars_rover/internal/domain/direction"
	. "mars_rover/internal/domain/planet"
	"mars_rover/internal/domain/rover/planetMap"
	. "mars_rover/internal/domain/rover/planetMap"
)

// TODO: extract
type Rover interface {
	TurnLeft()
	TurnRight()
	MoveForward() error
	MoveBackward() error
	Position() AbsoluteCoordinate
	Direction() Direction
	Map() Map
}

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

func (this *WrappingCollidingRover) TurnLeft() {
	this.direction = this.direction.DirectionOnTheLeft()
}

func (this *WrappingCollidingRover) TurnRight() {
	this.direction = this.direction.DirectionOnTheRight()
}

func (this *WrappingCollidingRover) MoveForward() error {
	futureCoordinate := *coord.SumOf(this.coordinate, this.direction.RelativePositionAhead())
	isOutOfBounds := this.planetMap.IsOutOfBounds(futureCoordinate)
	if isOutOfBounds {
		futureCoordinate = this.wrapAround(futureCoordinate)
	}
	willHitSomething := this.planetMap.CollidesWithObstacle(futureCoordinate)
	if willHitSomething {
		return errors.New("cannot move forward, something is in the way")
	}
	this.coordinate = futureCoordinate
	return nil
}

func (this *WrappingCollidingRover) MoveBackward() error {
	futureCoordinate := *coord.SumOf(this.coordinate, this.direction.RelativePositionBehind())
	isOutOfBounds := this.planetMap.IsOutOfBounds(futureCoordinate)
	if isOutOfBounds {
		futureCoordinate = this.wrapAround(futureCoordinate)
	}
	willHitSomething := this.planetMap.CollidesWithObstacle(futureCoordinate)
	if willHitSomething {
		return errors.New("cannot move backward, something is in the way")
	}
	this.coordinate = futureCoordinate
	return nil
}

func (this *WrappingCollidingRover) Position() AbsoluteCoordinate {
	return this.coordinate
}

func (this *WrappingCollidingRover) Direction() Direction {
	return this.direction
}

func (this *WrappingCollidingRover) Map() Map {
	return this.planetMap
}

func (this *WrappingCollidingRover) wrapAround(coordinate AbsoluteCoordinate) AbsoluteCoordinate {
	return *absoluteCoordinate.From(
		wrap(coordinate.X(), this.planetMap.Width()),
		wrap(coordinate.Y(), this.planetMap.Height()),
	)
}

func wrap(point, limit int) int {
	if point > limit {
		return 0
	} else if point < 0 {
		return limit
	}
	return point
}
