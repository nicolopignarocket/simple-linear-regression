package data

import (
	"os"
	"encoding/csv"
	"io"
	"log"
	"strconv"
)

type DataSet struct {
	Rows []Row
	Columns []string
}

type Row struct {
	Cells map[string]float64
}

func NewDataSet(csvPath string) *DataSet {
	dataSet := &DataSet{}

	file, err := os.Open(csvPath)
	if err != nil {
		log.Fatalln("Error opening csv file:", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	firstLine := true
	numColumns := 0

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalln("Error reading csv line:", err)
		}

		if firstLine {
			firstLine = false
			for i := 0; i < len(record); i++ {
				dataSet.Columns = append(dataSet.Columns, record[i])
			}

			numColumns = len(dataSet.Columns)

			continue
		}

		row := Row{
			Cells: make(map[string]float64, numColumns),
		}

		for i := 0; i < len(record); i++ {
			cellValue, err := strconv.ParseFloat(record[i], 64)
			if err != nil {
				log.Fatalln("Error parsing csv value:", err)
			}
			row.Cells[dataSet.Columns[i]] = cellValue
		}

		dataSet.Rows = append(dataSet.Rows, row)
	}

	return dataSet
}
