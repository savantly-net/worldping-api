package components

import (
	"io/ioutil"
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPhantomRender(t *testing.T) {

	Convey("Can render url", t, func() {
		tempDir, _ := ioutil.TempDir("", "img")
		renderer := &PhantomRenderer{ImagesDir: tempDir, PhantomDir: "../../_vendor/phantomjs/"}
		png, err := renderer.RenderToPng("http://www.google.com")
		So(err, ShouldBeNil)
		So(exists(png), ShouldEqual, true)

		//_, err = os.Stat(store.getFilePathForDashboard("hello"))
		//So(err, ShouldBeNil)
	})

}

func exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
