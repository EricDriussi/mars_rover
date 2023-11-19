package infra

import (
	"database/sql"
	"encoding/json"
	"log"
	. "mars_rover/internal/domain/planet/rockyPlanet"
	. "mars_rover/internal/domain/rover/wrappingCollidingRover"
	. "mars_rover/internal/infra/mappers"

	_ "github.com/mattn/go-sqlite3"
)

const (
	WrappingRoversTable = "wrapping_rovers"
	RockyPlanetsTable   = "rocky_planets"
)

type SQLiteRepository struct {
	db *sql.DB
}

func NewSQLite(db *sql.DB) *SQLiteRepository {
	return &SQLiteRepository{db: db}
}

func (r *SQLiteRepository) SaveWrappingRover(rover WrappingCollidingRover) error {
	roverAsBytes, err := json.Marshal(MapToPersistenceRover(rover))
	if err != nil {
		return err
	}
	roverAsString := string(roverAsBytes)

	_, err = r.db.Exec("INSERT INTO "+WrappingRoversTable+" (rover) VALUES (?)",
		roverAsString)
	return err
}

func (r *SQLiteRepository) SaveRockyPlanet(planet RockyPlanet) error {
	planetAsBytes, err := json.Marshal(MapToPersistenceRockyPlanet(planet))
	if err != nil {
		return err
	}
	planetAsString := string(planetAsBytes)

	_, err = r.db.Exec("INSERT INTO "+RockyPlanetsTable+" (planet) VALUES (?)",
		planetAsString)
	return err
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
	CREATE TABLE IF NOT EXISTS ` + WrappingRoversTable + ` (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		rover TEXT NOT NULL
	);
	CREATE TABLE IF NOT EXISTS ` + RockyPlanetsTable + ` (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		planet TEXT NOT NULL
	);
`)
	if err != nil {
		log.Fatal(err)
	}
}
