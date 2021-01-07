package launchpadmanagerdb

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"os"
)

const (
	psqlUser     = "POSTGRES_USER"
	psqlPassword = "POSTGRES_PASSWORD"
	psqlHost     = "POSTGRES_HOST"
	psqlSchema   = "POSTGRES_DB"
	psqlPort     = "POSTGRES_PORT"
)

var (
	Client   *sqlx.DB
	user     = os.Getenv(psqlUser)
	password = os.Getenv(psqlPassword)
	host     = os.Getenv(psqlHost)
	schema   = os.Getenv(psqlSchema)
	port     = os.Getenv(psqlPort)
)

func init() {
	datasourceName := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		user,
		password,
		host,
		port,
		schema,
	)

	var err error
	Client, err = sqlx.Connect("postgres", datasourceName)
	if err != nil {
		panic(err)
	}

	if err = Client.Ping(); err != nil {
		panic(err)
	}

	m, err := migrate.New(
		"file://./db/migrations",
		datasourceName,
	)
	if err != nil {
		panic(err)
	}
	if err = m.Up(); err != nil {
		fmt.Println("Migrations", err)
	}
}
