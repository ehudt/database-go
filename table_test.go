package main

import "testing"

func TestSelectBasic(t *testing.T) {
	table := Table{
		Schema: []Column{
			{Name: "id", Type: ColumnType{Int}},
			{Name: "firstName", Type: ColumnType{Text}},
			{Name: "lastName", Type: ColumnType{Text}},
			{Name: "age", Type: ColumnType{Float}},
		},
		ColMap: map[string]int{
			"id":        0,
			"firstName": 1,
			"lastName":  2,
			"age":       3,
		},
		Data: [][]Cell{
			{{1}, {2}, {3}, {4}},
			{{"Ehud"}, {"Neta"}, {"Shiri"}, {"Omer"}},
			{{"Tamir"}, {"Perri Tamir"}, {"Tamir"}, {"Tamir"}},
			{{37.2}, {35.1}, {1.7}, {1.7}},
		},
	}

	columns := []string{"firstName", "age"}
	res, err := Select(table, columns)
	if err != nil {
		t.Fatalf("Select failed: %v", err)
	}
	for i, c := range columns {
		if res.Schema[i].Name != c {
			t.Errorf("Result column %d mismatch. got %s, want %s", i, res.Schema[i].Name, c)
		}
	}
}
