package addressgen

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"time"
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

func getZips() []map[string]string {
	var zips []map[string]string
	f, err := os.Open("./addressgen/uszips.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	zipsListed := csvParser(f)
	for _, element := range zipsListed {
		cityStateZip := map[string]string{
			"zip":   element[0],
			"city":  element[3],
			"state": element[3]}
		zips = append(zips, cityStateZip)
	}
	return zips
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

func startAdressEngine() AddressEngine {
	//Going to load nessasry datapoints into memory
	streetNames := getStreet()
	zipCodes := getZips()
	ae := AddressEngine{
		streetAdd:          streetNames,
		streetAddLength:    len(streetNames),
		cityStateZip:       zipCodes,
		cityStateZipLength: len(zipCodes),
		seed:               time.Now().Unix(),
	}
	return ae

}

func AddressEngineCreate() {
	startAdressEngine()
	fmt.Println("Placeholder so go doesn't keep getting rid of FMT, FML")
}
