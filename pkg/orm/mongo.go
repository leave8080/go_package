package orm

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoConfig struct {
	MongoUri       string
	PoolSize       int
	Database       string
	ConnectTimeOut int
}

func InitMongo(c *MongoConfig) *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(c.ConnectTimeOut)*time.Second)
	defer cancel()
	if c.PoolSize == 0 {
		c.PoolSize = 100
	}
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(c.MongoUri))
	if err != nil {
		panic(err)
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}
	return client
}

func MongoClose(client *mongo.Client, timeout time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	if err := client.Disconnect(ctx); err != nil {
		panic(err)
	}
}
