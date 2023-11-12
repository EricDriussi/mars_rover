package rover

import (
	"errors"
	"mars_rover/internal/domain/location"
	"mars_rover/internal/domain/planet"
	planetMap "mars_rover/internal/domain/rover/planet_map"
)

// TODO: extract
type Rover interface {
	TurnLeft()
	TurnRight()
	MoveForward() error
	MoveBackward() error
	Location() *location.Location
	Map() *planetMap.Map
}

type WrappingCollidingRover struct {
	location  location.Location
	planetMap planetMap.Map
}

func Land(location location.Location, planet planet.Planet) (*WrappingCollidingRover, error) {
	mapOfPlanet := planetMap.Of(planet)
	if mapOfPlanet.CollidesWithObstacle(location.Position()) {
		return nil, errors.New("cannot land on obstacle")
	}
	return &WrappingCollidingRover{location: location, planetMap: *mapOfPlanet}, nil
}

func (this WrappingCollidingRover) Location() *location.Location {
	return &this.location
}

func (this WrappingCollidingRover) Map() *planetMap.Map {
	return &this.planetMap
}

func (this *WrappingCollidingRover) TurnLeft() {
	this.location.FaceLeft()
}

func (this *WrappingCollidingRover) TurnRight() {
	this.location.FaceRight()
}

func (this *WrappingCollidingRover) MoveForward() error {
	this.location.CalculatePositionAhead()
	if this.willBeOutOfBounds() {
		this.location.WrapAround(this.planetMap.Size())
	}
	if this.willHitSomething() {
		this.location.Reset()
		return errors.New("cannot move forward, something is in the way")
	}
	this.location.UpdatePosition()
	return nil
}

func (this *WrappingCollidingRover) MoveBackward() error {
	this.location.CalculatePositionBehind()
	if this.willBeOutOfBounds() {
		this.location.WrapAround(this.planetMap.Size())
	}
	if this.willHitSomething() {
		this.location.Reset()
		return errors.New("cannot move backward, something is in the way")
	}
	this.location.UpdatePosition()
	return nil
}

func (this WrappingCollidingRover) willBeOutOfBounds() bool {
	return this.planetMap.IsOutOfBounds(this.location.WillBeAt())
}

func (this WrappingCollidingRover) willHitSomething() bool {
	return this.planetMap.CollidesWithObstacle(this.location.WillBeAt())
}
