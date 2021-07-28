package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"net/url"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ziflex/waitfor/v2"
)

const Scheme = "mysql"

type MySQL struct {
	url *url.URL
}

func Use() waitfor.ResourceConfig {
	return waitfor.ResourceConfig{
		Scheme:  Scheme,
		Factory: New,
	}
}

func New(u *url.URL) (waitfor.Resource, error) {
	if u == nil {
		return nil, fmt.Errorf("%q: %w", "url", waitfor.ErrInvalidArgument)
	}

	return &MySQL{u}, nil
}

func (s *MySQL) Test(ctx context.Context) error {
	db, err := sql.Open(s.url.Scheme, strings.TrimPrefix(s.url.String(), Scheme+"://"))

	if err != nil {
		return err
	}

	defer db.Close()

	return db.PingContext(ctx)
}
