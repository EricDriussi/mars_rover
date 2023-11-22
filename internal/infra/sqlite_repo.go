package infra

import (
	"database/sql"
	"encoding/json"
	"log"
	. "mars_rover/internal/domain/planet"
	. "mars_rover/internal/domain/rover"
	. "mars_rover/internal/domain/rover/godModRover"
	. "mars_rover/internal/infra/mappers"

	_ "github.com/mattn/go-sqlite3"
)

const (
	RoversTable  = "rovers"
	PlanetsTable = "planets"
)

type SQLiteRepository struct {
	db *sql.DB
}

func NewSQLite(db *sql.DB) *SQLiteRepository {
	return &SQLiteRepository{db: db}
}

func (r *SQLiteRepository) SaveGame(rover Rover, planet Planet) error {
	planetId, err := r.savePlanet(planet)
	if err != nil {
		return err
	}

	return r.saveRover(rover, planetId)
}

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

func InitMem() *sql.DB {
	return setup(":memory:")
}

func InitFS() *sql.DB {
	return setup("./mars_rover.db")

}

func setup(location string) *sql.DB {
	db, err := sql.Open("sqlite3", location)
	if err != nil {
		log.Fatal(err)
	}

	createTables(db)
	return db
}

func createTables(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS ` + RoversTable + ` (
		id TEXT PRIMARY KEY,
		rover TEXT NOT NULL,
		godMod BOOLEAN NOT NULL,
		planet_id INTEGER,
		FOREIGN KEY(planet_id) REFERENCES ` + PlanetsTable + `(id)
	);
	CREATE TABLE IF NOT EXISTS ` + PlanetsTable + ` (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		planet TEXT NOT NULL
	);
`)
	if err != nil {
		log.Fatal(err)
	}
}
