package mongodb

import (
	"context"
	"fmt"
	"net/url"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/ziflex/waitfor/v2"
)

const Scheme = "mongodb"

type Mongo struct {
	uri string
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

	return &Mongo{uri: u.String()}, nil
}

func (m *Mongo) Test(ctx context.Context) error {
	client, err := mongo.NewClient(options.Client().ApplyURI(m.uri))

	if err != nil {
		return err
	}

	return client.Ping(ctx, readpref.Primary())
}
