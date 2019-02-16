package models

type Hierarchy struct {
	Parent int    `csv:"parent" valid:"required"`
	Child  int    `csv:"child" valid:"required"`
	Type   string `csv:"type"`
}
