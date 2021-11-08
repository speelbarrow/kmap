package kmap

import (
	"bufio"
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
		args []int
	)

	// Output the first question
	_, _ = out.WriteString("What is the size of the k-map? (3):\n")

	// Wait for user input, assume 3 if empty
	if s, e := r.ReadString('\n'); e != nil {
		return 2, e
	} else if s = strings.Trim(s, "\n"); s == "" {
		size = 3
	} else if size, e = strconv.Atoi(strings.Trim(s, "\n")); e != nil {
		return 2, e
	}

	// Output the second question
	_, _ = out.WriteString("What are the arguments to the k-map?:\n")

	// Wait for user input and then parse it
	if s, e := r.ReadString('\n'); e != nil {
		return 2, e
	} else {
		// Find the delimiter, if it exists
		var delim string
		if a := regexp.MustCompile(`^[0-9]+([^0-9]+)`).FindStringSubmatch(s); len(a) > 1 {
			delim = a[1]
		}

		// Parse the string
		if args, e = Parse(strings.Trim(s, "\n"), delim); e != nil {
			return 1, e
		}
	}

	// Generate the k-map and format it and output it
	if kmap, e := NewKmap(size, args...); e != nil {
		return 1, e
	} else {
		_, _ = out.WriteString(kmap.Format())
	}

	return 0, nil
}
