package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"log"
)

func main() {
	f := excelize.NewFile()
	sheet, err := f.NewSheet("Sheet1")
	if err != nil {
		log.Fatal(err)
	}

	for i, v := range data() {
		f.SetCellValue("Sheet1", fmt.Sprintf("A%d", i+1), v[0])
		f.SetCellValue("Sheet1", fmt.Sprintf("B%d", i+1), v[1])
		f.SetCellValue("Sheet1", fmt.Sprintf("C%d", i+1), v[2])
	}

	f.SetActiveSheet(sheet)
	if err = f.SaveAs("a.xlsx"); err != nil {
		log.Fatal(err)
	}
}

func data() [][]string {
	return [][]string{
		{"Id", "Name", "Age"},
		{"1", "mark", "33"},
		{"2", "lyn", "221"},
		{"3", "daniel", "899"},
	}
}
