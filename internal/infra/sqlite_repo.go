package infra

import (
	"database/sql"
	"encoding/json"
	"log"
	"mars_rover/internal/domain/coordinate"
	coord "mars_rover/internal/domain/coordinate"
	"mars_rover/internal/domain/location"
	"mars_rover/internal/domain/location/direction"
	"mars_rover/internal/domain/obstacle"
	planetTest "mars_rover/internal/domain/planet/test"
	"mars_rover/internal/domain/rover"
	"mars_rover/internal/domain/size"

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
	CREATE TABLE IF NOT EXISTS planets (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		width INTEGER NOT NULL,
		height INTEGER NOT NULL,
		obstacles TEXT NOT NULL
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

type RoverPersistenceEntity struct {
	Location  LocationPersistenceEntity  `json:"location"`
	PlanetMap PlanetMapPersistenceEntity `json:"planetMap"`
}

type PlanetMapPersistenceEntity struct {
	Size      size.Size           `json:"size"`
	Obstacles []obstacle.Obstacle `json:"obstacles"`
}

type ObstaclePersistenceEntity struct {
	Coordinate CoordinatePersistenceEntity `json:"coordinate"`
}

type CoordinatePersistenceEntity struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type LocationPersistenceEntity struct {
	CurrentCoord CoordinatePersistenceEntity `json:"currentCoord"`
	FutureCoord  CoordinatePersistenceEntity `json:"futureCoord"`
	Direction    string                      `json:"direction"`
}

func (this ObstaclePersistenceEntity) Occupies(coordinate coord.AbsoluteCoordinate) bool {
	return true
}

func (this ObstaclePersistenceEntity) IsBeyond(size size.Size) bool {
	return true
}

func (r *SQLiteRepository) mapRover(rover rover.Rover) RoverPersistenceEntity {
	return RoverPersistenceEntity{
		Location: LocationPersistenceEntity{
			CurrentCoord: CoordinatePersistenceEntity{
				X: rover.Location().Position().X(),
				Y: rover.Location().Position().Y(),
			},
			FutureCoord: CoordinatePersistenceEntity{
				X: rover.Location().WillBeAt().X(),
				Y: rover.Location().WillBeAt().Y(),
			},
			// TODO: This is a hack. There should probably be a factory or something
			Direction: "North",
		},
		PlanetMap: PlanetMapPersistenceEntity{
			Size: size.Size{
				Width:  rover.Map().Size().Width,
				Height: rover.Map().Size().Height,
			},
			Obstacles: rover.Map().Obstacles(),
		},
	}
}

func directionFromString(dirStr string) direction.Direction {
	switch dirStr {
	case "North":
		return &direction.North{}
	case "South":
		return &direction.South{}
	case "East":
		return &direction.East{}
	case "West":
		return &direction.West{}
	default:
		// Handle the default case or return an error
		// For now, let's return North as default
		return &direction.North{}
	}
}

func ConvertToRover(roverData RoverPersistenceEntity) (rover.Rover, error) {
	// Create a location based on the persisted data
	loc, err := location.From(
		*coordinate.NewAbsolute(roverData.Location.CurrentCoord.X, roverData.Location.CurrentCoord.Y),
		directionFromString(roverData.Location.Direction),
	)
	if err != nil {
		return nil, err
	}

	// Create a planet (for now, assume an empty planet)
	planet := new(planetTest.MockPlanet)
	planet.On("Size").Return(size.Size{})
	planet.On("Obstacles").Return([]obstacle.Obstacle{})

	// Land the rover on the planet
	roverInstance, err := rover.Land(*loc, planet)
	if err != nil {
		return nil, err
	}

	return roverInstance, nil
}

func (r *SQLiteRepository) SaveRover(rover rover.Rover) error {
	roverAsBytes, err := json.Marshal(r.mapRover(rover))
	if err != nil {
		return err
	}
	roverAsString := string(roverAsBytes)

	_, err = r.db.Exec("INSERT INTO rovers (rover) VALUES (?)",
		roverAsString)
	return err
}
