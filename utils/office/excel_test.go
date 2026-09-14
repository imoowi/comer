package office

import (
	"testing"

	"github.com/xuri/excelize/v2"
)

func TestNewExcel(t *testing.T) {
	if f := NewExcel(); f == nil {
		t.Fatal("NewExcel returned nil")
	}
}

func TestAddExcelSheet(t *testing.T) {
	f := excelize.NewFile()
	headers := []string{"Name", "Age"}
	rows := [][]interface{}{{"tom", 18}, {"jerry", 20}}
	if err := AddExcelSheet(f, "Users", headers, rows); err != nil {
		t.Fatal(err)
	}
	cases := map[string]string{
		"A1": "Name",
		"B1": "Age",
		"A2": "tom",
		"B2": "18",
		"A3": "jerry",
		"B3": "20",
	}
	for cell, want := range cases {
		if got, _ := f.GetCellValue("Users", cell); got != want {
			t.Errorf("cell %s = %q, want %q", cell, got, want)
		}
	}
}

func TestExportExcel(t *testing.T) {
	f, err := ExportExcel("Data", []string{"A", "B"}, [][]interface{}{{1, 2}})
	if err != nil {
		t.Fatal(err)
	}
	if f == nil {
		t.Fatal("ExportExcel returned nil file")
	}
	if got, _ := f.GetCellValue("Data", "A1"); got != "A" {
		t.Errorf("A1 = %q, want %q", got, "A")
	}
	if got, _ := f.GetCellValue("Data", "B2"); got != "2" {
		t.Errorf("B2 = %q, want %q", got, "2")
	}
}

func TestGetColumnName(t *testing.T) {
	cases := map[int]string{
		0:   "A",
		25:  "Z",
		26:  "AA",
		27:  "AB",
		52:  "BA",
		701: "ZZ",
	}
	for in, want := range cases {
		if got := string(getColumnName(in, 2)); got != want {
			t.Errorf("getColumnName(%d) = %q, want %q", in, got, want)
		}
	}
}

func TestGetColumnRowName(t *testing.T) {
	if got := getColumnRowName(getColumnName(0, 3), 1); got != "A1" {
		t.Errorf("getColumnRowName(A,1) = %q, want A1", got)
	}
	if got := getColumnRowName(getColumnName(0, 3), 23); got != "A23" {
		t.Errorf("getColumnRowName(A,23) = %q, want A23", got)
	}
	if got := getColumnRowName(getColumnName(1, 3), 5); got != "B5" {
		t.Errorf("getColumnRowName(B,5) = %q, want B5", got)
	}
}
