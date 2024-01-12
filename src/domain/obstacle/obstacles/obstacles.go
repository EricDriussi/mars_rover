package obstacles

import (
	"errors"
	. "mars_rover/src/domain/coordinate/absoluteCoordinate"
	. "mars_rover/src/domain/obstacle"
	. "mars_rover/src/domain/size"
)

type Obstacles struct {
	list []Obstacle
}

func FromList(list ...Obstacle) (*Obstacles, error) {
	if len(list) < 1 {
		return nil, errors.New("cannot create Obstacles with empty obstacle list")
	}
	if containsOverlappingCoordinates(list...) {
		return nil, errors.New("invalid Obstacles: different obstacles share the same Coordinate(s)")
	}
	return &Obstacles{list}, nil
}

func containsOverlappingCoordinates(obstacleList ...Obstacle) bool {
	coordinates := make(map[AbsoluteCoordinate]bool)
	for _, obstacle := range obstacleList {
		obstacleCoordinates := obstacle.Coordinates()
		for _, singleCoordinate := range obstacleCoordinates.List() {
			if coordinates[singleCoordinate] {
				return true
			}
			coordinates[singleCoordinate] = true
		}
	}
	return false
}

func Empty() *Obstacles {
	return &Obstacles{[]Obstacle{}}
}

func (this *Obstacles) List() []Obstacle {
	return this.list
}

func (this *Obstacles) Occupy(coordinate AbsoluteCoordinate) bool {
	for _, obs := range this.list {
		if obs.Occupies(coordinate) {
			return true
		}
	}
	return false
}

func (this *Obstacles) IsAnyBeyond(size Size) bool {
	for _, obs := range this.list {
		if obs.IsBeyond(size) {
			return true
		}
	}
	return false
}

func (this *Obstacles) Amount() int {
	return len(this.list)
}

func (this *Obstacles) Add(obstacle Obstacle) error {
	if this.anyCoordinateOverlaps(obstacle) {
		return errors.New("cannot add obstacle with overlapping coordinates")
	}
	this.list = append(this.list, obstacle)
	return nil
}

func (this *Obstacles) anyCoordinateOverlaps(obstacle Obstacle) bool {
	for _, obs := range this.list {
		coordinates := obs.Coordinates()
		if coordinates.ContainAnyOf(obstacle.Coordinates()) {
			return true
		}
	}
	return false
}
