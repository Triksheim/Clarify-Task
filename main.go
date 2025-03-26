package main

import (
	"fmt"
	"time"
)

type Reading struct {
	Timestamp time.Time `json:"timestamp"`
	Value     float64   `json:"measurement_value"`
	Unit      string    `json:"unit"`
}

func (r Reading) Print() {
	fmt.Printf(
		"[%s] %.2f %s\n",
		r.Timestamp.Format(time.RFC3339),
		r.Value,
		r.Unit,
	)
}

func main() {
	// load data from logfile
	filepath := "data/sensor_data.log"
	var sensorDataLines []string = LoadLinesFromFile(filepath)

	// parse logfile data
	var sensorReadings map[string][]Reading = ParseSensorData(sensorDataLines)

	// print all sensor readings in asc id order
	keysAscending := GetSortedKeys(sensorReadings)
	for _, key := range keysAscending {
		fmt.Printf("\nSensor %s has %d readings:\n", key, len(sensorReadings[key]))
		for _, r := range sensorReadings[key] {
			r.Print()
		}
	}

	// post to clarify with Go SDK
	credsFile := "creds/ClarifyCredentials_SensorData_cvgqt44dbtbc73fap23g.json"
	PostSensorReadingsWithSDK(sensorReadings, credsFile)
}
