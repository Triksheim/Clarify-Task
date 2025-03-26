package main

import (
	"context"
	"encoding/json"
	"time"

	clarify "github.com/clarify/clarify-go"
	"github.com/clarify/clarify-go/fields"
	"github.com/clarify/clarify-go/views"
)

func PostSensorReadingsWithSDK(sensorReadings map[string][]Reading, cfg Config) error {

	credentialsFile := cfg.Paths.ClarifyCredentials
	creds, err := clarify.CredentialsFromFile(credentialsFile)
	if err != nil {
		ErrorLog.Printf("Loading credentials failed: %v", err)
		return err
	}

	timeout := time.Duration(cfg.Net.TimeoutSeconds) * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	client := creds.Client(ctx)

	// build DataFrame for post
	df := views.DataFrame{}
	for sensorID, readings := range sensorReadings {
		timeseries := map[fields.Timestamp]float64{} // format {timestamp : value}
		for _, r := range readings {
			t := fields.AsTimestamp(r.Timestamp.UTC())
			timeseries[t] = r.Value
		}
		df[sensorID] = timeseries // format df[sensorID]{timestamp : value, timestamp : value ...}
	}

	// post req
	result, err := client.Insert(df).Do(ctx)
	if err != nil {
		ErrorLog.Printf("POST request failed: %v", err)
		return err
	}

	if cfg.Flags.LogUpload {
		resultJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			ErrorLog.Printf("Failed to encode result: %v", err)
		} else {
			UploadLog.Printf("\n%s", resultJSON)
		}

	}

	return nil
}
