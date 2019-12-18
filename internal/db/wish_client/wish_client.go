package wish_client

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"wishingwell/internal/db/wish"
)

type WishClient struct {
	*mongo.Collection
	contextFunc func(parent context.Context) (context.Context, context.CancelFunc)
}

func NewWishClient(collection *mongo.Collection, contextFunc func(parent context.Context) (context.Context, context.CancelFunc)) *WishClient {
	return &WishClient{Collection: collection, contextFunc: contextFunc}
}

func (client *WishClient) InsertOne(wish wish.Wish) error {
	ctx, cancel := client.contextFunc(context.Background())
	defer cancel()
	_, err := client.Collection.InsertOne(ctx, wish) // Use inserted ID?
	return err
}

func (client *WishClient) InsertMany(wishes []wish.Wish) error {
	ctx, cancel := client.contextFunc(context.Background())
	defer cancel()

	// ugh
	s := make([]interface{}, len(wishes))
	for i, w := range wishes {
		s[i] = w
	}
	_, err := client.Collection.InsertMany(ctx, s)
	return err
}

func (client *WishClient) Find(query bson.M) ([]wish.Wish, error) {
	ctx, cancel := client.contextFunc(context.Background())
	defer cancel()

	cursor, err := client.Collection.Find(ctx, query)
	if err != nil {
		return nil, err
	}

	res := make([]wish.Wish, 0, 10) // TODO: optimize this preallocation?

	for cursor.Next(context.TODO()) {
		var elem wish.Wish

		err := cursor.Decode(&elem)
		if err != nil {
			return nil, err
		}

		res = append(res, elem)
	}

	return res, nil
}
