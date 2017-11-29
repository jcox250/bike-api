package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strings"
)

const configFile = "config/dev.json"

type Config struct {
	Database `json:"database"`
	ApiKey   string `json:"apiKey"`
	Cors     `json:"cors"`
}

type Database struct {
	Type       string `json:"type"`
	ConnString string `json:"connectionString"`
}

type Cors struct {
	AllowedOrigins []string `json:"allowedOrigins"`
	AllowedHeaders []string `json:"allowedHeaders"`
	AllowedMethods []string `json:"allowedMethods"`
}

func Init() Config {
	var c Config
	// Read config file
	bytes, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatal("Can't set system configuration from config file: ", err)
	}

	// Decode the json in the config file into the Database object
	reader := strings.NewReader(string(bytes))
	err = json.NewDecoder(reader).Decode(&c)
	if err != nil {
		log.Fatal("Could not decode system configuration from config file: ", err)
	}
	return c
}
