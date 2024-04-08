package dao

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
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
	Query(ctx context.Context, name string) (*Hello, error)
}

func NewHelloDao(db *mongo.Database) HelloDao {
	return &helloDao{
		t: db.Collection("hello"),
	}
}

type helloDao struct {
	t *mongo.Collection
}

func (d *helloDao) Create(ctx context.Context, name string) error {
	now := time.Now()
	_, err := d.t.InsertOne(ctx, Hello{
		ID:        primitive.NewObjectID(),
		Name:      name,
		CreatedAt: now,
		UpdatedAt: now,
	})
	return err
}

func (d *helloDao) Query(ctx context.Context, name string) (*Hello, error) {
	var res Hello
	filter := bson.M{"name": name}
	err := d.t.FindOne(ctx, filter).Decode(&res)
	return &res, err
}
