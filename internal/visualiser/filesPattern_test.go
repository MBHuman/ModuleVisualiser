package visualiser

import (
	"regexp"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestNewFilePattern(t *testing.T) {
	convey.Convey("Given a pattern string, NewFilePattern should return a FilePattern object with a compiled regex", t, func() {
		pattern := ".*\\.txt"
		expectedRegex := regexp.MustCompile(pattern)

		filePattern, err := NewFilePattern(pattern)

		convey.So(err, convey.ShouldBeNil)
		convey.So(filePattern.Pattern, convey.ShouldEqual, pattern)
		convey.So(filePattern.Regex.String(), convey.ShouldEqual, expectedRegex.String())
	})

	convey.Convey("Given an invalid pattern string, NewFilePattern should return an error", t, func() {
		pattern := "[[["

		filePattern, err := NewFilePattern(pattern)

		convey.So(filePattern, convey.ShouldBeNil)
		convey.So(err, convey.ShouldNotBeNil)
	})
}
