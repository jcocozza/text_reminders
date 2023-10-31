package messaging

import (
    "encoding/json"
    "os"
    "log"
)

type Config struct {
	From string `json:"from"`
	Password string `json:"password"`
	To []string `json:"to"`
	SmtpHost string `json:"smtpHost"`
	SmtpPort string `json:"smtpPort"`
}

func ReadConfig() Config {
    // Read the JSON file
    jsonFile, err := os.ReadFile("config.json")
    if err != nil {
        log.Fatalf("Error reading JSON file: %v", err)
    }

    // Create a Passwords struct to unmarshal the JSON data into
    var config Config

    // Unmarshal the JSON data into the Passwords struct
    if err := json.Unmarshal(jsonFile, &config); err != nil {
        log.Fatalf("Error unmarshaling JSON: %v", err)
    }

	return config
}
