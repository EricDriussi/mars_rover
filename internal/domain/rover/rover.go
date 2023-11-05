package rover

import (
	"errors"
	"mars_rover/internal/domain/location"
	"mars_rover/internal/domain/planet"
	planetMap "mars_rover/internal/domain/rover/planet_map"
)

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

func (this *Rover) MoveForward() {
	this.location.CalculatePositionAhead()
	if this.willBeOutOfBounds() {
		this.location.WrapAround(this.planetMap.Size())
	}
	if this.willHitSomething() {
		// TODO: how do I "report the obstacle"?
		this.location.Reset()
		return
	}
	this.location.UpdatePosition()
}

func (this *Rover) MoveBackward() {
	this.location.CalculatePositionBehind()
	if this.willBeOutOfBounds() {
		this.location.WrapAround(this.planetMap.Size())
	}
	if this.willHitSomething() {
		// TODO: how do I "report the obstacle"?
		this.location.Reset()
		return
	}
	this.location.UpdatePosition()
}

func (this Rover) willBeOutOfBounds() bool {
	return this.planetMap.IsOutOfBounds(this.location.WillBeAt())
}

func (this Rover) willHitSomething() bool {
	return this.planetMap.CollidesWithObstacle(this.location.WillBeAt())
}
