package main

import "fmt"

type SQL struct {
	tables map[string]Table
}

type Table [][]string

func Constructor(names []string, columns []int) SQL {
	tables := make(map[string]Table, len(names))

	for _, name := range names {
		tables[name] = make([][]string, 0)
	}
	return SQL{
		tables,
	}
}

func (this *SQL) InsertRow(name string, row []string) {
	this.tables[name] = append(this.tables[name], row)
}

func (this *SQL) DeleteRow(name string, rowId int) {
	// this.tables[name][rowId] = []{};
}

func (this *SQL) SelectCell(name string, rowId int, columnId int) string {
	return this.tables[name][rowId-1][columnId-1]
}

func main() {
	tables := Constructor([]string{"table1"}, []int{2})

	tables.InsertRow("table1", []string{"c1", "c2"})
	//tables.DeleteRow("table1", 1)
	ans := tables.SelectCell("table1", 1, 1)
	fmt.Println(ans)
}

/**
 * Your SQL object will be instantiated and called as such:
 * obj := Constructor(names, columns);
 * obj.InsertRow(name,row);
 * obj.DeleteRow(name,rowId);
 * param_3 := obj.SelectCell(name,rowId,columnId);
 */
