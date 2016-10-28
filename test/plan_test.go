package test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAddToPlan(t *testing.T) {
	Convey("Tests add correctly to plan", t, func() {
		p := &Plan{}
		p.Add(&Test{})
		So(len(p.Tests), ShouldEqual, 1)
		p.Add(&Test{})
		p.Add(&Test{})
		So(len(p.Tests), ShouldEqual, 3)
	})
}
