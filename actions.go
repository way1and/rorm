package rorm

import (
	"errors"
	"github.com/redis/go-redis/v9"
	types "github.com/spf13/cast"
	"github.com/way1and/rorm/client"
	"github.com/way1and/rorm/models"
	"reflect"
	"time"
)

type DB struct {
	Client  *redis.Client
	Mapping map[string]*models.Model // 主键
	Target  *models.Model
	SyncDB  bool // 同步数据库
}

func (db *DB) Model(model interface{}) *DB {
	target := db.Mapping[parseStructToMappingK(model)]
	if target == nil {
		panic(errors.New("未迁移的对象"))
	}
	db.Target = target
	return db
}

// Get 获取值
func (db *DB) Get(model interface{}) bool {
	m := db.Mapping[parseStructToMappingK(model)]
	key := parseStructToK(model, m)
	v := client.Get(db.Client, key)
	// 不存在 返回false
	if v == nil {
		return false
	}
	// 值存在
	elem := reflect.ValueOf(model).Elem()

	for i, name := range m.FieldNames {
		field := elem.FieldByName(m.FieldRaws[i])
		t := m.FieldTypes[i]

		// 根据类型设置值
		if t == "string" {
			field.SetString(types.ToString(v[name]))
		} else if t == "int" {
			field.SetInt(types.ToInt64(v[name]))
		} else if t == "bool" {
			field.SetBool(types.ToBool(v[name]))
		} else if t == "float" {
			field.SetFloat(types.ToFloat64(v[name]))
		}
	}
	return true
}

// Exist 是否存在 ( 只对比键 )
func (db *DB) Exist(model interface{}) bool {
	return client.Exist(db.Client, parseStructToMappingK(model))
}

// GetEX 获取过期时间
func (db *DB) GetEX(model interface{}) time.Duration {
	return client.GetEX(db.Client, parseStructToMappingK(model))
}

// Set 设置值
func (db *DB) Set(model interface{}) bool {
	m := db.Mapping[parseStructToMappingK(model)]

	k, v := parseStructToKV(model, m)
	res := client.Sets(db.Client, k, v)
	if m.Expire { // 设置过期时间
		client.SetEX(db.Client, k, m.ExpireAfter)
	}
	return res
}

// IncrBy 更改值
func (db *DB) IncrBy(model interface{}, field string, change int) bool {
	m := db.Mapping[parseStructToMappingK(model)]
	key := parseStructToK(model, m)
	client.IncrBy(db.Client, key, parseName(field), int64(change))
	return false
}

// Delete 获取值
func (db *DB) Delete(model interface{}) bool {
	m := db.Mapping[parseStructToMappingK(model)]
	return client.Del(db.Client, parseStructToK(model, m))
}
