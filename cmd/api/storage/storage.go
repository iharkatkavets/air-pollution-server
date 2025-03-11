package storage

import (
	_ "github.com/mattn/go-sqlite3"
	"sensor/cmd/api/models"
)

type Storage interface {
	CreateMeasurement(m *models.Measurement) error
	GetAllMeasurements() ([]models.Measurement, error)
}
