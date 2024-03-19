package interfaces

type IModel interface {
	GetID() uint
	SetId(uint)
	TableName() string
}
