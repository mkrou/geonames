package models

type AdminCode5 struct {
	GeonameId  int    `csv:"geonameId" valid:"required"`
	AdminCode5 string `csv:"adm5code" valid:"required"`
}
