package main

import (
	"context"
	"encoding/json"
	"os"

	clarify "github.com/clarify/clarify-go"
	"github.com/clarify/clarify-go/fields"
	"github.com/clarify/clarify-go/views"
)

func PostSensorReadingsWithSDK(sensorReadings map[string][]Reading, credentialsFile string) error {

	creds, err := clarify.CredentialsFromFile(credentialsFile)
	if err != nil {
		return err
	}

	ctx := context.Background()
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
		return err
	}

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	enc.Encode(result)

	return nil
}
