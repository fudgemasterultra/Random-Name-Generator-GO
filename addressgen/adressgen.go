package addressgen

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func csvParser(f *os.File) [][]string {
	var allLines [][]string

	csvReader := csv.NewReader(f)
	for {
		rec, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		allLines = append(allLines, rec)
	}
	allLines = allLines[1:]
	return allLines
}

func getStreet() []string {
	var streets []string
	f, err := os.Open("./addressgen/street_names.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	streetsListed := csvParser(f)

	for _, element := range streetsListed {
		streets = append(streets, element[0])
	}
	return streets

}

func startAdressEngine() {
	//Going to load nessasry datapoints into memory
	streetNames := getStreet()
	_ = streetNames

}

func AddressEngineCreate() {
	startAdressEngine()
	fmt.Println("Placeholder so go doesn't keep getting rid of FMT, FML")
}
