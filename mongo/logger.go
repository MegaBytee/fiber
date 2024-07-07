package mongo

import (
	"fmt"
	"time"

	"github.com/gofiber/storage/mongodb/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Logger struct {
	Storage *mongodb.Storage
}

type Log struct {
	ObjectID   primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Key        string             `json:"key" bson:"key"`
	Value      []byte             `json:"value" bson:"value"`
	Expiration time.Time          `json:"exp,omitempty" bson:"exp,omitempty"`
}

func LoggerSettings(mongo_uri string) *Logger {
	storage := mongodb.New(mongodb.Config{
		ConnectionURI: mongo_uri,
		Database:      "mongo_logger",
		Collection:    "logs",
		Reset:         true,
	})
	return &Logger{
		Storage: storage,
	}

}

func (log *Logger) Set(key, value string) {
	expire := 32400 * time.Second
	err := log.Storage.Set(key, []byte(value), expire)
	if err != nil {
		fmt.Println("Logger-set-err:>", err.Error())
	}

}
