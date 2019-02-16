package models

type Language struct {
	Iso639_1 string `csv:"ISO 639-1"`
	Iso639_2 string `csv:"ISO 639-2"`
	Iso639_3 string `csv:"ISO 639-3" valid:"required"`
	Name     string `csv:"Language Name" valid:"required"`
}
