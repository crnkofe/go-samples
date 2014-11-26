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

	cols := []string{"key", "column1", "column2", "value"}
	vals := cassandra.Dict{
		"key":     "random",
		"column1": 0,
		"column2": "test2",
		"value":   "Funky test!"}

	err := cassandra.WriteRow("settings", cols, vals)
	if err != nil {
		fmt.Println("Error while inserting Cassandra row.")
	}

	cols = []string{"key", "column1", "column2", "value"}
	s := cassandra.Setting{
		Key:     "random",
		Column1: 0,
		Column2: "test2",
		Value:   cassandra.SettingBlob{Test: "Son of a gun!"}}

	err = cassandra.WriteSetting("settings", s)
	if err != nil {
		fmt.Println("Error while inserting Cassandra row.")
	}

	redSetting, err := cassandra.ReadSetting("settings", "random")
	fmt.Println("Deserialized setting:")
	fmt.Println(redSetting[0].Test)

	rows, err := cassandra.ReadRows("settings", "random")
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
