package namegen

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"time"
)

//This is used to get grab the file, and return the list of list of strings for parseing later on
func parseCSV(f *os.File) [][]string {
	var lines [][]string
	csvReader := csv.NewReader(f)

	for {
		rec, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		lines = append(lines, rec)
	}
	return lines
}

func readCSVLastName() []string {
	var allSurName []string
	f, err := os.Open("./namegen/lastname.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	allLines := parseCSV(f)
	allLines = allLines[1:]
	for index, element := range allLines {
		_ = index
		allSurName = append(allSurName, element[0])
	}
	return allSurName
}

func readCSVFirstName() [][]string {
	var allNames [][]string
	f, err := os.Open("./namegen/firstnames.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	allLines := parseCSV(f)
	//This removes the first line that holds just the data table refrence
	allLines = allLines[1:]
	for index, element := range allLines {
		_ = index
		allNames = append(allNames, []string{element[1], element[3]})
	}
	return allNames
}

//This will create the NameCreation Struct
func NameCreation() NameEngine {
	firstNames := readCSVFirstName()
	lastNames := readCSVLastName()
	nE := NameEngine{
		firstNames:       firstNames,
		lastNames:        lastNames,
		firstNamesLength: len(firstNames),
		lastNamesLength:  len(lastNames),
		intialSeed:       int(time.Now().Unix()),
	}
	return nE

}
