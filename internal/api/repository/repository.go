package repository

import "parqueadero-back/internal/api/model"

type Repository interface {
	//municipios
	GetMunicipios() ([]model.Municipio, error)
	//parqueadero
	GetParqueaderos() ([]model.Parqueadero, error)
	CreateParqueadero(model.Parqueadero) error
	//bus
	GetBuses() ([]model.Bus, error)
	CreateBus(model.Bus) error
	//programacion
	GetProgramaciones() ([]model.Programacion, error)
	CreateProgramacion(model.CreateProgramacion) error
}
