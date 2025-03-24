package main

import (
	"strconv"
	"strings"
	"time"
)

func ParseSensorData(dataLines []string) map[string][]Reading {

	groupedSensorData := make(map[string][]Reading)

	for _, line := range dataLines {

		fields := strings.Split(line, ";")

		timestampString := fixTimestamp(fields[0]) // fix missing(?) colon in timezone
		timestamp, _ := time.Parse(time.RFC3339, timestampString)
		timestamp = timestamp.UTC() // Normalize to UTC time

		sensorID := fields[1]
		value, _ := strconv.ParseFloat(fields[2], 64)
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

func fixTimestamp(time string) string {
	// convert "+0000" → "+00:00"
	if time[len(time)-3] != ':' {
		return time[:len(time)-2] + ":" + time[len(time)-2:]
	}
	return time
}
