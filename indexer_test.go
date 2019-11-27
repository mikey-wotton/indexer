package indexer

import (
	"bytes"
	"io/ioutil"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCreateIndex(t *testing.T) {
	Convey("Given a call to CreateIndex", t, func() {
		Convey("A successful call's output should match Golden File", func() {
			fileOutName := "./testdata/output/index.html"
			err := CreateIndex("testdata\\directory", fileOutName)
			So(err, ShouldBeNil)
			if !fileCompareHelper(t, fileOutName) {
				t.Error("file generated by successful call does not match Golden File")
			}
		})
		Convey("An non-existent directory should return an error", func() {
			err := CreateIndex("", "index.html")
			//todo find specific error
			So(err, ShouldBeError)
		})

	})
}

func fileCompareHelper(t *testing.T, fileOutName string) bool {
	goldenFileName := "testdata\\golden\\index.html"
	f1, err := ioutil.ReadFile(goldenFileName)
	if err != nil {
		t.Error(err)
	}
	f2, err := ioutil.ReadFile(fileOutName)
	if err != nil {
		t.Error(err)
	}
	return bytes.Equal(f1, f2)
}
