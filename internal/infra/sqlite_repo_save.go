package infra

import (
	"encoding/json"
	. "mars_rover/internal/domain/planet"
	. "mars_rover/internal/domain/rover"
	. "mars_rover/internal/domain/rover/godModRover"
	. "mars_rover/internal/infra/mappers"

	_ "github.com/mattn/go-sqlite3"
)

func (r *SQLiteRepository) SaveGame(rover Rover, planet Planet) error {
	planetId, err := r.savePlanet(planet)
	if err != nil {
		return err
	}
	return r.saveRover(rover, planetId)
}

func (r *SQLiteRepository) savePlanet(planet Planet) (int64, error) {
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

func (r *SQLiteRepository) saveRover(rover Rover, planetId int64) error {
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