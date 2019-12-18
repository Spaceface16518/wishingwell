package db_init

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"time"
	"wishingwell/internal/db"
	"wishingwell/internal/db/wish_client"
)

const TIMEOUT = 20 * time.Second

var connString string

func init() {
	connString = os.Getenv("MONGODB_URI")
	if connString == "" {
		panic("Must provide a MONGODB_URI environmental variable")
	}
}

func EnsureInit() error {
	if db.DB == nil {
		c, err := mongo.NewClient(options.Client().ApplyURI(connString))
		if err != nil {
			panic(err)
		}
		err = c.Connect(context.TODO())
		if err != nil {
			return err
		}

		wishClient := wish_client.NewWishClient(c.Database("wishlist").Collection("wishes"), func(parent context.Context) (ctx context.Context, cancelFunc context.CancelFunc) {
			return context.WithTimeout(parent, 30*time.Second)
		})
		db.DB = db.NewDatabase(wishClient)
	}
	return nil
}
