package resources

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoDB struct {
	uri string
}

func NewMongoDB(uri string) Resource {
	return &MongoDB{uri}
}

func (m *MongoDB) Test(ctx context.Context) error {
	client, err := mongo.NewClient(options.Client().ApplyURI(m.uri))

	if err != nil {
		return err
	}

	return client.Ping(ctx, readpref.Primary())
}
