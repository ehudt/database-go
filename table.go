package main

import "fmt"

// Column type names
const (
	Text  string = "TEXT"
	Int   string = "INT"
	Float string = "FLOAT"
)

type ColumnType struct {
	Name string
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

func Select(t Table, columns []string) (*Table, error) {
	res := &Table{
		Schema: make([]Column, len(columns)),
		ColMap: make(map[string]int, len(columns)),
		Data:   make([][]Cell, len(columns)),
	}
	for i, c := range columns {
		ci, ok := t.ColMap[c]
		if !ok {
			return nil, fmt.Errorf("Column not found: %s", c)
		}
		res.Schema[i] = t.Schema[ci]
		res.ColMap[c] = i
		res.Data[i] = t.Data[ci]
	}
	return res, nil
}
