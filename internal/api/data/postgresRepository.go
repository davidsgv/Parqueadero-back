package data

import (
	"database/sql"
	"fmt"
	"time"
)

type ConnectionInfo struct {
	Host     string
	Port     string
	User     string
	Password string
	DBname   string
}

type postgresConfig struct {
	ConnMaxLifetime time.Duration //SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	ConnMaxIdleTime time.Duration //SetConnMaxIdleTime sets the maximum amount of time a connection may be idle.
	MaxIdleConns    int           //SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	MaxOpenConns    int           //SetMaxOpenConns sets the maximum number of open connections to the database.
}

type PostgresRepository struct {
	db *sql.DB
}

// func NewMysqlRepository(conf MysqlConfig) (repository.IRepositoryEmpresa, error) {
func NewPostgresRepository(conInfo ConnectionInfo) (*PostgresRepository, error) {
	conf := postgresConfig{
		ConnMaxLifetime: 2,
		ConnMaxIdleTime: 2,
		MaxIdleConns:    2,
		MaxOpenConns:    2,
	}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		conInfo.Host, conInfo.Port, conInfo.User, conInfo.Password, conInfo.DBname)

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
