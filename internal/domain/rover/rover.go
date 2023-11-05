package rover

import (
	"errors"
	"mars_rover/internal/domain/location"
	"mars_rover/internal/domain/planet"
	planetMap "mars_rover/internal/domain/rover/planet_map"
)

type IRover interface {
	TurnLeft()
	TurnRight()
	MoveForward() error
	MoveBackward() error
}

type Rover struct {
	location  location.Location
	planetMap planetMap.Map
}

func Land(location location.Location, planet planet.Planet) (*Rover, error) {
	mapOfPlanet := planetMap.Of(planet)
	if mapOfPlanet.CollidesWithObstacle(location.Position()) {
		return nil, errors.New("cannot land on obstacle")
	}
	return &Rover{location: location, planetMap: *mapOfPlanet}, nil
}

func (this Rover) Location() *location.Location {
	return &this.location
}

func (this *Rover) TurnLeft() {
	this.location.FaceLeft()
}

func (this *Rover) TurnRight() {
	this.location.FaceRight()
}

func (this *Rover) MoveForward() error {
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

func (this *Rover) MoveBackward() error {
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

func (this Rover) willBeOutOfBounds() bool {
	return this.planetMap.IsOutOfBounds(this.location.WillBeAt())
}

func (this Rover) willHitSomething() bool {
	return this.planetMap.CollidesWithObstacle(this.location.WillBeAt())
}
