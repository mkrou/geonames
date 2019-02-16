package models

type AlternateNameModification struct {
	Id           int    `csv:"alternateNameId" valid:"required"`
	GeonameId    int    `csv:"geonameid" valid:"required"`
	IsoLanguage  string `csv:"isolanguage"`
	Name         string `csv:"alternate name" valid:"required"`
	IsPreferred  bool   `csv:"isPreferredName,omitempty"`
	IsShort      bool   `csv:"isShortName,omitempty"`
	IsColloquial bool   `csv:"isColloquial,omitempty"`
	IsHistoric   bool   `csv:"isHistoric,omitempty"`
}
