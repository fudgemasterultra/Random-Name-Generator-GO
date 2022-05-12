package main

import (
	"sync"

	"random.name/names/addressgen"
	"random.name/names/namegen"
)

func main() {
	var wg sync.WaitGroup
	var ae addressgen.AddressEngine
	var ne namegen.NameEngine
	wg.Add(2)
	go func() {
		ae = addressgen.AddressEngineCreate()
		wg.Done()
	}()
	go func() {
		ne = namegen.NameCreation()
		wg.Done()
	}()
	wg.Wait()
	for i := 0; i < 300; i++ {
		ae.GenerateAddress()
		ne.CreateName()
	}

}
