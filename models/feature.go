package models

type FeatureCode struct {
	Code        string `csv:"code" valid:"required"`
	Name        string `csv:"name" valid:"required"`
	Description string `csv:"description"`
}
