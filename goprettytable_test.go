package goprettytable

import "testing"

func TestPrint(t *testing.T) {
	table := NewTable('*')
	table.AddField([]string{"One", "Two", "Three"})
	table.AddField([]string{"Four", "Fife", "Six"})
	table.AddField([]string{"Seven", "Eight", "Nine"})
	table.Print()
}
