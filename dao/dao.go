package dao

import (
	"context"
	"github.com/dmzlingyin/utils/config"
	"github.com/dmzlingyin/utils/ioc"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func NewDatabase() *mongo.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.Get("mongo.uri").String()))
	if err != nil {
		panic(err)
	}
	return client.Database(config.Get("mongo.database").String())
}

func init() {
	ioc.Put(NewDatabase, "database")
	ioc.Put(NewHelloDao, "dao.hello")
}
