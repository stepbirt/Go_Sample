package store

import (
	"context"

	"github.com/stepbirt/api/todo"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type MongoDBStore struct {
	*mongo.Collection
}

func NewMongoDBStore(col *mongo.Collection) *MongoDBStore {
	return &MongoDBStore{Collection: col}
}

func (s *MongoDBStore) New(todo *todo.Todo) error {
	_, errr := s.Collection.InsertOne(context.Background(), todo)

	return errr
}
