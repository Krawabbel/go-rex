package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Krawabbel/go-rex/rex"
)

func main() {

	rgx := os.Args[1]

	path := os.Args[2]

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	for line := 0; scanner.Scan(); line++ {

		found, err := rex.Match(rgx, scanner.Bytes())

		if err != nil {
			panic(err)
		}

		if found {
			fmt.Printf("%d: %s\n", line, scanner.Text())
		}
	}
}
