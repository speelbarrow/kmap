package kmap

// Kmap represents a Karnaugh map, or k-map (https://en.wikipedia.org/wiki/Karnaugh_map). It provides values contained in the k-map it represents as well as properties containing data about the k-map (these properties should NOT be modified). A Kmap should be created using the NewKmap() function.
type Kmap struct {
	Values           [][]bool
	Size, Rows, Cols int
}

// NewKmap is the constructor for the Kmap type.
func NewKmap(size int, args ...int) (*Kmap, error) {
	return nil, nil
}

// Minterms represents a true/false value with each index representing the corresponding minterm in the k-map.
func (k *Kmap) Minterms() []bool {
	return nil
}
