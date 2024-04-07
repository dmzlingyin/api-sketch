package dao

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type Hello struct {
	ID        primitive.ObjectID `bson:"_id"`
	Name      string             `bson:"name"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

type HelloDao interface {
	Create(ctx context.Context, name string) error
}

func NewHelloDao(db *mongo.Database) HelloDao {
	return &helloDao{
		collection: db.Collection("hello"),
	}
}

type helloDao struct {
	collection *mongo.Collection
}

func (d *helloDao) Create(ctx context.Context, name string) error {
	now := time.Now()
	_, err := d.collection.InsertOne(ctx, Hello{
		ID:        primitive.NewObjectID(),
		Name:      name,
		CreatedAt: now,
		UpdatedAt: now,
	})
	return err
}
