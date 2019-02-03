package models

import (
	"fmt"
	"strconv"
	"time"
)

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

type Altername struct {
	Id           int
	GeonameId    int
	IsoLanguage  string
	Name         string
	IsPreferred  bool
	IsShort      bool
	IsColloquial bool
	IsHistoric   bool
	From         time.Time
	To           time.Time
}

func (a *Altername) IsAlpha2() bool {
	return len(a.IsoLanguage) == 2
}
func (a *Altername) IsAlpha3() bool {
	return len(a.IsoLanguage) == 3
}

func getBool(val string) bool {
	return val == "1"
}

func ParseAltername(parts []string) (*Altername, error) {
	if len(parts) != 10 {
		return nil, fmt.Errorf("Line contains wrong number of columns: %d", len(parts))
	}

	id, err := strconv.Atoi(parts[0])
	if err != nil {
		return nil, err
	}

	geoId, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, err
	}

	return &Altername{
		Id:           id,
		GeonameId:    geoId,
		IsoLanguage:  parts[2],
		Name:         parts[3],
		IsPreferred:  getBool(parts[4]),
		IsShort:      getBool(parts[5]),
		IsColloquial: getBool(parts[6]),
		IsHistoric:   getBool(parts[7]),
		From:         parseDate(parts[8]),
		To:           parseDate(parts[9]),
	}, nil
}

func parseDate(date string) time.Time {
	if date == "" {
		return time.Time{}
	}

	if res, err := time.Parse("2006-01-02", date); err == nil {
		return res
	}

	if res, err := time.Parse("02 January 2006", date); err == nil {
		return res
	}

	if res, err := time.Parse("2006", date); err == nil {
		return res
	}

	if res, err := time.Parse("200601", date); err == nil {
		return res
	}

	if res, err := time.Parse("20060102", date); err == nil {
		return res
	}

	if res, err := time.Parse("02-01-2006", date); err == nil {
		return res
	}

	fmt.Println(date)
	return time.Time{}
}
