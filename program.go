package kmap

import (
	"bufio"
	"flag"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Program is run by the main function in the executable package. It returns an exit code.
func Program(in, out *os.File) (int, error) {
	// Create a reader to read user input
	r := bufio.NewReader(in)

	// Create variables to hold input
	var (
		size int
		args string
	)

	flag.IntVar(&size, "s", 0, "Shorthand for -size.")
	flag.IntVar(&size, "size", 0, "The size of the k-map (the number of variables). Valid values are 2, 3, and 4.")

	flag.StringVar(&args, "a", "", "Shorthand for -args.")
	flag.StringVar(&args, "args", "", "The arguments to the k-map.")

	flag.Parse()

	// Output the first question if it was not passed as a command line argument
	if size == 0 {
		_, _ = out.WriteString("What is the size of the k-map? (3):\n")

		// Wait for user input, assume 3 if empty
		if s, e := r.ReadString('\n'); e != nil {
			return 2, e
		} else if s = strings.Trim(s, "\n"); s == "" {
			size = 3
		} else if size, e = strconv.Atoi(strings.Trim(s, "\n")); e != nil {
			return 2, e
		}
	}

	// Output the second question if it was not passed as a command line argument
	if args == "" {
		_, _ = out.WriteString("What are the arguments to the k-map?:\n")

		// Wait for user input and then parse it
		var e error
		if args, e = r.ReadString('\n'); e != nil {
			return 2, e
		}
	}

	// Find the delimiter of the args string, if it exists
	var delim string
	if a := regexp.MustCompile(`^[0-9]+([^0-9]+)`).FindStringSubmatch(args); len(a) > 1 {
		delim = a[1]
	}

	// Parse the args string
	if a, e := Parse(strings.Trim(args, "\n"), delim); e != nil {
		return 1, e
	} else {
		// Generate the k-map and format it and output it
		if kmap, e := NewKmap(size, a, nil); e != nil {
			return 1, e
		} else {
			_, _ = out.WriteString(kmap.Format() + "\n")
		}

		return 0, nil
	}
}
