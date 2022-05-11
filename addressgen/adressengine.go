package addressgen

import (
	"fmt"
	"math/rand"
)

type AddressEngine struct {
	streetAdd          []string
	streetAddLength    int
	cityStateZip       []map[string]string
	cityStateZipLength int
	seed               int64
}

func (ad *AddressEngine) GenerateAddress() {
	rand.Seed(ad.seed)
	streetAddressIndex := rand.Intn(ad.streetAddLength)
	cityStateZipIndex := rand.Intn(ad.cityStateZipLength)
	finalStreet := ad.streetAdd[streetAddressIndex]
	finalCityStateZipIndex := ad.cityStateZip[cityStateZipIndex]
	houseNumber := rand.Intn(9999)
	fmt.Println(houseNumber, finalStreet, finalCityStateZipIndex["city"], finalCityStateZipIndex["state"], finalCityStateZipIndex["zip"])

	ad.seed = ad.seed + 10

}
