package read_output

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func ReadOutput(file string) {
	// open input file
	fi, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	scanner := bufio.NewScanner(fi)
	for scanner.Scan() {
		fmt.Println("_")
		fmt.Println(scanner.Text())
	}

	// close fi on exit and check for its returned error
	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()
}
