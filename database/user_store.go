package database

import (
	"context"

	"github.com/hraghu25/hotelreservationsystem/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// This is a user collection details
const userCollection = "users"

// We are creating Userstore interface that implement user specific methods.
type UserStore interface {
	GetUserByID(context.Context, string) (*types.User, error)
}

type MongoUserStore struct {
	client   *mongo.Client
	collName *mongo.Collection
}

func NewMongoUserStore(client *mongo.Client) *MongoUserStore {
	return &MongoUserStore{
		client:   client,
		collName: client.Database(dbname).Collection(userCollection),
	}
}

// GetUserByID implements UserStore.
func (s *MongoUserStore) GetUserByID(ctx context.Context, id string) (*types.User, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var usr types.User
	if err := s.collName.FindOne(ctx, bson.M{"_id": oid}).Decode(&usr); err != nil {
		return nil, err
	}
	return &usr, nil
}
