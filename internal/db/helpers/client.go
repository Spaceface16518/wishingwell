package helpers

import (
	"go.mongodb.org/mongo-driver/bson"
	"wishingwell/internal/db/wish"
)

type Database interface {
	InsertOne(wish wish.Wish) error
	InsertMany(wishes []wish.Wish) error
	Find(query bson.M) ([]wish.Wish, error)
}
