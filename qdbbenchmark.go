package main

import (
	"errors"
	"io"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func LoadQdbBenchmark(path string) error {
	_, err := os.Stat("./tool/bin/qdb-benchmark")
	if err != nil {
		err = exec.Command("mkdir", "-p", "tool").Run()
		if err != nil {
			return errors.New("Could not create tool folder")
		}
		if path == "" {
			err := installQdbBenchmark()
			if err != nil {
				return err
			}
		} else {
			exec.Command("cp", path, "./tool/qdb-benchark").Run()
		}
	}
	return nil
}

func installQdbBenchmark() error {
	err := downloadQdbBenchmark()
	if err != nil {
		return errors.New("Could not download qdb-benchmark")
	}
	err = extractQdbBenchmark()
	if err != nil {
		return errors.New("Could not extract qdb-benchmark")
	}
	return nil
}

func extractQdbBenchmark() error {
	cmd := exec.Command("tar", "zxvf", "./qdb-benchmark.tar.gz")
	cmd.Dir = "./tool"
	err := cmd.Run()
	return err
}

func downloadQdbBenchmark() error {
	out, err := os.Create("./tool/qdb-benchmark.tar.gz")
	if err != nil {
		return err
	}
	defer out.Close()

	resp, err := http.Get("https://download.quasardb.net/quasardb/nightly/bench/qdb-benchmark-2.0.0-Linux.tar.gz")
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	_, err = io.Copy(out, resp.Body)
	return err
}

func getTestsFromInput(lines []string, databases []string) map[string][]string {
	isTest := false
	databaseTests := make(map[string][]string)
	for _, database := range databases {
		var tests []string
		toMatch := "\\s*"
		toMatch += database
		toMatch += "_([^\\s]*).*"
		re, err := regexp.Compile(toMatch)
		if err != nil {
			return nil
		}
		for _, line := range lines {
			if line == "Available tests:" {
				isTest = true
				continue
			}
			if isTest == true {
				matched := re.FindStringSubmatch(line)
				if len(matched) > 0 {
					tests = append(tests, matched[1])
				}
			}
		}
		databaseTests[database] = tests
	}
	return databaseTests
}

func getSupportedTests(databases []string) []string {
	out, err := exec.Command("./tool/bin/qdb-benchmark", "-h").Output()
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(out[:len(out)]), "\n")

	databaseTests := getTestsFromInput(lines, databases)

	// suppress tests that are not in every database tested
	var results []string
	for index, database := range databases {
		tests := databaseTests[database]
		if index == 0 {
			results = append(results, tests...)
		} else {
			for index, result := range results {
				found := false
				for _, test := range tests {
					if test == result {
						found = true
						break
					}
				}
				if found == false {
					results = append(results[:index], results[:index+1]...)
				}
			}
		}
	}
	return results
}
