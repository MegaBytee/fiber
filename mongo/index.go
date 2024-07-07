package mongo

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func IndexUnique(key string, value bool) mongo.IndexModel {
	return mongo.IndexModel{
		Keys:    bson.M{key: 1}, // index in ascending order or -1 for descending order
		Options: options.Index().SetUnique(value),
	}
}

func IndexText(key string, weight int) mongo.IndexModel {
	return mongo.IndexModel{
		Keys:    bson.D{{Key: key, Value: "text"}},
		Options: options.Index().SetWeights(bson.D{{Key: key, Value: weight}}),
	}
}
