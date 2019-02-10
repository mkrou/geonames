package models

type Hierarchy struct {
	Parent int    `csv:"parent"`
	Child  int    `csv:"child"`
	Type   string `csv:"type"`
}
