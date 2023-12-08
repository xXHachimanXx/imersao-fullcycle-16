package main

import (
	"encoding/csv"
	"os"
	"sort"
	"strconv"
	"strings"
)

type DataMatrix [][]string

func (dm DataMatrix) Len() int {
	return len(dm)
}

func (dm DataMatrix) Swap(i, j int) {
	dm[i], dm[j] = dm[j], dm[i]
}

func (dm DataMatrix) Less(i, j int) bool {
	name1, name2 := dm[i][0], dm[j][0]
	if name1 == name2 {
		age1, err := strconv.Atoi(dm[i][1])
		if err != nil {
			panic(err)
		}

		age2, err := strconv.Atoi(dm[j][1])
		if err != nil {
			panic(err)
		}

		return age1 < age2
	}

	return strings.ToLower(name1) < strings.ToLower(name2)
}

func readCSVFile(sourceFilePath string) ([][]string, []string) {
	csvFile, err := os.Open(sourceFilePath)

	if err != nil {
		panic(err)
	}
	defer csvFile.Close()

	csvReader := csv.NewReader(csvFile)

	header, err := csvReader.Read()
	if err != nil {
		panic(err)
	}

	rows, err := csvReader.ReadAll()
	if err != nil {
		panic(err)
	}

	return rows, header
}

func writeCSVFile(destinyFilePath string, dataMatrix DataMatrix, header []string) {
	csvFile, err := os.Create(destinyFilePath)
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)
	defer writer.Flush()

	if err := writer.Write(header); err != nil {
		panic(err)
	}

	for _, dataMatrixRow := range dataMatrix {
		if err := writer.Write(dataMatrixRow); err != nil {
			panic(err)
		}
	}
}

func main() {
	sourceFilePath, destinyFilePath := os.Args[1], os.Args[2]
	rows, header := readCSVFile(sourceFilePath)

	dataMatrix := DataMatrix(rows)
	sort.Sort(dataMatrix)

	writeCSVFile(destinyFilePath, dataMatrix, header)
}
