package persistence

import (
	"encoding/json"
	. "mars_rover/src/domain/planet"
	. "mars_rover/src/domain/rover"
	. "mars_rover/src/domain/rover/godModRover"
	. "mars_rover/src/infra/persistence/mappers"

	_ "github.com/mattn/go-sqlite3"
)

func (r *SQLiteRepository) SaveGame(rover Rover, planet Planet) error {
	planetId, err := r.AddPlanet(planet)
	if err != nil {
		return err
	}
	return r.AddRover(rover, planetId)
}

func (r *SQLiteRepository) AddPlanet(planet Planet) (int64, error) {
	planetAsBytes, err := json.Marshal(MapToPersistencePlanet(planet))
	if err != nil {
		return 0, err
	}

	planetInsertResult, err := r.db.Exec("INSERT INTO "+PlanetsTable+" (planet) VALUES (?)",
		string(planetAsBytes))
	if err != nil {
		return 0, err
	}
	return planetInsertResult.LastInsertId()
}

func (r *SQLiteRepository) AddRover(rover Rover, planetId int64) error {
	roverAsBytes, err := json.Marshal(MapToPersistenceRover(rover))
	if err != nil {
		return err
	}

	_, err = r.db.Exec("INSERT INTO "+RoversTable+" (id, rover, godMod, planet_id) VALUES (?, ?, ?, ?)",
		rover.Id().String(),
		string(roverAsBytes),
		isGodMod(rover),
		planetId,
	)
	return err
}

func isGodMod(rover Rover) bool {
	_, isGodMod := rover.(*GodModRover)
	return isGodMod
}
