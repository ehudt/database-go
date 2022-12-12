package main

import "fmt"

// Column type names
const (
	Char    string = "CHAR"
	VarChar string = "VARCHAR"
	Text    string = "TEXT"
	Int     string = "INT"
	Float   string = "FLOAT"
)

type ColumnType struct {
	Name string
	Size *int
}

type Column struct {
	Name string
	Type ColumnType
}

type Cell struct {
	Value interface{}
}

type Table struct {
	Schema []Column
	ColMap map[string]int
	Data   [][]Cell // Column major order
}

func (t *Table) Select(columns []string) (*Table, error) {
	// Find the indices of columns in the table.
	cIndices := make([]int, len(columns))
	for i, c := range columns {
		ci, ok := t.ColMap[c]
		if !ok {
			return nil, fmt.Errorf("Column not found: %s", c)
		}
		cIndices[i] = ci
	}

}
