package main

import "fmt"
import "flag"
import "os"

func main() {

	// client := ConnectClient()

	// id, url := CreateSheet(client)

	// fmt.Printf("id: %s\n", id)
	// fmt.Printf("url: %s\n", url)

	var config = flag.String("c", "config.json", "Specify a config file")
	var listDatabases = flag.Bool("list-databases", false, "List the supported databases")
	var listTests = flag.Bool("list-tests", false, "List the supported tests")
	flag.Parse()

	fileConfig := MustReadConfig(*config)
	jsonConfig := MustConvertConfig(fileConfig)
	MustLoadQdbBenchmark(jsonConfig.QdbBenchmark)

	if *listDatabases == true {
		fmt.Println("Supported databases:")
		for _, test := range supportedDatabases() {
			fmt.Println("\t", test)
		}
		fmt.Println("")
	}

	if *listTests == true {
		fmt.Println("Supported tests:")
		for _, test := range supportedTests(jsonConfig.Databases) {
			fmt.Println("\t", test)
		}
		fmt.Println("")
	}

	if *listTests == false && *listDatabases == false {
		err := CheckConfig(jsonConfig)
		if err != nil {
			fmt.Println("Error while checking config: ", err)
			os.Exit(-1)
		}
	}

}
