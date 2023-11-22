package infra

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type SQLiteRepository struct {
	db *sql.DB
}

func InitMem() (*sql.DB, *SQLiteRepository) {
	db := setup(":memory:")
	return db, &SQLiteRepository{db: db}
}

func InitFS() (*sql.DB, *SQLiteRepository) {
	db := setup("./mars_rover.db")
	return db, &SQLiteRepository{db: db}
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

const (
	RoversTable  = "rovers"
	PlanetsTable = "planets"
)
