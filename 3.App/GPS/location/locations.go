package location

import "sync"

var lock = &sync.Mutex{}

type Location struct {
	GPS
}

var singleInstance *Location

func GetInstance() *Location {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()

		if singleInstance == nil {
			singleInstance = &Location{
				GPS: GPS{
					Buses: make([]Bus, 0),
				},
			}
		}
	}

	return singleInstance
}
