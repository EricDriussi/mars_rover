package infra

import (
	"database/sql"
	"encoding/json"
	"log"
	"mars_rover/internal/domain/rover"

	_ "github.com/mattn/go-sqlite3"
)

type SQLiteRepository struct {
	db *sql.DB
}

func InitMem() *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS rovers (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		rover TEXT NOT NULL
	);
`)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func NewSQLite(db *sql.DB) *SQLiteRepository {
	return &SQLiteRepository{db: db}
}

func (r *SQLiteRepository) SaveRover(rover rover.Rover) error {
	roverAsBytes, err := json.Marshal(r.mapToPersistenceRover(rover))
	if err != nil {
		return err
	}
	roverAsString := string(roverAsBytes)

	_, err = r.db.Exec("INSERT INTO rovers (rover) VALUES (?)",
		roverAsString)
	return err
}
