package data

import (
	"errors"
	"parqueadero-back/internal/api/model"

	_ "github.com/lib/pq"
)

func (repo *PostgresRepository) GetMunicipios() ([]model.Municipio, error) {
	var query string = `
		SELECT id, nombre 
		FROM municipio
	`

	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	municipios := []model.Municipio{}
	for rows.Next() {
		municipio := model.Municipio{}
		rows.Scan(&municipio.Id, &municipio.Nombre)
		municipios = append(municipios, municipio)
	}

	return municipios, nil
}

func (repo *PostgresRepository) GetParqueaderos() ([]model.Parqueadero, error) {
	var query string = `
		SELECT 
			par.id
			, par.nombre
			, par.capacidad
			, par.latitud
			, par.longitud
			, mun.id
			, mun.nombre
		FROM parqueadero par
		INNER JOIN municipio mun
			ON par.municipio_id = mun.id
	`

	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	parqueaderos := []model.Parqueadero{}
	for rows.Next() {
		parqueadero := model.Parqueadero{
			Municipio: model.Municipio{},
		}

		rows.Scan(&parqueadero.Id, &parqueadero.Nombre, &parqueadero.Capacidad, &parqueadero.Latitud,
			&parqueadero.Longitud, &parqueadero.Municipio.Id, &parqueadero.Municipio.Nombre)

		parqueaderos = append(parqueaderos, parqueadero)
	}

	return parqueaderos, nil
}

func (repo *PostgresRepository) CreateParqueadero(parqueadero model.Parqueadero) error {
	var query string = `
		INSERT INTO parqueadero (nombre, capacidad, latitud, longitud, municipio_id)
		VALUES ($1, $2, $3, $4, $5)
	`

	result, err := repo.db.Exec(query, parqueadero.Nombre, parqueadero.Capacidad, parqueadero.Latitud,
		parqueadero.Longitud, parqueadero.Municipio.Id)
	if err != nil {
		return err
	}

	id, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if id < 0 {
		return errors.New("no inserted data")
	}

	return nil
}

func (repo *PostgresRepository) GetBuses() ([]model.Bus, error) {
	var query string = `
		SELECT id, placa, capacidad FROM bus
	`

	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	buses := []model.Bus{}
	for rows.Next() {
		bus := model.Bus{}

		rows.Scan(&bus.Id, &bus.Placa, &bus.Capacidad)

		buses = append(buses, bus)
	}

	return buses, nil
}

func (repo *PostgresRepository) CreateBus(bus model.Bus) error {
	var query string = `
		INSERT INTO bus (placa, capacidad)
		VALUES ($1, $2)
	`

	result, err := repo.db.Exec(query, bus.Placa, bus.Capacidad)
	if err != nil {
		return err
	}

	id, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if id < 0 {
		return errors.New("no inserted data")
	}

	return nil
}

func (repo *PostgresRepository) GetProgramaciones() ([]model.Programacion, error) {
	var query string = `
		SELECT 
			bus.id,
			bus.placa,
			bus.capacidad,
			pro.id,
			pro.llegada,
			pro.salida,
			pro.estadia,
			parS.id,
			parS.nombre,
			parS.capacidad,
			parS.latitud,
			parS.longitud,
			munS.id,
			munS.nombre,
			parL.id,
			parL.nombre,
			parL.capacidad,
			parL.latitud,
			parL.longitud,
			munL.id,
			munL.nombre
		FROM bus
		INNER JOIN programacion pro
			ON bus.id = pro.bus_id
		INNER JOIN parqueadero parS
			ON pro.parqueadero_salida_id = parS.id
		INNER JOIN municipio munS
			ON parS.municipio_id = munS.id
		INNER JOIN parqueadero parL
			ON pro.parqueadero_llegada_id = parL.id
		INNER JOIN municipio munL
			ON parL.municipio_id = munL.id
	`

	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	programaciones := []model.Programacion{}
	for rows.Next() {
		programa := model.Programacion{
			Bus:                model.Bus{},
			ParqueaderoSalida:  model.Parqueadero{},
			ParqueaderoLlegada: model.Parqueadero{},
		}

		rows.Scan(&programa.Bus.Id, &programa.Bus.Placa, &programa.Bus.Capacidad,
			&programa.Id, &programa.Llegada, &programa.Salida, &programa.Estadia,

			&programa.ParqueaderoSalida.Id, &programa.ParqueaderoSalida.Nombre, &programa.ParqueaderoSalida.Capacidad,
			&programa.ParqueaderoSalida.Latitud, &programa.ParqueaderoSalida.Longitud,
			&programa.ParqueaderoSalida.Municipio.Id, &programa.ParqueaderoSalida.Municipio.Nombre,

			&programa.ParqueaderoLlegada.Id, &programa.ParqueaderoLlegada.Nombre, &programa.ParqueaderoLlegada.Capacidad,
			&programa.ParqueaderoLlegada.Latitud, &programa.ParqueaderoLlegada.Longitud,
			&programa.ParqueaderoLlegada.Municipio.Id, &programa.ParqueaderoLlegada.Municipio.Nombre,
		)

		programaciones = append(programaciones, programa)
	}

	return programaciones, nil
}

func (repo *PostgresRepository) CreateProgramacion(programacion model.CreateProgramacion) error {
	var query string = `
		INSERT INTO programacion (llegada, salida, estadia, bus_id, parqueadero_salida_id, parqueadero_llegada_id)
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	result, err := repo.db.Exec(query, programacion.Llegada, programacion.Salida, programacion.Estadia,
		programacion.BusId, programacion.ParqueaderoSalidaId, programacion.ParqueaderoLlegadaId)
	if err != nil {
		return err
	}

	id, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if id < 0 {
		return errors.New("no inserted data")
	}

	return nil
}
