package goprettytable

type Table struct {
	Fields    [][]string
	Delimeter rune
}

func NewTable(del rune) (t *Table) {
	return &Table{Delimeter: del}
}

func (t *Table) AddField(el []string) {
	t.Fields = append(t.Fields, el)
}

func (t *Table) Print() {
	// Print Table:
	// TODO: Remove delimeter in start of table
	//
	// -----------
	// 1    2    3
	// -----------
	// 4    5    6
	// -----------
	// 7    8    9
	// -----------

	for row := 0; row < len(t.Fields); row++ {
		t.PrintDelimeter()
		for col := 0; col < len(t.Fields[row]); col++ {
			print(string(t.Delimeter), tab(), t.Fields[row][col], padding(t.Fields[row][col]))
			if col == len(t.Fields[row])-1 {
				print(string(t.Delimeter))
			}
		}
		print("\n")
	}
	t.PrintDelimeter()

}

func (t *Table) PrintDelimeter() {
	for i := 0; i < int(getMaxDelimeterLength(t.Fields)); i++ {
		print(string(t.Delimeter))
	}
	print("\n")
}

func stringsLength(strs []string) uint {
	// Get Length of concatenated strings from array + tab width
	var res uint
	if len(strs) == 0 {
		return 0
	}
	for i, s := range strs {
		res += uint(len(s))
		if i != len(strs)-1 {
			res += uint(8 - len(s))
			res += uint(16)
		}
	}
	return res
}

func padding(s string) string {
	var pad string
	padding := int(uint(8 - len(s)))
	padding += 4
	for i := 0; i < padding; i++ {
		pad += " "
	}
	return pad
}

func tab() string {
	// Idk its prtints 8 spaces, like tab but no
	return "        "
}

func getMaxDelimeterLength(source [][]string) uint {
	// Return longest field from array of fields
	// fmt.Println(source)
	res := make([]uint, 1)
	for _, strs := range source {
		res = append(res, stringsLength(strs))
	}
	var max uint = 0
	for _, el := range res {
		if el > max {
			max = el
		}
	}
	return max + uint(len(source[0][len(source[0])-1])) + (11 - uint(len(source[0][len(source[0])-1])))
}
