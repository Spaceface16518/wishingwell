package db

import (
	"go.mongodb.org/mongo-driver/bson"
	"wishingwell/internal/db/helpers"
	"wishingwell/internal/db/wish"
)

var DB *Database

type Database struct {
	client helpers.Database
}

//NewDatabse will construct a database container from an database connection. The argument must be initialized before
// passing it to this function because it will become unavailable once it is passed.
func NewDatabase(databaseHelper helpers.Database) *Database {
	return &Database{databaseHelper}
}

func (db *Database) InsertOne(wish wish.Wish) error {
	return db.client.InsertOne(wish)
}

func (db *Database) InsertMany(wish []wish.Wish) error {
	return db.client.InsertMany(wish)
}

func (db *Database) Find(query bson.M) ([]wish.Wish, error) {
	return db.client.Find(query)
}

//GetInternal provides direct access to the internal database helper associated with this Database struct.
func (db *Database) GetInternal() *helpers.Database {
	return &db.client
}
