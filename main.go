package main

import (
	"fmt"

	"random.name/names/namegen"
)

func main() {
	NE := namegen.NameCreation()
	fmt.Println(NE.CreateName())
}
