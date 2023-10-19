package location

type GPS struct {
	Buses []Bus
}

func (gps *GPS) Register(placa string) {
	for _, bus := range gps.Buses {
		if bus.Placa == placa {
			return
		}
	}

	bus := Bus{
		Placa: placa,
	}
	gps.Buses = append(gps.Buses, bus)
}

func (gps *GPS) Update(placa string, latitud, longitud float64) {
	for i, bus := range gps.Buses {
		if bus.Placa == placa {
			gps.Buses[i].Latitud = latitud
			gps.Buses[i].Longitud = longitud
			return
		}
	}
	gps.Register(placa)
}

func (gps *GPS) GetBuses() []Bus {
	return gps.Buses
}
