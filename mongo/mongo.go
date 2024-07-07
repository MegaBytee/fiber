package mongo

import (
	"context"
	"fmt"
	"net/url"
	"time"

	paginate "github.com/gobeam/mongo-go-pagination"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct {
	uri       string
	db        *mongo.Database
	connected bool
	logger    *Logger
}

func New(config ...Config) *Mongo {
	// Set default config
	cfg := configDefault(config...)

	// Create data source name
	var dsn string

	// Check if user supplied connection string
	if cfg.ConnectionURI != "" {
		dsn = cfg.ConnectionURI
	} else {
		dsn = "mongodb://"
		if cfg.Username != "" {
			dsn += url.QueryEscape(cfg.Username)
		}
		if cfg.Password != "" {
			dsn += ":" + cfg.Password
		}
		if cfg.Username != "" || cfg.Password != "" {
			dsn += "@"
		}
		dsn += fmt.Sprintf("%s:%d", url.QueryEscape(cfg.Host), cfg.Port)
	}

	// Set mongo options
	opt := options.Client().ApplyURI(dsn)

	// Create and connect the mongo client in one step
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, opt)
	if err != nil {
		panic(err)
	}

	// verify that the client can connect
	if err = client.Ping(ctx, nil); err != nil {
		panic(err)
	}
	/*
		defer func() {
			if err = client.Disconnect(ctx); err != nil {
				panic(err)
			}
		}()*/
	// Get collection from database
	db := client.Database(cfg.Database)
	mg := &Mongo{
		uri:       dsn,
		db:        db,
		connected: true,
		logger:    LoggerSettings(dsn),
	}
	return mg
}
func (m *Mongo) Connected() bool {
	return m.connected
}
func (m *Mongo) GetURI() string {
	return m.uri
}

func (m *Mongo) Collection(name string) *mongo.Collection {
	return m.db.Collection(name)
}

func (m *Mongo) CreateIndex(name string, indexModel mongo.IndexModel) Error {
	var err error
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if m.connected {
		_, err = m.Collection(name).Indexes().CreateOne(ctx, indexModel)
	}

	return m.handleErrors(name, INDEX_CREATE, err)

}

func (m *Mongo) Get(name string, filter primitive.D, opts *options.FindOneOptions) *mongo.SingleResult {
	return m.Collection(name).FindOne(context.Background(), filter, opts)
}

func (m *Mongo) CheckSaved(name string, filter primitive.D) bool {
	var a any

	err := m.Get(name, filter, nil).Decode(&a)
	if err != nil && err == mongo.ErrNoDocuments {
		return false
	}
	return true
}

func (m *Mongo) Save(name string, data any) Error {
	//fmt.Println("mongo:Save:>", m)
	_, err := m.Collection(name).InsertOne(context.Background(), data)
	return m.handleErrors(name, SAVE, err)
}

func (m *Mongo) Update(name string, filter primitive.D, data primitive.D) Error {
	_, err := m.Collection(name).UpdateOne(context.Background(), filter, data)
	return m.handleErrors(name, UPDATE, err)
}

func (m *Mongo) Delete(name string, filter primitive.D) Error {
	_, err := m.Collection(name).DeleteOne(context.Background(), filter)
	return m.handleErrors(name, DELETE, err)
}

func (m *Mongo) Incr(name, key, value string, filter primitive.D) Error {
	data := bson.D{{Key: "$inc", Value: bson.D{{Key: key, Value: StringToInt(value)}}}}
	return m.Update(name, filter, data)
}

func (m *Mongo) Paginate(name string, filter primitive.D, limit, page int64, x any) (paginate.PaginationData, Error) {
	c := m.Collection(name)
	pQuery := paginate.New(c).Context(context.Background()).Limit(limit).Page(page).Filter(filter)
	paginatedData, err := pQuery.Decode(x).Find()
	return paginatedData.Pagination, m.handleErrors(name, PAGINATE, err)

}
