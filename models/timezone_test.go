package models

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestParseTimeZone(t *testing.T) {
	Convey("Given a time zone with a starting values", t, func() {
		a := &TimeZone{
			Id:          "America/Anchorage",
			CountryCode: "US",
			gmtOffset:   -9.0,
			dstOffset:   -8.0,
			rawOffset:   -9.0,
		}

		Convey("When the language is parsed", func() {
			b, err := ParseTimeZone(a.toArray())

			Convey("The error should be nill", func() {
				So(err, ShouldBeNil)
			})

			Convey("Objects must be equal", func() {
				So(*b, ShouldResemble, *a)
			})
		})
	})
}

func (l *TimeZone) toArray() []string {
	return []string{
		l.CountryCode,
		l.Id,
		fmt.Sprintf("%f", l.gmtOffset),
		fmt.Sprintf("%f", l.dstOffset),
		fmt.Sprintf("%f", l.rawOffset),
	}
}
