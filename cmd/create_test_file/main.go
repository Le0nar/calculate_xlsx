package main

import (
	"log"
	"strconv"

	"github.com/xuri/excelize/v2"
)

func main() {
	// Создаем новый файл Excel
	f := excelize.NewFile()

	// Явно задаем имя листа, так как по умолчанию оно может быть пустым
	sheet := "Sheet1" // Название листа, можно выбрать любое

	// Заполняем 100000 строк значением 1 в столбце A
	for i := 1; i <= 250000; i++ {
		cellA := "A" + strconv.Itoa(i)
		err := f.SetCellValue(sheet, cellA, 1)
		if err != nil {
			log.Fatal(err)
		}

		cellB := "B" + strconv.Itoa(i)
		err = f.SetCellValue(sheet, cellB, 1)
		if err != nil {
			log.Fatal(err)
		}

		cellC := "C" + strconv.Itoa(i)
		err = f.SetCellValue(sheet, cellC, 1)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Сохраняем файл
	err := f.SaveAs("test_file.xlsx")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Файл успешно сохранен!")
}
