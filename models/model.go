package models

import "time"

type Model struct {
	Name          string   // 模型名 (区分表)
	FieldRaws     []string // 字段列表
	FieldNames    []string // 字段映射名
	FieldTypes    []string
	KeyFieldRaws  []string      // 主键 字段列表
	KeyFieldNames []string      // 主键 映射列表
	Sync          bool          // 是否同步
	SyncFieldRaws []string      // 同步字段列表
	Expire        bool          // 是否过期
	ExpireAfter   time.Duration // 过期时间
}

func (model *Model) AddField(field *Field) {
	// 设置到 model
	model.FieldTypes = append(model.FieldTypes, field.Type) // 类型
	model.FieldRaws = append(model.FieldRaws, field.Raw)    // 原始名
	model.FieldNames = append(model.FieldNames, field.Name) // 映射名
	if field.Sync {                                         // 是否同步
		model.Sync = true
		model.SyncFieldRaws = append(model.SyncFieldRaws, field.Raw)
	}

	if field.IsKey { // 是否是键
		model.KeyFieldRaws = append(model.KeyFieldRaws, field.Raw)
		model.KeyFieldNames = append(model.KeyFieldNames, field.Name)
	}
}

func (model *Model) SetExpireAfter(expire time.Duration) {
	model.Expire = true
	model.ExpireAfter = expire
}

type Field struct {
	Name    string // 字段映射
	Raw     string // 字段名
	IsKey   bool   // 是否是主键
	IsField bool   // 是否映射
	Sync    bool   // 是否开启同步 到 mysql 默认关闭
	Type    string // 类型名
}
