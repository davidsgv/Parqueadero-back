package location

type Location struct {
	locations    []Vehicle
	EventHandler Events
}

func (gps *Location) notifyUpdate() {
	if gps.EventHandler != nil {
		gps.EventHandler.OnLocationUpdate()
	}
}

func (gps *Location) Update(id string, v Vehicle) {
	defer gps.notifyUpdate()

	for i, bus := range gps.locations {
		if bus.Plate == v.Plate {
			gps.locations[i].identifier = id
			gps.locations[i].Latitude = v.Latitude
			gps.locations[i].Longitude = v.Longitude
			return
		}
	}

	ve := Vehicle{
		identifier: id,
		Plate:      v.Plate,
		Latitude:   v.Latitude,
		Longitude:  v.Longitude,
	}
	gps.locations = append(gps.locations, ve)
}

func (gps *Location) Remove(id string) {
	defer gps.notifyUpdate()

	for i, bus := range gps.locations {
		if bus.identifier == id {
			gps.locations[i] = gps.locations[len(gps.locations)-1] // Copy last element to index i.
			gps.locations[len(gps.locations)-1] = Vehicle{}        // Erase last element (write zero value).
			gps.locations = gps.locations[:len(gps.locations)-1]   // Truncate slice.
			return
		}
	}
}

func (gps *Location) GetLocations() []Vehicle {
	return gps.locations
}
