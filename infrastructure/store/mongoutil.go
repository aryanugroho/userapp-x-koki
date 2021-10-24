package store

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func oid(id string) (primitive.ObjectID, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return primitive.ObjectID{}, nil
	}
	return oid, nil
}

func withParam(key string, value interface{}) (setQuery bson.D) {
	setQuery = bson.D{
		{
			Key:   key,
			Value: value,
		},
	}
	return
}

func withPagination(pageNumber, pageSize int) *options.FindOptions {
	skip := int64((pageNumber - 1) * pageSize)

	limit := int64(pageSize)

	return &options.FindOptions{
		Limit: &limit,
		Skip:  &skip,
	}
}

func mapToParam(filter map[string]interface{}) (bson.D, error) {
	if len(filter) == 0 {
		return bson.D{}, nil
	}

	var params bson.D
	for key, value := range filter {
		if key == "_id" {
			objectID, err := oid(value.(string))
			if err != nil {
				return bson.D{}, err
			}
			params = append(params, primitive.E{Key: key, Value: objectID})
		} else {
			params = append(params, primitive.E{Key: key, Value: value})
		}
	}
	return params, nil
}
