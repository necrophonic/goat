package spec

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestImportSpecFile(t *testing.T) {

	Convey("Simple spec file correctly imported", t, func() {
		s := Import("../testing/simple-1.test.yml")

		So(s.Name, ShouldEqual, "Simple test 1")
		So(s.Description, ShouldEqual, "A simple test layout")
		So(len(s.Scenarios), ShouldEqual, 2)

	})
}
