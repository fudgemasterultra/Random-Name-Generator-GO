package namegen

import (
	"math/rand"
)

type NameEngine struct {
	firstNames       []string
	lastNames        []string
	firstNamesLength int
	lastNamesLength  int
	intialSeed       int
}

func (n *NameEngine) CreateName() [2]string {
	rand.Seed(int64(n.intialSeed))
	n.intialSeed = n.intialSeed + 10
	randomFirstIndex := rand.Intn(n.firstNamesLength)
	randomSecondIndex := rand.Intn(n.lastNamesLength)
	fullName := [2]string{n.firstNames[randomFirstIndex], n.lastNames[randomSecondIndex]}
	return fullName
}
