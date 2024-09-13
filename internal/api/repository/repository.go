package repository

import "parqueadero-back/internal/api/model"

type iMunicipio interface {
	GetMunicipios() ([]model.Municipio, error)
}

type iParqueadero interface {
	GetParqueaderos() ([]model.Parqueadero, error)
	CreateParqueadero(model.Parqueadero) error
	// DeleteParqueadero(int64) error
}

type iBus interface {
	GetBuses() ([]model.Bus, error)
	CreateBus(model.Bus) error
	// DeleteParqueadero(int64) error
}

type iprogramacion interface {
	GetProgramaciones() ([]model.Programacion, error)
	CreateProgramacion(model.CreateProgramacion) error
	// DeleteParqueadero(int64) error
}

type Repository interface {
	iMunicipio
	iParqueadero
	iBus
	iprogramacion
}
