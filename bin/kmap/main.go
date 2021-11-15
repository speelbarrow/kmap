package main

import (
	"fmt"
	"github.com/noah-friedman/kmap"
	"os"
)

func main() {
	fmt.Println("kmap v1.0.1")
	fmt.Println("Created by Noah Friedman")
	fmt.Println()

	c, e := kmap.Program(os.Stdin, os.Stdout)

	if e != nil {
		fmt.Printf("ERROR: %s", e.Error())
	}

	os.Exit(c)
}
