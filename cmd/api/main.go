package main

import (
	"fmt"

	"github.com/thedatashed/xlsxreader"
)

func main() {
	xl, _ := xlsxreader.OpenFile("./teste.xlsx")

	defer xl.Close()

	for row := range xl.ReadRows(xl.Sheets[0]) {
		fmt.Println("id: " + row.Cells[0].Value)
		fmt.Println("nome: " + row.Cells[1].Value)
		fmt.Println("presença: " + row.Cells[2].Value)

	}

}
