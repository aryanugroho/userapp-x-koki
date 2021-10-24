package store

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/aryanugroho/userapp-x-koki/common/log"
	"github.com/aryanugroho/userapp-x-koki/config"
	"github.com/aryanugroho/userapp-x-koki/infrastructure"
)

func NewStore(ctx context.Context, config *config.Config) (infrastructure.Store, error) {
	connString := fmt.Sprintf("mongodb://%s:%s@%s:%s", config.DB.UserName, config.DB.Password, config.DB.Host, config.DB.Port)
	db, err := mongo.Connect(ctx, options.Client().ApplyURI(connString))
	if err != nil {
		log.Error(ctx, "failed init db", err)
		return nil, err
	}
	err = db.Ping(ctx, nil)
	if err != nil {
		log.Error(ctx, "failed ping db", err)
		return nil, err
	}

	store := MongoStore{
		user: UserMongoStore{db, db.Database(config.DB.Name).Collection(userCollection)},
	}

	return &store, nil
}

type MongoStore struct {
	user UserMongoStore
}

func (s *MongoStore) User() infrastructure.UserStore {
	return &s.user
}
