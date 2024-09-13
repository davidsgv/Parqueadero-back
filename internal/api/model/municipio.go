package model

type Municipio struct {
	Id     int64  `json:"id"  validate:"required"`
	Nombre string `json:"nombre"`
}
