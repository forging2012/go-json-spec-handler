package jsc

import (
	"net/url"
	"testing"

	"github.com/derekdowling/go-json-spec-handler"
	"github.com/derekdowling/jsh-api"
	. "github.com/smartystreets/goconvey/convey"
)

const testURL = "https://httpbin.org"

func TestClientRequest(t *testing.T) {

	Convey("Client Tests", t, func() {

		Convey("->setPath()", func() {
			url := &url.URL{Host: "test"}

			Convey("should format properly", func() {
				setPath(url, "tests")
				So(url.String(), ShouldEqual, "//test/tests")
			})

			Convey("should respect an existing path", func() {
				url.Path = "admin"
				setPath(url, "test")
				So(url.String(), ShouldEqual, "//test/admin/test")
			})
		})

		Convey("->setIDPath()", func() {
			url := &url.URL{Host: "test"}

			Convey("should format properly an id url", func() {
				setIDPath(url, "tests", "1")
				So(url.String(), ShouldEqual, "//test/tests/1")
			})
		})

	})
}

func TestResponseParsing(t *testing.T) {

	Convey("Response Parsing Tests", t, func() {

		Convey("->ParseObject()", func() {

			obj, objErr := jsh.NewObject("123", "test", map[string]string{"test": "test"})
			So(objErr, ShouldBeNil)

			response, err := mockObjectResponse(obj)
			So(err, ShouldBeNil)

			Convey("should parse successfully", func() {
				doc, err := Document(response)

				So(err, ShouldBeNil)
				So(doc.HasData(), ShouldBeTrue)
				So(doc.First().ID, ShouldEqual, "123")
			})
		})

		Convey("->GetList()", func() {

			obj, objErr := jsh.NewObject("123", "test", map[string]string{"test": "test"})
			So(objErr, ShouldBeNil)

			list := jsh.List{obj, obj}

			response, err := mockListResponse(list)
			So(err, ShouldBeNil)

			Convey("should parse successfully", func() {
				doc, err := Document(response)

				So(err, ShouldBeNil)
				So(doc.HasData(), ShouldBeTrue)
				So(doc.First().ID, ShouldEqual, "123")
			})
		})
	})
}

// not a great for this, would much rather have it in test_util, but it causes an
// import cycle wit jsh-api
func testAPI() *jshapi.API {
	resource := jshapi.NewMockResource("test", 1, nil)
	api := jshapi.New("")
	api.Add(resource)

	return api
}
