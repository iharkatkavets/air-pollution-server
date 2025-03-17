package storage

import (
	"database/sql"
	"log"
	"sensor/cmd/api/models"
	"strconv"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

type SqlStorage struct {
	db *sql.DB
}

func NewSqlStorage(db *sql.DB) *SqlStorage {
	return &SqlStorage{db: db}
}

func (s *SqlStorage) InitDB() {
	s.createTables()
}

func (s *SqlStorage) createTables() {
	createMeasurementTable := `
    CREATE TABLE IF NOT EXISTS measurement (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        sensor STRING NOT NULL,
        parameter STRING NOT NULL,
        value REAL NOT NULL,
        unit STRING NOT NULL,
        timestamp DATETIME NOT NULL 
    )
    `
	_, err := s.db.Exec(createMeasurementTable)
	if err != nil {
		log.Fatal(err)
	}
}

func (s *SqlStorage) CreateMeasurement(m *models.Measurement) error {
	query := `INSERT INTO measurement 
    (sensor, parameter, value, unit, timestamp) VALUES 
    (?, ?, ?, ?, ?)`
	_, err := s.db.Exec(query, m.Sensor, m.Parameter, m.Value, m.Unit, m.Timestamp)
	return err
}

func (s *SqlStorage) GetAllMeasurements(filters map[string]string) ([]models.Measurement, error) {
	query := "SELECT id, sensor, parameter, value, unit, timestamp FROM measurement"
	var args []interface{}

	var conditions []string
	i := 1
	for key, value := range filters {
		conditions = append(conditions, key+" = $"+strconv.Itoa(i))
		args = append(args, value)
		i++
	}

	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	rows, err := s.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var measurements []models.Measurement
	for rows.Next() {
		var mt models.Measurement
		err := rows.Scan(&mt.ID, &mt.Sensor, &mt.Parameter, &mt.Value, &mt.Unit, &mt.Timestamp)
		if err != nil {
			return nil, err
		}
		measurements = append(measurements, mt)
	}

	return measurements, nil
}
