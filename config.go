package db-compare

import (
	"encoding/json"
	"errors"
)

type JsonConfig struct {
	TestsConfig struct {
		QdbBenchmark string   `json:"qdb-benchmark"`
		Databases    []string `json:"databases"`
		Tests        []struct {
			Name     string   `json:"name"`
			Subtests []string `json:"subtests"`
		} `json:"tests"`
		Nodes          []int    `json:"nodes"`
		Threads        []int    `json:"threads"`
		Sizes          []string `json:"sizes"`
		NumberElements []string `json:"number-elements"`
		Pause          string   `json:"pause"`
		Duration       string   `json:"duration"`
	} `json:"tests_config"`
}

func ReadConfig(config []byte) (JsonConfig, error) {
	var jsonConfig JsonConfig
	err := json.Unmarshal(config, &jsonConfig)
	return jsonConfig, err
}

func CheckConfig(jsonConfig JsonConfig) error {
	for database := range jsonConfig.TestsConfig.Databases {
		found := false
		for supported := range supportedDatabases() {
			if database == supported {
				found = true
				break
			}
		}
		if found == false {
			return errors.New("Database is not supported, aborting.\nTo see a list supported databases run: ./db-compare --list-databases")
		}
	}
	return nil
}

func supportedDatabases() []string {
	return []string{"qdb"}
}
