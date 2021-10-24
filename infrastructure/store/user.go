package store

import (
	"context"

	"github.com/aryanugroho/userapp-x-koki/common/log"
	"github.com/aryanugroho/userapp-x-koki/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	userCollection = "user"
)

type UserMongoStore struct {
	db   *mongo.Client
	user *mongo.Collection
}

func (s *UserMongoStore) Add(ctx context.Context, model *domain.User) (*domain.User, error) {
	result, err := s.user.InsertOne(ctx, model)
	if err != nil {
		log.Error(ctx, "failed to Add", err)
		return nil, err
	}
	objectID := result.InsertedID.(primitive.ObjectID)
	insertedID := objectID.Hex()
	model.ID = insertedID
	return model, nil
}
func (s *UserMongoStore) Update(ctx context.Context, model *domain.User) (*domain.User, error) {
	id, err := primitive.ObjectIDFromHex(model.ID)
	if err != nil {
		log.Error(ctx, "failed to Update", err)
		return nil, err
	}

	_, err = s.user.UpdateOne(ctx, withParam("_id", id), withParam("$set", model))
	if err != nil {
		log.Error(ctx, "failed to Update", err)
		return nil, err
	}

	return model, nil
}
func (s *UserMongoStore) Get(ctx context.Context, id string) (*domain.User, error) {
	var user domain.User
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Error(ctx, "failed to Get", err)
		return nil, err
	}
	res := s.user.FindOne(ctx, withParam("_id", oid))
	err = res.Decode(&user)
	if err != nil {
		log.Error(ctx, "failed to Get", err)
		return nil, err
	}
	return &user, nil
}

func (s *UserMongoStore) GetAll(ctx context.Context, params map[string]interface{}, pageNum, size int) ([]*domain.User, error) {
	var users []*domain.User
	param, err := mapToParam(params)
	if err != nil {
		log.Error(ctx, "failed to GetAll", err)
		return nil, err
	}

	cur, err := s.user.Find(ctx, param, withPagination(pageNum, size))
	if err != nil {
		log.Error(ctx, "failed to GetAll", err)
		return nil, err
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		user := domain.User{}
		err = cur.Decode(&user)
		if err != nil {
			log.Error(ctx, "failed to GetAll", err)
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}

func (s *UserMongoStore) CountAll(ctx context.Context, params map[string]interface{}) (int64, error) {
	count, err := s.user.CountDocuments(ctx, nil)
	if err != nil {
		log.Error(ctx, "failed to CountAll", err)
		return 0, err
	}

	return count, nil
}

func (s *UserMongoStore) Delete(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Error(ctx, "failed to Delete", err)
		return err
	}

	_, err = s.user.DeleteOne(ctx, withParam("_id", oid))
	if err != nil {
		log.Error(ctx, "failed to Delete", err)
		return err
	}

	return nil
}

func (s *UserMongoStore) FindByPhoneNumber(ctx context.Context, phoneNumber string) (*domain.User, error) {
	var user domain.User
	res := s.user.FindOne(ctx, withParam("phone_number", phoneNumber))
	err := res.Decode(&user)
	if err != nil {
		log.Error(ctx, "failed to FindByPhoneNumber", err)
		return nil, err
	}
	return &user, nil
}
