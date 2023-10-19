package model

type Parqueadero struct {
	Id        int64     `json:"id"`
	Nombre    string    `json:"nombre"`
	Capacidad int       `json:"capacidad"`
	Longitud  float64   `json:"longitud"`
	Latitud   float64   `json:"latitud"`
	Municipio Municipio `json:"municipio"`
}
