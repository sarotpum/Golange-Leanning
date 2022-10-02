package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // The database driver in use.
)

type Config struct {
	User string
	Password string
	Host string
	Port string
	Name string
	DisableTLS bool
}

func OpenPgSql() (*sqlx.DB, error)  {
	// set env
	var cfg Config
	cfg.Host = `localhost`
	cfg.Name = `menu`
	cfg.User = `postgres`
	cfg.Password = `admin1234`
	cfg.Port = `5432`
	cfg.DisableTLS = true
	sslmode := `disable`

	var dataSource = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name, sslmode)
	return sqlx.Connect("postgres", dataSource)
}