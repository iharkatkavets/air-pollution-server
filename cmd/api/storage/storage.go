package storage

import (
	_ "github.com/mattn/go-sqlite3"
	"sensor/cmd/api/models"
)

type Storage interface {
	CreateMeasurement(m *models.Measurement) error
	GetAllMeasurements(filters map[string]string) ([]models.Measurement, error)
}
