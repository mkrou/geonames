package models

import (
	. "github.com/smartystreets/goconvey/convey"
	"strconv"
	"testing"
	"time"
)

func TestParseAltername(t *testing.T) {
	Convey("Given some altername with a starting values", t, func() {
		from, _ := time.Parse("2006-01-02", "2018-02-02")
		to, _ := time.Parse("2006-01-02", "2018-03-03")
		a := &Altername{
			Id:           1,
			GeonameId:    2,
			IsoLanguage:  "a",
			Name:         "b",
			IsPreferred:  true,
			IsShort:      false,
			IsColloquial: true,
			IsHistoric:   false,
			From:         from,
			To:           to,
		}

		Convey("When the altername is parsed", func() {
			b, err := ParseAltername(a.toArray())

			Convey("The error should be nill", func() {
				So(err, ShouldBeNil)
			})

			Convey("Objects must be equal", func() {
				So(*b, ShouldResemble, *a)
			})
		})
	})
}

func (a *Altername) toArray() []string {
	return []string{
		strconv.Itoa(a.Id),
		strconv.Itoa(a.GeonameId),
		a.IsoLanguage,
		a.Name,
		boolToStr(a.IsPreferred),
		boolToStr(a.IsShort),
		boolToStr(a.IsColloquial),
		boolToStr(a.IsHistoric),
		a.From.Format("2006-02-01"),
		a.To.Format("2006-02-01"),
	}
}

func boolToStr(b bool) string {
	if b {
		return "1"
	} else {
		return ""
	}
}
