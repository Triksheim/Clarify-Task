# Clarify-Task - Sensor Uploader

Go program that reads sensor data from a log file, processes it, and optionally uploads the data to Clarify using their Go SDK.

---

## Approach

- Read config from `config.yaml`
    - Loaded into Go `Config` struct
- Read sensor data from a log file (one reading per line)   
    - `<timestamp>` ; `<sensor_id>` ; `<measurement_value>` ; `<unit>`
- Parse and normalize:
    - Timestamps are converted to UTC
    - Units are normalized (°F → °C, psi → bar, gpm → l/s)
    - Grouped by sensor ID into Go `Reading` structs
- Print readings to console (if enabled)
  - Output is grouped and sorted by sensor ID
- Upload to Clarify using the Go SDK (if enabled)
    - Loads Clarify credentials from JSON
    - Converts readings to Clarify `DataFrame`
- Error handling
    - Detects malformed lines and parse errors
    - Catches file read issues
    - Handles failed requests with timeout (set in config)
- Logging
    - Errors are always logged to console
    - Errors are logged to file if enabled
    - Request results are logged to file if enabled


---

## How to Use

1. **Edit `config.yaml`**  
   Set paths, flags, and Clarify credentials.

2. **Run the program**  
   ```bash
   go run .