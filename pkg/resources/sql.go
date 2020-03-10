package resources

import (
	"context"
	"database/sql"
	"errors"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

type (
	SQLDriver int

	SQL struct {
		driverName string
		connStr    string
	}
)

const (
	UNKNOWN_DRIVER  SQLDriver = 0
	POSTGRES_DRIVER SQLDriver = 1
	MYSQL_DRIVER    SQLDriver = 2
)

var driverNames = map[SQLDriver]string{
	POSTGRES_DRIVER: "postgres",
	MYSQL_DRIVER:    "mysql",
}

func NewSQL(driver SQLDriver, connStr string) (*SQL, error) {
	driverName, found := driverNames[driver]

	if !found {
		return nil, errors.New("unsupported SQL driver type")
	}

	return &SQL{driverName, connStr}, nil
}

func (s SQL) Test(ctx context.Context) error {
	db, err := sql.Open(s.driverName, s.connStr)

	if err != nil {
		return err
	}

	defer db.Close()

	return db.PingContext(ctx)
}
