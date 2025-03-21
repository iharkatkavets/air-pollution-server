package service

import (
	"sensor/cmd/api/models"
	"sensor/cmd/api/storage"
	"time"
)

type MeasurementService struct {
	storage storage.Storage
}

func NewMeasurementService(s storage.Storage) *MeasurementService {
	return &MeasurementService{storage: s}
}

type CreateMeasurementReq struct {
	MassDensityPM1_0   float64   `json:"mass_density_pm1_0"`
	MassDensityPM2_5   float64   `json:"mass_density_pm2_5"`
	MassDensityPM4_0   float64   `json:"mass_density_pm4_0"`
	MassDensityPM10    float64   `json:"mass_density_pm10"`
	MassDensityUnit    string    `json:"mass_density_unit"`
	ParticleCountPM0_5 float64   `json:"particle_count_pm0_5"`
	ParticleCountPM1_0 float64   `json:"particle_count_pm1_0"`
	ParticleCountPM2_5 float64   `json:"particle_count_pm2_5"`
	ParticleCountPM4_0 float64   `json:"particle_count_pm4_0"`
	ParticleCountPM10  float64   `json:"particle_count_pm10"`
	ParticleCountUnit  string    `json:"particle_count_unit"`
	ParticleSize       float64   `json:"particle_size"`
	ParticleSizeUnit   string    `json:"particle_size_unit"`
	Timestamp          time.Time `json:"timestamp"`
}

func (s *MeasurementService) CreateMeasurement(req *CreateMeasurementReq) error {
	measurements := []models.Measurement{
		{Sensor: models.SensorMassDensity, Parameter: models.ParameterPM1_0, Value: req.MassDensityPM1_0, Unit: req.MassDensityUnit, Timestamp: req.Timestamp},
		{Sensor: models.SensorMassDensity, Parameter: models.ParameterPM2_5, Value: req.MassDensityPM2_5, Unit: req.MassDensityUnit, Timestamp: req.Timestamp},
		{Sensor: models.SensorMassDensity, Parameter: models.ParameterPM4_0, Value: req.MassDensityPM4_0, Unit: req.MassDensityUnit, Timestamp: req.Timestamp},
		{Sensor: models.SensorMassDensity, Parameter: models.ParameterPM10, Value: req.MassDensityPM10, Unit: req.MassDensityUnit, Timestamp: req.Timestamp},

		{Sensor: models.SensorParticleCount, Parameter: models.ParameterPM0_5, Value: req.ParticleCountPM0_5, Unit: req.ParticleCountUnit, Timestamp: req.Timestamp},
		{Sensor: models.SensorParticleCount, Parameter: models.ParameterPM1_0, Value: req.ParticleCountPM1_0, Unit: req.ParticleCountUnit, Timestamp: req.Timestamp},
		{Sensor: models.SensorParticleCount, Parameter: models.ParameterPM2_5, Value: req.ParticleCountPM2_5, Unit: req.ParticleCountUnit, Timestamp: req.Timestamp},
		{Sensor: models.SensorParticleCount, Parameter: models.ParameterPM4_0, Value: req.ParticleCountPM4_0, Unit: req.ParticleCountUnit, Timestamp: req.Timestamp},
		{Sensor: models.SensorParticleCount, Parameter: models.ParameterPM10, Value: req.ParticleCountPM10, Unit: req.ParticleCountUnit, Timestamp: req.Timestamp},

		{Sensor: models.SensorParticleSize, Parameter: models.ParameterSize, Value: req.ParticleSize, Unit: req.ParticleSizeUnit, Timestamp: req.Timestamp},
	}

	for _, measurement := range measurements {
		if err := s.storage.CreateMeasurement(&measurement); err != nil {
			return err
		}
	}
	return nil
}

func (s *MeasurementService) GetAllMeasurements(filters map[string]string) ([]models.Measurement, error) {
	readings, err := s.storage.GetAllMeasurements(filters)
	if err != nil {
		return nil, err
	}
	return readings, nil
}
