package service

import (
	"api/model"
	"api/repository"
	"time"
)

type Service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (servicio *Service) GetMunicipios() ([]model.Municipio, error) {
	//validar datos
	//err := rol.ValidateNombre()

	//crear el registro
	municipios, err := servicio.repo.GetMunicipios()
	if err != nil {
		return nil, err
	}

	return municipios, nil
}

func (servicio *Service) GetParqueaderos() ([]model.Parqueadero, error) {
	//validar datos
	//err := rol.ValidateNombre()

	//crear el registro
	parqueaderos, err := servicio.repo.GetParqueaderos()
	if err != nil {
		return nil, err
	}

	return parqueaderos, nil
}

func (servicio *Service) CreateParqueadero(parqueadero model.Parqueadero) ([]model.Parqueadero, error) {
	err := servicio.repo.CreateParqueadero(parqueadero)
	if err != nil {
		return nil, err
	}
	return servicio.GetParqueaderos()
}

func (servicio *Service) GetBuses() ([]model.Bus, error) {
	//validar datos
	//err := rol.ValidateNombre()

	//crear el registro
	buses, err := servicio.repo.GetBuses()
	if err != nil {
		return nil, err
	}

	return buses, nil
}

func (servicio *Service) CreateBus(bus model.Bus) ([]model.Bus, error) {
	err := servicio.repo.CreateBus(bus)
	if err != nil {
		return nil, err
	}
	return servicio.GetBuses()
}

func (servicio *Service) GetProgramaciones() ([]model.Programacion, error) {
	//validar datos
	//err := rol.ValidateNombre()

	//crear el registro
	programas, err := servicio.repo.GetProgramaciones()
	if err != nil {
		return nil, err
	}

	for index, programa := range programas {
		programas[index].LlegadaI = programa.Llegada.Unix()
		programas[index].SalidaI = programa.Salida.Unix()
	}

	return programas, nil
}

func (servicio *Service) CreateProgramacion(programacion model.CreateProgramacion) ([]model.Programacion, error) {
	programacion.Llegada = time.Unix(programacion.LlegadaI, 0)
	programacion.Salida = time.Unix(programacion.SalidaI, 0)

	err := servicio.repo.CreateProgramacion(programacion)
	if err != nil {
		return nil, err
	}
	return servicio.GetProgramaciones()
}
