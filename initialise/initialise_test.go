package initialise

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

const (
	errStub = "folder not ready to initialise: "
)

func TestInitialiseFolder(t *testing.T) {
	Convey("Folder should not exist", t, func() {
		err := Folder("./invalid")
		So(err, ShouldNotBeNil)
		So(err.Error(), ShouldEqual, errStub+"target [./invalid] does not exist")
	})

	Convey("Folder is not empty", t, func() {
		err := Folder(".")
		So(err, ShouldNotBeNil)
		So(err.Error(), ShouldEqual, errStub+"target [.] is not empty")
	})

	Convey("Target is not a folder", t, func() {
		err := Folder("./initialise.go")
		So(err, ShouldNotBeNil)
		So(err.Error(), ShouldEqual, errStub+"target [./initialise.go] is not a folder")
	})
}
