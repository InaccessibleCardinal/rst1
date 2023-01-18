package xl

import (
	"archive/zip"
	"fmt"

	"github.com/xuri/excelize/v2"
)

const path = "/Users/kennethlandsbaum/go/test-book.xlsx"

func openFile(f *zip.File) {
	println("open zip file...")
	zippedFile, err := f.Open()
	if err != nil {
		println(err.Error())
	}
	defer zippedFile.Close()
	b := make([]byte, 10000)
	zippedFile.Read(b)
	println("reading file...")
	println(string(b))
}

func UnzipXL() {
	reader, err := zip.OpenReader(path)
	if err != nil {
		fmt.Printf("error opening reader: %s", err.Error())
	}
	defer reader.Close()

	for _, f := range reader.File {
		fmt.Println(f.Name)
		if f.Name == "xl/worksheets/_rels/sheet1.xml.rels" {
			openFile(f)
		}
	}
}

func ReadXL() {
	f, err := excelize.OpenFile(path)
	if err != nil {
		fmt.Printf("error opening file: %s", err.Error())
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Printf("error closing file: %s", err.Error())
		}
	}()

	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Printf("error getting rows: %s", err.Error())
	}
	for _, row := range rows {
		for _, d := range row {
			println(d)
		}
	}
}
