package db-compare

import "fmt"
import "flag"

func main() {

	// client := ConnectClient()

	// id, url := CreateSheet(client)

	// fmt.Printf("id: %s\n", id)
	// fmt.Printf("url: %s\n", url)

	var listDatabases = flag.Bool("--list-databases", false, "List the supported databases")
	flag.Parse()

	if *listDatabases == true {
		fmt.Println("Supported databases: ", supportedDatabases())
	}
}
