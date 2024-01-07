package persistence

import (
	"encoding/json"
	. "mars_rover/src/domain"
	. "mars_rover/src/domain/rover"
	. "mars_rover/src/infra/persistence/mappers"

	_ "github.com/mattn/go-sqlite3"
)

func (this *SQLiteRepository) UpdateRover(rover Rover) *RepositoryError {
	roverAsBytes, err := json.Marshal(MapToPersistenceRover(rover))
	if err != nil {
		return CouldNotMap(err)
	}
	roverAsString := string(roverAsBytes)
	_, err = this.db.Exec("UPDATE "+RoversTable+" SET rover = ? WHERE id = ?",
		roverAsString,
		rover.Id().String(),
	)
	if err != nil {
		return CouldNotUpdate(err)
	}
	return nil
}
