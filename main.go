package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

type HelloWorld struct {
	Message string `json:"message"`
}

type CSV struct {
	Result []Record `json:"result"`
}

type Record struct {
	No       string
	Type     string
	Size     string
	Quantity string
	NokType  string
	DinType  string
}

func createList(data [][]string) []Record {
	var list []Record

	for i, row := range data {
		// skip the header
		if i > 0 {
			var record Record

			for j, rec := range row {
				switch j {
				case 0:
					record.No = rec
				case 1:
					record.Type = rec
				case 2:
					record.Size = rec
				case 3:
					record.Quantity = rec
				case 4:
					record.NokType = rec
				case 5:
					record.DinType = rec
				}
			}

			list = append(list, record)
		}
	}

	return list
}

func readCsv() []Record {
	f, err := os.Open("sample.csv")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	records := createList(data)

	return records
}

func main() {
	e := echo.New()
	fmt.Println("Hello world!", e)
	fmt.Println("Hello world!")

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, HelloWorld{Message: "hello world!"})
	})

	e.GET("/csv", func(c echo.Context) error {
		records := readCsv()
		return c.JSON(http.StatusOK, CSV{Result: records})
	})

	readCsv()

	// test comment
	e.Logger.Fatal(e.Start(":3030"))
}
