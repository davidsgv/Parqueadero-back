package model

type Parqueadero struct {
	Id        int64     `json:"id" validate:"Required"`
	Nombre    string    `json:"nombre" validate:"Required"`
	Capacidad int       `json:"capacidad" validate:"Required,gt=10"`
	Longitud  float64   `json:"longitud" validate:"Required,gte=-90,lte=90"`
	Latitud   float64   `json:"latitud" validate:"Required,gte=-90,lte=90"`
	Municipio Municipio `json:"municipio" validate:"Required"`
}
