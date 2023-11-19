package infra

import (
	"database/sql"
	"encoding/json"
	"log"
	. "mars_rover/internal/domain/planet"
	"mars_rover/internal/domain/rover"
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

func (r *SQLiteRepository) SaveRover(rover rover.Rover) error {
	roverAsBytes, err := json.Marshal(MapToPersistenceRover(rover))
	if err != nil {
		return err
	}
	roverAsString := string(roverAsBytes)

	_, err = r.db.Exec("INSERT INTO "+RoversTable+" (rover) VALUES (?)",
		roverAsString)
	return err
}

func (r *SQLiteRepository) SavePlanet(planet Planet) error {
	planetAsBytes, err := json.Marshal(MapToPersistencePlanet(planet))
	if err != nil {
		return err
	}
	planetAsString := string(planetAsBytes)

	_, err = r.db.Exec("INSERT INTO "+PlanetsTable+" (planet) VALUES (?)",
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
	CREATE TABLE IF NOT EXISTS ` + RoversTable + ` (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		rover TEXT NOT NULL
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
