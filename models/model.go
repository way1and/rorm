package models

type Model struct {
	Name          string   // 模型名 (区分表)
	FieldRaws     []string // 字段列表
	FieldNames    []string // 字段映射名
	FieldTypes    []string
	KeyFieldRaws  []string // 主键 字段列表
	KeyFieldNames []string // 主键 映射列表
	Sync          bool     // 是否同步
	SyncFieldRaws []string // 同步字段列表
}

type Field struct {
	Name  string // 字段映射
	Raw   string // 字段名
	IsKey bool   // 是否是主键
	Sync  bool   // 是否开启同步 到 mysql 默认关闭
	Type  string // 类型名
}
