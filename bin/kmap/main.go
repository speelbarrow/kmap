package main

import (
	"fmt"
	"os"

	"github.com/noah-friedman/kmap"
)

func main() {
	fmt.Println("kmap v1.0.2")
	fmt.Println("Created by Noah Friedman")
	fmt.Println()

	c, e := kmap.Program(os.Stdin, os.Stdout)

	if e != nil {
		fmt.Printf("ERROR: %s", e.Error())
	}

	os.Exit(c)
}
