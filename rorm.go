package rorm

import (
	"github.com/redis/go-redis/v9"
	"github.com/way1and/rorm/models"
	"github.com/way1and/rorm/options"
)

func Open(redisClient *redis.Client, options ...options.Options) *DB {
	return &DB{
		Client:  redisClient,
		Mapping: make(map[string]*models.Model),
	}
}

// AppendModel 添加模型
func (db *DB) AppendModel(models ...interface{}) {
	for _, model := range models {
		db.Mapping[parseStructToMappingK(model)] = parseStructToModel(model)
	}
}
