package interfaces

// 模型接口
type IModel interface {
	GetID() uint
	SetId(uint)
	TableName() string
}
