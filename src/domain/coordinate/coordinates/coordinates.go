package coordinates

import (
	"errors"
	. "mars_rover/src/domain/coordinate/absoluteCoordinate"
	. "mars_rover/src/domain/size"
	"sort"
)

type Coordinates struct {
	list []AbsoluteCoordinate
}

func New(list ...AbsoluteCoordinate) (*Coordinates, error) {
	if len(list) == 0 {
		return nil, errors.New("cannot create Coordinates with empty coordinate list")
	}
	return &Coordinates{filterUnique(list)}, nil
}

func filterUnique(list []AbsoluteCoordinate) []AbsoluteCoordinate {
	uniqueList := make([]AbsoluteCoordinate, 0)
	for _, coordinate := range list {
		if !contains(uniqueList, coordinate) {
			uniqueList = append(uniqueList, coordinate)
		}
	}
	return uniqueList
}

func (this *Coordinates) Contain(coordinate AbsoluteCoordinate) bool {
	return contains(this.list, coordinate)
}

func (this *Coordinates) Overflow(size Size) bool {
	for _, coordinate := range this.list {
		if coordinate.X() > size.Width() || coordinate.Y() > size.Height() {
			return true
		}
	}
	return false
}

func (this *Coordinates) List() []AbsoluteCoordinate {
	return this.list
}

func (this *Coordinates) First() AbsoluteCoordinate {
	return this.list[0]
}

func (this *Coordinates) HasOnlyOne() bool {
	return len(this.list) == 1
}

func (this *Coordinates) Amount() int {
	return len(this.list)
}

func (this *Coordinates) AreContiguous() bool {
	if len(this.list) <= 1 {
		return true
	}
	return areAllAdjacent(this.list)
}

func contains(list []AbsoluteCoordinate, coordinate AbsoluteCoordinate) bool {
	for _, containedCoordinate := range list {
		if containedCoordinate.Equals(coordinate) {
			return true
		}
	}
	return false
}

func areAllAdjacent(coordinates []AbsoluteCoordinate) bool {
	sortedCoords := make([]AbsoluteCoordinate, len(coordinates))
	copy(sortedCoords, coordinates)
	sort.Slice(sortedCoords, sortCoordinates(sortedCoords))
	for i := 0; i < len(sortedCoords)-1; i++ {
		currentCoordinate := sortedCoords[i]
		nextCoordinate := sortedCoords[i+1]
		if !currentCoordinate.IsAdjacentTo(nextCoordinate) {
			return false
		}
	}
	return true
}

func sortCoordinates(coordinates []AbsoluteCoordinate) func(i int, j int) bool {
	return func(i, j int) bool {
		if coordinates[i].X() != coordinates[j].X() {
			return coordinates[i].X() < coordinates[j].X()
		}
		return coordinates[i].Y() < coordinates[j].Y()
	}
}
