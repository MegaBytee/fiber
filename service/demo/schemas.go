package demo

import (
	"github.com/MegaBytee/fiber/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Hello struct {
	Name  string
	Value string
}

func (x *Hello) FilterID() primitive.D {
	return bson.D{{Key: "name", Value: x.Name}}
}

var HELLO = mongo.NewSchema("hello").
	SetIndex(mongo.IndexUnique("name", true))
