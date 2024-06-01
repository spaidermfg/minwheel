package main

import (
	"github.com/tealeg/xlsx/v3"
	"log"
)

func main() {
	create()
}

func create() {
	file := xlsx.NewFile()

	sheet, err := file.AddSheet("Sheet1")
	if err != nil {
		log.Fatal(err)
	}

	row := sheet.AddRow()
	row.AddCell().SetString("in")
	row.AddCell().SetString("code")
	row.AddCell().SetString("we")
	row.AddCell().SetString("trust")

	if err = file.Save("./a.xlsx"); err != nil {
		log.Fatal(err)
	}
}

func reader() {

}
