package rorm

import (
	types "github.com/spf13/cast"
	"github.com/way1and/rorm/models"
	"github.com/way1and/rorm/tags"
	"reflect"
	"strings"
	"unicode"
)

// 解析 传入结构体 成 rorm 映射值
func parseStructToMappingK(s interface{}) string {
	return reflect.TypeOf(s).Name()
}

// 解析 传入的结构体 为 rorm 模型
func parseStructToModel(s interface{}) *models.Model {
	t := reflect.TypeOf(s).Elem()
	nums := t.NumField()
	model := new(models.Model)
	model.Name = parseName(t.Name())

	for n := 0; n < nums; n++ {
		f := t.Field(n)
		tag := f.Tag.Get("rorm")
		// 禁止解析
		if tag == "-" {
			continue
		}

		raw := f.Name

		// 解析 field tag
		field := parseFieldTag(tag)
		if field.Name == "" {
			field.Name = parseName(raw)
		}

		// 设置到 model
		model.FieldTypes = append(model.FieldTypes, f.Type.String()) // 类型
		model.FieldRaws = append(model.FieldRaws, raw)               // 原始名
		model.FieldNames = append(model.FieldNames, field.Name)      // 映射名
		if field.Sync {                                              // 是否同步
			model.Sync = true
			model.SyncFieldRaws = append(model.SyncFieldRaws, raw)
		}

		if field.IsKey { // 是否是键
			model.KeyFieldRaws = append(model.KeyFieldRaws, raw)
			model.KeyFieldNames = append(model.KeyFieldNames, field.Name)
		}
	}
	// 设置默认 key
	if len(model.KeyFieldRaws) == 0 {
		model.KeyFieldRaws = append(model.KeyFieldRaws, "ID")
	}
	return model
}

// 将 变量名 转化 为 映射名
func parseName(name string) string {
	name = strings.Replace(name, "ID", "id", 1)
	s := name[0:1]

	rs := []rune(name)
	for i := 1; i < len(rs); i++ {
		if unicode.IsUpper(rs[i]) { // 如果是大写
			s += "_"
			s += string(rs[i])
			continue
		}
		s += string(rs[i])
	}
	return strings.ToLower(s)
}

// 解析 结构体一行 标签
func parseFieldTag(tag string) *models.Field {

	options := strings.Split(tag, ";")
	field := new(models.Field)

	for _, option := range options {
		kv := strings.Split(option, ":")

		// 设置
		if kv[0] == tags.TagOptionName {
			field.Name = kv[1]
		} else if kv[0] == tags.TagOptionKey {
			field.IsKey = true
		} else if kv[0] == tags.TagOptionSync {
			field.Sync = true
		}
	}

	return field
}

// 将结构体 解析为 映射键值
func parseStructToKV(s interface{}, m *models.Model) (string, map[string]any) {
	data := make(map[string]any)
	key := m.Name

	v := reflect.ValueOf(s).Elem()

	for i, raw := range m.FieldRaws {
		field := v.FieldByName(raw)

		// 跳过 0 值
		if field.IsZero() {
			continue
		}
		data[m.FieldNames[i]] = field.Interface()
	}

	// 取主键值 构造 主键
	for _, name := range m.KeyFieldNames {
		key += "_" + types.ToString(data[name])
	}

	return key, data
}

// 将结构体 解析为 映射键
func parseStructToK(s interface{}, m *models.Model) string {
	key := m.Name
	v := reflect.ValueOf(s).Elem()

	// 取主键值 构造 主键
	for _, raw := range m.KeyFieldRaws {
		key += "_" + types.ToString(v.FieldByName(raw).Interface())
	}
	return key
}
