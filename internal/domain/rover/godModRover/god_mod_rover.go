package godModRover

import (
	"github.com/google/uuid"
	. "github.com/google/uuid"
	"mars_rover/internal/domain/coordinate"
	. "mars_rover/internal/domain/coordinate/absoluteCoordinate"
	. "mars_rover/internal/domain/planet"
	. "mars_rover/internal/domain/rover/direction"
	"mars_rover/internal/domain/rover/planetMap"
	. "mars_rover/internal/domain/rover/planetMap"
)

type GodModRover struct {
	id         UUID
	planetMap  Map
	coordinate AbsoluteCoordinate
	direction  Direction
}

func Land(coordinate AbsoluteCoordinate, planet Planet) *GodModRover {
	return &GodModRover{
		id:         uuid.New(),
		planetMap:  *planetMap.OfPlanet(planet),
		coordinate: coordinate,
		direction:  North{},
	}
}

// TODO.LM: should be 4 separate constructors: LandFacing{North, East, South, West}
func LandFacing(direction Direction, coordinate AbsoluteCoordinate, planet Planet) *GodModRover {
	return &GodModRover{
		id:         uuid.New(),
		planetMap:  *planetMap.OfPlanet(planet),
		coordinate: coordinate,
		direction:  direction,
	}
}

func (this *GodModRover) MoveForward() error {
	this.coordinate = *coordinate.SumOf(this.coordinate, this.direction.RelativeCoordinateAhead())
	return nil
}

func (this *GodModRover) MoveBackward() error {
	this.coordinate = *coordinate.SumOf(this.coordinate, this.direction.RelativeCoordinateBehind())
	return nil
}

func (this *GodModRover) TurnLeft() {
	this.direction = this.direction.DirectionOnTheLeft()
}

func (this *GodModRover) TurnRight() {
	this.direction = this.direction.DirectionOnTheRight()
}

func (this *GodModRover) Id() UUID {
	return this.id
}

func (this *GodModRover) Coordinate() AbsoluteCoordinate {
	return this.coordinate
}

func (this *GodModRover) Direction() Direction {
	return this.direction
}

func (this *GodModRover) Map() Map {
	return this.planetMap
}
