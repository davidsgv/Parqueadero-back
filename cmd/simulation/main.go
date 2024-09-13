package main

import (
	"parqueadero-back/internal/simulation"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	//ride(Programaciones[0])
	for i := range simulation.Programaciones {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			simulation.Ride(simulation.Programaciones[i])
		}(i)
	}
	wg.Wait()
}
