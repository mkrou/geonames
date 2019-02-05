package models

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestParseLanguage(t *testing.T) {
	Convey("Given a language with a starting values", t, func() {
		a := &Language{
			Iso639_1: "aa",
			Iso639_2: "aar",
			Iso639_3: "aar",
			Name:     "Afar",
		}

		Convey("When the language is parsed", func() {
			b, err := ParseLanguage(a.toArray())

			Convey("The error should be nill", func() {
				So(err, ShouldBeNil)
			})

			Convey("Objects must be equal", func() {
				So(*b, ShouldResemble, *a)
			})
		})
	})
}

func (l *Language) toArray() []string {
	return []string{
		l.Iso639_3,
		l.Iso639_2,
		l.Iso639_1,
		l.Name,
	}
}
