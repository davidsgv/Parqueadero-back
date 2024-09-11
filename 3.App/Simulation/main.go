package main

import (
	"sync"
)

func main() {
	var wg sync.WaitGroup

	//ride(Programaciones[0])
	for i := range Programaciones {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			ride(Programaciones[i])
		}(i)
	}
	wg.Wait()
}
