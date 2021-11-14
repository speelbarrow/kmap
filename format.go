package kmap

import "fmt"

// Strings used to format output
var formatStrings = map[int]string{
	2: `        y
  ---------
  | %c | %c |
  ---------
x | %c | %c |
  ---------`,

	3: `              y
  -----------------
  | %c | %c | %c | %c |
  -----------------
x | %c | %c | %c | %c |
  -----------------
          z`,

	4: `              y
  -----------------
  | %c | %c | %c | %c |
  -----------------
  | %c | %c | %c | %c |
  ----------------- x
  | %c | %c | %c | %c |
w -----------------
  | %c | %c | %c | %c |
  -----------------
          z`,
}

// Format creates a formatted string that can be outputted containing the data of the Kmap.
func (kmap *Kmap) Format() string {
	// Create a flattened version of the k-map that uses '1', 'X' and '0' values instead of &true, &false and nil

	var flat []interface{} // interface is used to allow tuple argument in fmt.Sprintf call
	for _, v := range kmap.Values {
		for _, v := range v {
			if v == nil {
				flat = append(flat, '0')
			} else if *v {
				flat = append(flat, '1')
			} else {
				flat = append(flat, 'X')
			}
		}
	}

	// Format and return the output
	return fmt.Sprintf(formatStrings[kmap.Size], flat...)
}
