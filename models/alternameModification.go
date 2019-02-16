package models

type AlternameModification struct {
	Id           int    `csv:"alternateNameId"`
	GeonameId    int    `csv:"geonameid"`
	IsoLanguage  string `csv:"isolanguage"`
	Name         string `csv:"alternate name"`
	IsPreferred  bool   `csv:"isPreferredName,omitempty"`
	IsShort      bool   `csv:"isShortName,omitempty"`
	IsColloquial bool   `csv:"isColloquial,omitempty"`
	IsHistoric   bool   `csv:"isHistoric,omitempty"`
}
