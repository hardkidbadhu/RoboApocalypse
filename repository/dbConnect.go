package repository

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"github.com/RobotApocalypse/configuration"
	"github.com/RobotApocalypse/constants"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

var db *sql.DB

type DBConnect interface {
	ConnectDB() *sql.DB
	Migrate()
}

type dbConnect struct {
	cfg configuration.Config
	log *logrus.Entry
}

func (d dbConnect) ConnectDB() *sql.DB {
	var err error
	port, _ := strconv.Atoi(d.cfg.GetString(constants.DBPort))
	psqlConn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		d.cfg.GetString(constants.DBHost), port, d.cfg.GetString(constants.DBUserName), d.cfg.GetString(constants.DBPassword),
		d.cfg.GetString(constants.DBName))

	d.log.Infoln("psqlConn", psqlConn)
	db, err = sql.Open("postgres", psqlConn)
	if err != nil {
		d.log.Fatalf("error: exec schema migration table - %s", err.Error())
	}

	if err = db.Ping(); err != nil {
		d.log.Fatalf("error: pinging database - %s", err.Error())
	}

	db.SetMaxOpenConns(d.cfg.GetInt(constants.DBMaxOpenConnection))
	db.SetMaxIdleConns(d.cfg.GetInt(constants.DBMaxIdleConnection))
	db.SetConnMaxLifetime(time.Duration(d.cfg.GetInt(constants.DBMaxLifeTime)) * time.Minute)
	return db
}

/**
 * @Description: Migrate updates the schema and table init in database listed in the file in directory  /migrations/mysql
 */
func (d dbConnect) Migrate() {

	d.log.Infoln("starting migration.....")
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		d.log.Fatalf("error: migrate (WithInstance) - %s", err.Error())
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://db/migrations/postgres",
		"postgres", driver)
	if err != nil {
		d.log.Fatalf("error: migrate NewWithDatabaseInstance  - %s", err.Error())
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		d.log.Fatalf("error: migration up - %s", err.Error())
	}

	d.log.Infoln("exiting migration.....")
}

func NewDBConnect(cfg configuration.Config, log *logrus.Entry) DBConnect {
	return &dbConnect{
		cfg: cfg,
		log: log,
	}
}
