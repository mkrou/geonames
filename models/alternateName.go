package models

/*
alternateNameId   : the id of this alternate name, int
geonameid         : geonameId referring to id in table 'geoname', int
isolanguage       : iso 639 language code 2- or 3-characters; 4-characters 'post' for postal codes and 'iata','icao' and faac for airport codes, fr_1793 for French Revolution names,  abbr for abbreviation, link to a website (mostly to wikipedia), wkdt for the wikidataid, varchar(7)
alternate name    : alternate name or name variant, varchar(400)
isPreferredName   : '1', if this alternate name is an official/preferred name
isShortName       : '1', if this is a short name like 'California' for 'State of California'
isColloquial      : '1', if this alternate name is a colloquial or slang term. Example: 'Big Apple' for 'New York'.
isHistoric        : '1', if this alternate name is historic and was used in the past. Example 'Bombay' for 'Mumbai'.
from		  : from period when the name was used
to		  : to period when the name was used
*/

type AlternateName struct {
	Id           int    `csv:"alternateNameId" valid:"required"`
	GeonameId    int    `csv:"geonameid" valid:"required"`
	IsoLanguage  string `csv:"isolanguage"`
	Name         string `csv:"alternate name" valid:"required"`
	IsPreferred  bool   `csv:"isPreferredName,omitempty"`
	IsShort      bool   `csv:"isShortName,omitempty"`
	IsColloquial bool   `csv:"isColloquial,omitempty"`
	IsHistoric   bool   `csv:"isHistoric,omitempty"`
	From         Time   `csv:"from"`
	To           Time   `csv:"to"`
}

func (a *AlternateName) IsAlpha2() bool {
	return len(a.IsoLanguage) == 2
}
func (a *AlternateName) IsAlpha3() bool {
	return len(a.IsoLanguage) == 3
}
