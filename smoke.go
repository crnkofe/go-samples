package main

import (
	"github.com/docopt/docopt-go"
	// "github.com/go-samples/files"
)

func main() {

	usage := `GO Samples Smoke test.

	Usage:
  	  smoke -input <filename>
  	  smoke -h | --help
  	  smoke --version

	Options:
  	  -h --help     Show this screen.
  	  --version     Show version.
  	  --input		<FILENAME>`

	arguments, _ := docopt.Parse(usage, nil, true, "GO Samples Smoke test", false)
	_, ok := arguments["input"]
	if ok {
		// read_output.ReadOutput(filename)
	} else {

	}
}
