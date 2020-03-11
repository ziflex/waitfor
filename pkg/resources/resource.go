package resources

import (
	"context"
	"fmt"
	"net/url"
	"strings"
)

type Resource interface {
	Test(ctx context.Context) error
}

func New(location string) (Resource, error) {
	u, err := url.Parse(location)

	if err != nil {
		return nil, err
	}

	switch u.Scheme {
	case "http", "https":
		return NewHTTP(location), nil
	case "proc":
		return NewProcess(location), nil
	case "file":
		return NewFile(strings.TrimPrefix(location, "file://")), nil
	case "mongodb":
		return NewMongoDB(location), nil
	case "postgres":
		return NewSQL(POSTGRES_DRIVER, location)
	case "mysql", "mariadb":
		return NewSQL(MYSQL_DRIVER, strings.TrimPrefix(location, "mysql://"))
	default:
		return nil, fmt.Errorf("unsupported location type: %s", u.Scheme)
	}
}
