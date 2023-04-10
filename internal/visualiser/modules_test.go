package visualiser

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestFileModules(t *testing.T) {
	convey.Convey("Given a FileModules struct", t, func() {
		fileModules := FileModules{
			File: "example.txt",
			Modules: []string{
				"module1",
				"module2",
			},
		}

		convey.Convey("The file name should match", func() {
			convey.So(fileModules.File, convey.ShouldEqual, "example.txt")
		})

		convey.Convey("The modules list should match", func() {
			convey.So(len(fileModules.Modules), convey.ShouldEqual, 2)
			convey.So(fileModules.Modules[0], convey.ShouldEqual, "module1")
			convey.So(fileModules.Modules[1], convey.ShouldEqual, "module2")
		})
	})
}
