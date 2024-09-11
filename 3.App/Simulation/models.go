package main

type Coordenada struct {
	Latitud  float64
	Longitud float64
}

type Viaje struct {
	Origen  Coordenada
	Destino Coordenada
	Tiempo  int64 //miliseconds
}

type Bus struct {
	Placa string
}

type Programacion struct {
	Bus
	Viaje
	Inicio int64
}

type Message struct {
	Latitud  float64 `json:"latitude"`
	Longitud float64 `json:"longitude"`
	Placa    string  `json:"plate"`
}
