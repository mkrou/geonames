package models

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestParseGeoname(t *testing.T) {
	Convey("Given some geoname with a starting values", t, func() {
		date, _ := time.Parse("2006-01-02", "2018-02-02")
		a := &Geoname{
			Id:                    1,
			Name:                  "a",
			AsciiName:             "b",
			AlternateNames:        []string{"c", "d"},
			Latitude:              2.1,
			Longitude:             2.2,
			Class:                 "e",
			Code:                  "f",
			CountryCode:           "g",
			AlternateCountryCodes: []string{"i", "j"},
			Admin1Code:            "k",
			Admin2Code:            "l",
			Admin3Code:            "m",
			Admin4Code:            "n",
			Population:            3,
			Elevation:             4,
			DigitalElevationModel: 5,
			Timezone:              "Europe/Andorra",
			ModificationDate:      date,
		}

		Convey("When the geoname is parsed", func() {
			b, err := ParseGeoname(a.toArray())

			Convey("The error should be nill", func() {
				So(err, ShouldBeNil)
			})

			Convey("Objects must be equal", func() {
				So(*b, ShouldResemble, *a)
			})
		})
	})
}

func (g *Geoname) toArray() []string {
	return []string{
		strconv.Itoa(g.Id),
		g.Name,
		g.AsciiName,
		strings.Join(g.AlternateNames, ","),
		fmt.Sprintf("%f", g.Latitude),
		fmt.Sprintf("%f", g.Longitude),
		g.Class,
		g.Code,
		g.CountryCode,
		strings.Join(g.AlternateCountryCodes, ","),
		g.Admin1Code,
		g.Admin2Code,
		g.Admin3Code,
		g.Admin4Code,
		strconv.Itoa(g.Population),
		strconv.Itoa(g.Elevation),
		strconv.Itoa(g.DigitalElevationModel),
		g.Timezone,
		g.ModificationDate.Format("2006-02-01"),
	}
}
