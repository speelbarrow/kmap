package kmap

import (
	"fmt"
	"math"
)

// Kmap represents a Karnaugh map, or k-map (https://en.wikipedia.org/wiki/Karnaugh_map). It provides values contained in the k-map it represents as well as properties containing data about the k-map (these properties should NOT be modified). A Kmap should be created using the NewKmap() function.
type Kmap struct {
	Values           [][]bool
	Size, Rows, Cols int
}

// NewKmap is the constructor for the Kmap type.
func NewKmap(size int, args ...int) (*Kmap, error) {
	// Define rows and cols in such a way that at most one variable's value must be changed to correctly reflect the size
	rows, cols := 2, 4
	switch size {
	case 3: // If size is 3, values are already correct
	case 2:
		cols = 2
	case 4:
		rows = 4
	default: // If size is not 2, 3, or 4, return an error
		return nil, fmt.Errorf("invalid size %d: must be 2, 3, or 4", size)
	}

	// Set up the k-map with the correct size, applying any provided arguments
	var vals [][]bool
	if len(args) != 0 {
		max := int(math.Pow(2, float64(size))) // Define the max limit an argument may be

		// Create a flat k-map with true values at indices of arguments
		a := make([]bool, rows*cols)
		for _, v := range args {
			if v >= max {
				return nil, fmt.Errorf("invalid argument %d with size %d: must be between 0 and %.0f", v, size, math.Pow(2, float64(size))-1)
			}
			a[v] = true
		}

		// Un-flatten the k-map and store in vals
		for i, s := cols, rows*cols; i <= s; i += cols {
			vals = append(vals, a[i-cols:i])
		}

		// Swap last two columns and last two rows if applicable
		mintermConvert(&vals)
	} else {
		// Create rows and fill with boolean arrays of proper length
		for i := 0; i < rows; i++ {
			vals = append(vals, make([]bool, cols))
		}
	}

	return &Kmap{vals, size, rows, cols}, nil
}

// Minterms represents a true/false value with each index representing the corresponding minterm in the k-map.
func (kmap *Kmap) Minterms() []bool {
	var r []bool
	for _, v := range mintermConvert(kmap.Values) {
		r = append(r, v...)
	}

	return r
}

// mintermConvert switches k-map values between a proper layout and a layout in order of minterms. v may be [][]bool, in which case a new but updated array will be returned, or *[][]bool, in which case nil will be returned and the original array will be modified.
func mintermConvert(v interface{}) [][]bool {
	// Determine type of v, if [][]bool copy into r and set vals to be a pointer to r, otherwise set vals to be v
	var (
		r    [][]bool
		vals *[][]bool
	)
	switch v.(type) {
	case [][]bool:
		for _, v := range v.([][]bool) {
			r = append(r, v)
		}
		vals = &r
	case *[][]bool:
		vals = v.(*[][]bool)
	default:
		panic(fmt.Errorf("invalid type passed to mintermConvert function"))
	}

	switch rows := len(*vals); true {
	case rows == 4:
		(*vals)[3], (*vals)[2] = (*vals)[2], (*vals)[3]
		fallthrough
	case len((*vals)[0]) == 4:
		for i := 0; i < rows; i++ {
			(*vals)[i][3], (*vals)[i][2] = (*vals)[i][2], (*vals)[i][3]
		}
	}

	return r
}
