package persistence

import (
	"encoding/json"
	. "mars_rover/src/domain"
	. "mars_rover/src/domain/planet"
	. "mars_rover/src/domain/rover"
	. "mars_rover/src/domain/rover/godModRover"
	. "mars_rover/src/infra/persistence/mappers"

	_ "github.com/mattn/go-sqlite3"
)

func (this *SQLiteRepository) AddPlanet(planet Planet) (int, *RepositoryError) {
	planetAsBytes, err := json.Marshal(MapToPersistencePlanet(planet))
	if err != nil {
		return -1, CouldNotMap(err)
	}

	planetInsertResult, err := this.db.Exec("INSERT INTO "+PlanetsTable+" (planet) VALUES (?)",
		string(planetAsBytes))
	if err != nil {
		return -1, CouldNotAdd(err)
	}
	num, err := planetInsertResult.LastInsertId()
	if err != nil {
		return -1, PersistenceMalfunction(err)
	}
	return int(num), nil
}

func (this *SQLiteRepository) AddRover(rover Rover, planetId int) *RepositoryError {
	roverAsBytes, err := json.Marshal(MapToPersistenceRover(rover))
	if err != nil {
		return CouldNotMap(err)
	}

	_, err = this.db.Exec("INSERT INTO "+RoversTable+" (id, rover, type, planet_id) VALUES (?, ?, ?, ?)",
		rover.Id().String(),
		string(roverAsBytes),
		typeOf(rover),
		planetId,
	)
	if err != nil {
		return CouldNotAdd(err)
	}
	return nil
}

func typeOf(rover Rover) string {
	if _, ok := rover.(*GodModRover); ok {
		return "godmod"
	}
	return "normal"
}
