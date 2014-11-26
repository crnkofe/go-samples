package main

import (
	"fmt"
	"github.com/docopt/docopt-go"
	"github.com/go-samples/cassandra"
	"github.com/go-samples/files"
)

func main() {

	usage := `GO Samples Smoke test.

	Usage:
  	  smoke --input <filename>
  	  smoke -h | --help
  	  smoke --version

	Options:
  	  -h --help     Show this screen.
  	  --version     Show version.
  	  --input		<FILENAME>`

	arguments, _ := docopt.Parse(usage, nil, true, "GO Samples Smoke test", false)
	filename, ok := arguments["--input"]
	if ok {
		read_output.Cat(fmt.Sprint(filename))
	} else {
		fmt.Println("No file passed as --input")
	}

	rows, err := cassandra.ReadRows("settings", "general_settings")
	if err != nil {
		fmt.Println("Error while reading Cassandra row.")
	}

	for _, row := range rows {
		fmt.Println("{")
		for k, v := range row {
			fmt.Printf("%s: %s\n", k, v)
		}
		fmt.Println("}")
	}
}
