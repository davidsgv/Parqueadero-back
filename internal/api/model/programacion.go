package model

import "time"

type Programacion struct {
	Id                 int64       `json:"id"`
	Llegada            time.Time   `json:"llegadat"`
	LlegadaI           int64       `json:"llegada"`
	Salida             time.Time   `json:"salidat"`
	SalidaI            int64       `json:"salida"`
	Estadia            int         `json:"estadia"` //segundos
	Bus                Bus         `json:"bus"`
	ParqueaderoSalida  Parqueadero `json:"parqueaderoSalida"`
	ParqueaderoLlegada Parqueadero `json:"parqueaderoLlegada"`
}

type CreateProgramacion struct {
	Id                   int64     `json:"id" validate:"Required"`
	Llegada              time.Time `json:"llegadat" validate:"Required"`
	LlegadaI             int64     `json:"llegada" validate:"Required"`
	Salida               time.Time `json:"salidat" validate:"Required"`
	SalidaI              int64     `json:"salida" validate:"Required"`
	Estadia              int       `json:"estadia" validate:"Required"` //segundos
	BusId                int64     `json:"bus" validate:"Required"`
	ParqueaderoSalidaId  int64     `json:"parqueaderoSalida" validate:"Required"`
	ParqueaderoLlegadaId int64     `json:"parqueaderoLlegada" validate:"Required"`
}
