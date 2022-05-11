package namegen

import (
	"math/rand"
)

type NameEngine struct {
	firstNames       [][]string
	lastNames        []string
	firstNamesLength int
	lastNamesLength  int
	intialSeed       int
}

func (n *NameEngine) CreateName() [3]string {
	rand.Seed(int64(n.intialSeed))
	n.intialSeed = n.intialSeed + 10
	randomFirstIndex := rand.Intn(n.firstNamesLength)
	randomSecondIndex := rand.Intn(n.lastNamesLength)
	fullName := [3]string{n.firstNames[randomFirstIndex][0], n.lastNames[randomSecondIndex], n.firstNames[randomFirstIndex][1]}
	return fullName
}
