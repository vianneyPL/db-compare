package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

// JSONConfig : a json type describing the config file
type JSONConfig struct {
	QdbBenchmark string   `json:"qdb-benchmark"`
	Databases    []string `json:"databases"`
	Tests        []string `json:"tests"`
	TestsConfig  struct {
		Threads  []int    `json:"threads"`
		Sizes    []string `json:"sizes"`
		Packs    []string `json:"packs"`
		Pause    string   `json:"pause"`
		Duration string   `json:"duration"`
	} `json:"tests-config"`
	Clusters struct {
		Servers []struct {
			Location string `json:"location"`
			System   string `json:"system"`
			Nodes    []int  `json:"nodes"`
			Threads  []int  `json:"threads"`
		} `json:"servers"`
		Clients []struct {
			Location string `json:"location"`
			System   string `json:"system"`
			Nodes    []int  `json:"nodes"`
			Threads  []int  `json:"threads"`
		} `json:"clients"`
	} `json:"clusters"`
	Transient bool `json:"transient"`
}

// MustReadConfig : Open and return the bytes of the config file
//
// Panic on error
func MustReadConfig(path string) []byte {
	config, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return config
}

// MustConvertConfig : Transform the config file into json for later use
//
// Panic on error
func MustConvertConfig(fileConfig []byte) JSONConfig {
	var jsonConfig JSONConfig
	err := json.Unmarshal(fileConfig, &jsonConfig)
	if err != nil {
		panic(err)
	}
	return jsonConfig
}

// CheckConfig : Check the configuration for any error
func CheckConfig(jsonConfig JSONConfig) error {
	for database := range jsonConfig.Databases {
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
	for _, test := range jsonConfig.Tests {
		found := false
		for _, supportedTest := range supportedTests(jsonConfig.Databases) {
			if supportedTest == test {
				found = true
			}
		}
		if found == false {
			return errors.New("Test is not supported in this configuration, aborting.\nTo see a list supported tests run: ./db-compare -test_list")
		}
	}
	return nil
}

// supportedDatabases : return a list of supported databases
func supportedDatabases() []string {
	return []string{"qdb"}
}

var supportedTestsStrings []string

// supportedTests : return a list of supported tests
// ask the qdb-benchmark tool
// erase any test that are not present in every databases tested
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
