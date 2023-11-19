package infra

import (
	"database/sql"
	"encoding/json"
	"log"
	. "mars_rover/internal/domain/planet/rockyPlanet"
	. "mars_rover/internal/domain/rover/wrappingCollidingRover"

	_ "github.com/mattn/go-sqlite3"
)

type SQLiteRepository struct {
	db *sql.DB
}

func NewSQLite(db *sql.DB) *SQLiteRepository {
	return &SQLiteRepository{db: db}
}

func (r *SQLiteRepository) SaveWrappingRover(rover WrappingCollidingRover) error {
	roverAsBytes, err := json.Marshal(mapToPersistenceRover(rover))
	if err != nil {
		return err
	}
	roverAsString := string(roverAsBytes)

	_, err = r.db.Exec("INSERT INTO wrapping_rovers (rover) VALUES (?)",
		roverAsString)
	return err
}

func (r *SQLiteRepository) SaveRockyPlanet(planet RockyPlanet) error {
	planetAsBytes, err := json.Marshal(mapToPersistenceRockyPlanet(planet))
	if err != nil {
		return err
	}
	planetAsString := string(planetAsBytes)

	_, err = r.db.Exec("INSERT INTO rocky_planets (planet) VALUES (?)",
		planetAsString)
	return err

}

func InitMem() *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS wrapping_rovers (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		rover TEXT NOT NULL
	);
	CREATE TABLE IF NOT EXISTS rocky_planets (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		planet TEXT NOT NULL
	);
`)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func InitFS() *sql.DB {
	db, err := sql.Open("sqlite3", "./mars_rover.db")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS wrapping_rovers (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		rover TEXT NOT NULL
	);
	CREATE TABLE IF NOT EXISTS rocky_planets (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		planet TEXT NOT NULL
	);
`)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
