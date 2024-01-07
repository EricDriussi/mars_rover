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
		return 0, CouldNotMap(err)
	}
	alreadyExists, err := this.checkIfPlanetAlreadyExists(planet)
	if err != nil {
		return 0, PersistenceMalfunction(err)
	}
	if alreadyExists {
		num, err := this.getPlanetId(planet)
		if err != nil {
			return 0, PersistenceMalfunction(err)
		}
		return num, AlreadyExists()
	}

	planetInsertResult, err := this.db.Exec("INSERT INTO "+PlanetsTable+" (planet) VALUES (?)",
		string(planetAsBytes))
	if err != nil {
		return 0, CouldNotAdd(err)
	}
	num, err := planetInsertResult.LastInsertId()
	if err != nil {
		return 0, PersistenceMalfunction(err)
	}
	return int(num), nil
}

func (this *SQLiteRepository) checkIfPlanetAlreadyExists(planet Planet) (bool, error) {
	planetAsBytes, err := json.Marshal(MapToPersistencePlanet(planet))
	if err != nil {
		return false, err
	}
	var count int
	err = this.db.QueryRow("SELECT COUNT(*) FROM "+PlanetsTable+" WHERE planet = ?", string(planetAsBytes)).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (this *SQLiteRepository) getPlanetId(planet Planet) (int, error) {
	planetAsBytes, err := json.Marshal(MapToPersistencePlanet(planet))
	if err != nil {
		return 0, err
	}
	var planetId int
	err = this.db.QueryRow("SELECT id FROM "+PlanetsTable+" WHERE planet = ?", string(planetAsBytes)).Scan(&planetId)
	if err != nil {
		return 0, err
	}
	return planetId, nil
}

func (this *SQLiteRepository) AddRover(rover Rover, planetId int) *RepositoryError {
	roverAsBytes, err := json.Marshal(MapToPersistenceRover(rover))
	if err != nil {
		return CouldNotMap(err)
	}
	alreadyExists, err := this.checkIfRoverAlreadyExists(rover)
	if err != nil {
		return PersistenceMalfunction(err)
	}

	if alreadyExists {
		return AlreadyExists()
	}

	_, err = this.db.Exec("INSERT INTO "+RoversTable+" (id, rover, godMod, planet_id) VALUES (?, ?, ?, ?)",
		rover.Id().String(),
		string(roverAsBytes),
		isGodMod(rover),
		planetId,
	)
	if err != nil {
		return CouldNotAdd(err)
	}
	return nil
}

func (this *SQLiteRepository) checkIfRoverAlreadyExists(rover Rover) (bool, error) {
	roverAsBytes, err := json.Marshal(MapToPersistenceRover(rover))
	if err != nil {
		return false, err
	}
	var count int
	err = this.db.QueryRow("SELECT COUNT(*) FROM "+RoversTable+" WHERE rover = ?", string(roverAsBytes)).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func isGodMod(rover Rover) bool {
	_, isGodMod := rover.(*GodModRover)
	return isGodMod
}
