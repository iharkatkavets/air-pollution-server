package models

import "time"

const (
	SensorMassDensity   string = "mass_density"
	SensorParticleCount string = "particle_count"
	SensorParticleSize  string = "particle_size"
	ParameterPM0_5      string = "pm0.5"
	ParameterPM1_0      string = "pm1.0"
	ParameterPM2_5      string = "pm2.5"
	ParameterPM4_0      string = "pm4.0"
	ParameterPM10       string = "pm10"
	ParameterSize       string = "size"
)

type Measurement struct {
	ID        int
	Sensor    string
	Parameter string
	Value     float64
	Unit      string
	Timestamp time.Time
}
