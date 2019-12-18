package wish

import "go.mongodb.org/mongo-driver/bson"

func OwnerQuery(owner string) bson.M {
	return bson.M{"owner": owner}
}
