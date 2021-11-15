package kmap

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Program is run by the main function in the executable package. It returns an exit code.
func Program(in, out *os.File) (int, error) {
	// Create variables to hold flag input
	var (
		size                 int
		argsStr, dontCareStr string
		args, dontCare       []int
	)

	flag.IntVar(&size, "s", 0, "Shorthand for -size.")
	flag.IntVar(&size, "size", 0, "The size of the k-map (the number of variables). Valid values are 2, 3, and 4.")

	flag.StringVar(&argsStr, "a", "", "Shorthand for -args.")
	flag.StringVar(&argsStr, "args", "", "The arguments to the k-map.")

	flag.StringVar(&dontCareStr, "dc", "", "Shorthand for -dont-care.")
	flag.StringVar(&dontCareStr, "dont-care", "", "The don't care conditions of the k-map.")

	flag.Parse()

	// Create a reader to read user input
	r := bufio.NewReader(in)

	// Get size input if necessary
	for size == 0 {
		if _, e := out.WriteString("What is the size of the k-map? (3):\n"); e != nil {
			return 2, e
		} else if s, e := r.ReadString('\n'); e != nil {
			return 2, e
		} else if s == "\n" {
			size = 3
		} else if i, e := strconv.Atoi(strings.TrimSuffix(s, "\n")); e != nil || i < 2 || i > 4 {
			if _, e := out.WriteString("Invalid input. Please try again.\n"); e != nil {
				return 2, e
			}
			continue
		} else {
			size = i
		}
	}

	// Get arguments and don't care condition inputs if necessary, making sure any provided input is parsed properly before considering it valid
	for _, strc := range []struct {
		Pointer        *[]int
		String, Prompt string
	}{
		{&args, argsStr, "arguments to"},
		{&dontCare, dontCareStr, "don't care conditions of"},
	} {
		if strc.String == "" {
			if _, e := out.WriteString(fmt.Sprintf("What are the %s the k-map?:\n", strc.Prompt)); e != nil {
				return 2, e
			} else if strc.String, e = r.ReadString('\n'); e != nil {
				return 2, e
			}
		}

		var delim string
		if regex := regexp.MustCompile(`^[0-9]+([^0-9]+)`).FindStringSubmatch(strc.String); len(regex) >= 2 {
			delim = regex[1]
		}

		if p, e := Parse(strings.TrimSuffix(strc.String, "\n"), delim); e != nil {
			return 1, e
		} else {
			*strc.Pointer = p
		}
	}

	// Generate the Kmap object
	if k, e := NewKmap(size, args, dontCare); e != nil {
		return 1, e
	} else {
		// Output the formatted Kmap
		if _, e := out.WriteString(k.Format() + "\n"); e != nil {
			return 2, e
		}
	}

	return 0, nil
}
