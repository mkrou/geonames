package geonames

import (
	"fmt"
	"github.com/mkrou/geonames/models"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestReflection(t *testing.T) {
	Convey("Given a some language name", t, func() {
		val := "some language"

		Convey("When arguments filled", func() {
			var x *models.AlternateName
			err := fillArgument(func(m *models.AlternateName) error {
				x = m
				return nil
			}, func(v interface{}) error {
				x := v.(*models.AlternateName)
				x.Name = val
				return nil
			})

			Convey("The language name must be correct", func() {
				So(x.Name, ShouldEqual, val)
			})

			Convey("The error should be nill", func() {
				So(err, ShouldBeNil)
			})
		})
	})

	fmt.Println()
}
