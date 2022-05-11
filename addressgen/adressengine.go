package addressgen

import "math/rand"

type AddressEngine struct {
	streetAdd          []string
	streetAddLength    int
	cityStateZip       []map[string]string
	cityStateZipLength int
	seed               int64
}

func (ad *AddressEngine) GenerateAddress() {
	rand.Seed(ad.seed)
	ad.seed = ad.seed + 10

}
