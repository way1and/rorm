package rorm

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"github.com/way1and/rorm/models"
	"github.com/way1and/rorm/options"
)

var ctx = context.Background()

func Open(redisClient *redis.Client, options ...options.Options) (*DB, error) {
	if redisClient == nil {
		return nil, errors.New("rorm Open: val redisClient in parameters should not be <nil>")
	}

	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	return &DB{
		Client:  redisClient,
		Mapping: make(map[string]*models.Model),
	}, nil
}

// AppendModel 添加模型
func (db *DB) AppendModel(structs ...interface{}) *DB {
	for _, s := range structs {
		model := parseStructToModel(s)
		if model == nil {
			panic(errors.New("rorm AppendModel: Invalid struct no enable fields" + parseStructToMappingK(s)))
		}

		db.Mapping[parseStructToMappingK(s)] = model
	}
	return db
}
