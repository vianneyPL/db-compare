package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
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

func MustOpenConfig(path string) []byte {
	config, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return config
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
			return errors.New("Database is not supported, aborting.\nTo see a list supported databases run: ./db-compare -db_list")
		}
	}
	for _, test := range jsonConfig.TestsConfig.Tests {
		for _, subtest := range test.Subtests {
			fulltest := test.Name + "_" + subtest
			found := false
			for _, supportedTest := range supportedTests(jsonConfig.TestsConfig.Databases) {
				if supportedTest == fulltest {
					found = true
				}
			}
			if found == false {
				return errors.New("Test is not supported in this configuration, aborting.\nTo see a list supported tests run: ./db-compare -test_list")
			}
		}
	}
	return nil
}

func supportedDatabases() []string {
	return []string{"qdb"}
}

var supportedTestsStrings []string

func supportedTests(databases []string) []string {
	if supportedTestsStrings == nil {
		supportedTestsStrings = getSupportedTests(databases)
	}
	return supportedTestsStrings
}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}
