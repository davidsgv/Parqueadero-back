package location

import "sync"

var lock = &sync.Mutex{}

var singleInstance *Location

func GetInstance() *Location {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()

		if singleInstance == nil {
			singleInstance = &Location{
				locations: make([]Vehicle, 0),
			}
		}
	}

	return singleInstance
}
