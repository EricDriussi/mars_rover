package persistence

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type SQLiteRepository struct {
	db    *sql.DB
	store string
}

func InitMem() (*sql.DB, *SQLiteRepository) {
	repo := &SQLiteRepository{store: ":memory:"}
	repo.connect()
	createTables(repo.db)
	return repo.db, repo
}

func InitFS() *SQLiteRepository {
	repo := &SQLiteRepository{store: "./rover.db"}
	repo.connect()
	createTables(repo.db)
	return repo
}

func (this *SQLiteRepository) connect() {
	db, err := sql.Open("sqlite3", this.store)
	if err != nil {
		log.Fatal(err)
	}
	this.db = db
}

func createTables(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS ` + RoversTable + ` (
		id TEXT PRIMARY KEY,
		rover TEXT NOT NULL,
		type TEXT NOT NULL,
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
