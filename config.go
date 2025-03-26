package main

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Paths struct {
		SensorData         string `yaml:"sensor_data"`
		ClarifyCredentials string `yaml:"clarify_credentials"`
		ErrorLog           string `yaml:"error_log"`
		UploadLog          string `yaml:"upload_log"`
	} `yaml:"paths"`

	Flags struct {
		PrintReadings bool `yaml:"print_readings"`
		PostReadings  bool `yaml:"post_readings"`
		LogError      bool `yaml:"log_error"`
		LogUpload     bool `yaml:"log_upload"`
	} `yaml:"flags"`

	Net struct {
		TimeoutSeconds int `yaml:"timeout_seconds"`
	} `yaml:"net"`
}

func LoadConfig(path string) (Config, error) {
	var cfg Config
	data, err := os.ReadFile(path)
	if err != nil {
		return cfg, err
	}
	err = yaml.Unmarshal(data, &cfg)
	return cfg, err
}
