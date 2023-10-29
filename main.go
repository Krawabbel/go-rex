package main

import (
	"fmt"
	"os"

	"github.com/Krawabbel/go-rex/rex"
)

func main() {
	re := os.Args[1]
	msg := os.Args[2]
	if found, ctx := rex.Match(re, []byte(msg)); found {
		fmt.Printf("[SUCCESS] regular expression \"%s\" found:\n%s\n", re, ctx)
	}
	fmt.Printf("[FAIL] regular expression \"%s\" not found\n", re)
}
