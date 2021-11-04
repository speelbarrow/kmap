package kmap

import "fmt"

// Strings used to format output
var formatStrings = map[int]string{
	2: `        y
  ---------
  | %d | %d |
  ---------
x | %d | %d |
  ---------`,

	3: `              y
  -----------------
  | %d | %d | %d | %d |
  -----------------
x | %d | %d | %d | %d |
  -----------------
          z`,

	4: `              y
  -----------------
  | %d | %d | %d | %d |
  -----------------
  | %d | %d | %d | %d |
  ----------------- x
  | %d | %d | %d | %d |
w -----------------
  | %d | %d | %d | %d |
  -----------------
          z`,
}

// Format creates a formatted string that can be outputted containing the data of the Kmap.
func (kmap *Kmap) Format() string {
	// Create a flattened version of the k-map that uses '0' and '1' values instead of 'true' and 'false'
	var flat []interface{} // interface is used to allow tuple argument in fmt.Sprintf call
	for _, v := range kmap.Values {
		for _, v := range v {
			if v {
				flat = append(flat, 1)
			} else {
				flat = append(flat, 0)
			}
		}
	}

	// Format and return the output
	return fmt.Sprintf(formatStrings[kmap.Size], flat...)
}
