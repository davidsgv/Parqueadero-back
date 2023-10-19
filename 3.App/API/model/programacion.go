package model

import "time"

type Programacion struct {
	Id          int64       `json:"id"`
	Llegada     time.Time   `json:"llegadat"`
	LlegadaI    int64       `json:"llegada"`
	Salida      time.Time   `json:"salidat"`
	SalidaI     int64       `json:"salida"`
	Estadia     int         `json:"estadia"` //segundos
	Bus         Bus         `json:"bus"`
	Parqueadero Parqueadero `json:"parqueadero"`
}
