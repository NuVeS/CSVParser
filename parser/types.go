package parser

type Operation int8

const (
	More Operation = iota
	Less
	Equal
	NotEqual
	MoreOrEqual
	LessOrEqual
	none
)

type Expression struct {
	ColumnName string
	Operation  Operation
	Value      string
}
