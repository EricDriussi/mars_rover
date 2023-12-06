package persistence

import (
	"encoding/json"
	. "mars_rover/src/domain/rover"
	. "mars_rover/src/infra/persistence/mappers"

	_ "github.com/mattn/go-sqlite3"
)

func (r *SQLiteRepository) UpdateRover(rover Rover) error {

	roverAsBytes, err := json.Marshal(MapToPersistenceRover(rover))
	if err != nil {
		return err
	}
	roverAsString := string(roverAsBytes)
	_, err = r.db.Exec("UPDATE "+RoversTable+" SET rover = ? WHERE id = ?",
		roverAsString,
		rover.Id().String(),
	)
	return err
}