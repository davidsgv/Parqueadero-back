package data

import (
	"database/sql"
	"errors"
	"fmt"
	"parqueadero-back/internal/api/model"
	"time"

	_ "github.com/lib/pq"
)

const (
	host     = "postgres"
	port     = 5432
	user     = "postgres"
	password = "example"
	dbname   = "postgres"
)

type postgresConfig struct {
	ConnMaxLifetime time.Duration //SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	ConnMaxIdleTime time.Duration //SetConnMaxIdleTime sets the maximum amount of time a connection may be idle.
	MaxIdleConns    int           //SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	MaxOpenConns    int           //SetMaxOpenConns sets the maximum number of open connections to the database.
	//URL             string        //connection string BD
}

type PostgresRepository struct {
	db *sql.DB
}

// func NewMysqlRepository(conf MysqlConfig) (repository.IRepositoryEmpresa, error) {
func NewPostgresRepository() (*PostgresRepository, error) {
	conf := postgresConfig{
		ConnMaxLifetime: 2,
		ConnMaxIdleTime: 2,
		MaxIdleConns:    2,
		MaxOpenConns:    2,
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	//defer db.Close()

	db.SetConnMaxLifetime(time.Second * conf.ConnMaxLifetime)
	db.SetConnMaxIdleTime(time.Second * conf.ConnMaxIdleTime)
	db.SetMaxIdleConns(conf.MaxIdleConns)
	db.SetMaxOpenConns(conf.MaxOpenConns)

	pingErr := db.Ping()
	if pingErr != nil {
		panic(pingErr)
	}

	return &PostgresRepository{
		db: db,
	}, nil
}

func (repo *PostgresRepository) GetMunicipios() ([]model.Municipio, error) {
	var query string = `
		SELECT id, nombre 
		FROM municipio
	`

	rows, err := repo.db.Query(query)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

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
	defer rows.Close()
	if err != nil {
		return nil, err
	}

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
		return errors.New("No inserted data")
	}

	return nil
}

func (repo *PostgresRepository) GetBuses() ([]model.Bus, error) {
	var query string = `
		SELECT id, placa, capacidad FROM bus
	`

	rows, err := repo.db.Query(query)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

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
		return errors.New("No inserted data")
	}

	return nil
}

func (repo *PostgresRepository) GetProgramaciones() ([]model.Programacion, error) {
	var query string = `
		SELECT 
			bus.id, bus.placa, bus.capacidad,
			pro.id, pro.llegada, pro.salida, pro.estadia,
			par.id, par.nombre, par.capacidad,par.latitud, par.longitud,
			mun.id, mun.nombre
		FROM bus
		INNER JOIN programacion pro
			ON bus.id = pro.bus_id
		INNER JOIN parqueadero par
			ON pro.parqueadero_id = par.id
		INNER JOIN municipio mun
			ON par.municipio_id = mun.id
	`

	rows, err := repo.db.Query(query)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	programaciones := []model.Programacion{}
	for rows.Next() {
		programa := model.Programacion{
			Bus:         model.Bus{},
			Parqueadero: model.Parqueadero{},
		}

		rows.Scan(&programa.Bus.Id, &programa.Bus.Placa, &programa.Bus.Capacidad,
			&programa.Id, &programa.Llegada, &programa.Salida, &programa.Estadia,
			&programa.Parqueadero.Id, &programa.Parqueadero.Nombre, &programa.Parqueadero.Capacidad, &programa.Parqueadero.Latitud, &programa.Parqueadero.Longitud,
			&programa.Parqueadero.Municipio.Id, &programa.Parqueadero.Municipio.Nombre)

		programaciones = append(programaciones, programa)
	}

	return programaciones, nil
}

func (repo *PostgresRepository) CreateProgramacion(programacion model.CreateProgramacion) error {
	var query string = `
		INSERT INTO programacion (llegada, salida, estadia, bus_id, parqueadero_id)
		VALUES ($1, $2, $3, $4, $5)
	`

	result, err := repo.db.Exec(query, programacion.Llegada, programacion.Salida, programacion.Estadia,
		programacion.BusId, programacion.ParqueaderoId)
	if err != nil {
		return err
	}

	id, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if id < 0 {
		return errors.New("No inserted data")
	}

	return nil
}
