package model

type Bus struct {
	Id        int64  `json:"id"`
	Placa     string `json:"placa"`
	Capacidad int    `json:"capacidad"`
}
