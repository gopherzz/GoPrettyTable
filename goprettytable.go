package goprettytable

type Table struct {
	Fields    [][]string // Fields of table
	Delimeter rune       // Column and Row Delimeter symbol of table
}

// Returns New Table Object
func NewTable(del rune) (t *Table) {
	return &Table{Delimeter: del}
}

// Add Field to Table
func (t *Table) AddField(el []string) {
	t.Fields = append(t.Fields, el)
}

// Print Pretty Table From Your Data
func (t *Table) Print() {
	// FIXME: Fix print more then 3 columns
	for row := 0; row < len(t.Fields); row++ {
		t.printDelimeter()
		for col := 0; col < len(t.Fields[row]); col++ {
			print(string(t.Delimeter), tab(), t.Fields[row][col], padding(t.Fields[row][col]))
			if col == len(t.Fields[row])-1 {
				print(string(t.Delimeter))
			}
		}
		print("\n")
	}
	t.printDelimeter()

}

func (t *Table) printDelimeter() {
	for i := 0; i < int(getMaxDelimeterLength(t.Fields)); i++ {
		print(string(t.Delimeter))
	}
	print("\n")
}

// Get Length of concatenated strings from array + tab width
func stringsLength(strs []string) uint {

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

// Idk its prtints 8 spaces, like tab but no
func tab() string {
	return "        "
}

// Return longest field from array of fields
func getMaxDelimeterLength(source [][]string) uint {
	// FIXME: Fix find length width text in cells less then 4
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
