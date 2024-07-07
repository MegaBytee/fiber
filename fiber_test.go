package fiber

import (
	"fmt"
	"testing"

	"github.com/MegaBytee/fiber/service/demo"
	. "github.com/smartystreets/goconvey/convey"
)

func TestMain(t *testing.T) {
	Convey("Should be able to create new fiber microservice with mongo and run it", t, func() {

		x := NewFiber().SetService(*demo.Service)
		go x.Run()
		fmt.Println("fiber=", x)

		So(x.MongoDB.Connected(), ShouldEqual, true)

	})
}
