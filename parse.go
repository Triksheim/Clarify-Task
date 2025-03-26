package main

import (
	"strconv"
	"strings"
	"time"
)

func ParseSensorData(dataLines []string) map[string][]Reading {

	groupedSensorData := make(map[string][]Reading)

	for i, line := range dataLines {

		fields := strings.Split(line, ";")
		if len(fields) != 4 {
			ErrorLog.Printf("Line %d: invalid format: %s", i+1, line)
			continue
		}

		timestamp, err := time.Parse(time.RFC3339, fields[0])
		if err != nil {
			timestampString := fixTimestamp(fields[0]) // fix missing(?) colon in timezone
			timestamp, err = time.Parse(time.RFC3339, timestampString)
		}
		if err != nil {
			ErrorLog.Printf("Line %d: invalid timestamp: %s", i+1, line)
		}
		timestamp = timestamp.UTC() // Normalize to UTC time

		sensorID := fields[1]

		value, err := strconv.ParseFloat(fields[2], 64)
		if err != nil {
			ErrorLog.Printf("Line %d: invalid value: %s", i+1, line)
		}

		unit := fields[3]

		// Normalize if needed
		switch {
		case strings.HasPrefix(sensorID, "TEMP") && unit == "F":
			value = FahrenheitToCelsius(value)
			unit = "C"
		case strings.HasPrefix(sensorID, "PRESS") && unit == "psi":
			value = PsiToBar(value)
			unit = "bar"
		case strings.HasPrefix(sensorID, "FLOW") && unit == "gpm":
			value = GpmToLitrePerSecond(value)
			unit = "l/s"
		}

		reading := Reading{
			Timestamp: timestamp,
			Value:     value,
			Unit:      unit,
		}

		groupedSensorData[sensorID] = append(groupedSensorData[sensorID], reading)

	}
	return groupedSensorData
}

func fixTimestamp(ts string) string {
	// convert "+0000" → "+00:00"
	if len(ts) < 6 {
		return ts
	}
	if ts[len(ts)-3] != ':' {
		return ts[:len(ts)-2] + ":" + ts[len(ts)-2:]
	}
	return ts
}
