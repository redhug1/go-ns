package client

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"testing"

	"github.com/ONSdigital/go-ns/log"
	"github.com/gorilla/mux"
	. "github.com/smartystreets/goconvey/convey"
)

func TestUnitClient(t *testing.T) {
	portChan := make(chan int)
	go mockZebedeeServer(portChan)

	port := <-portChan

	cli := NewZebedeeClient(fmt.Sprintf("http://localhost:%d", port))

	Convey("test Get()", t, func() {
		Convey("test get sucessfully returns response from zebedee", func() {
			b, err := cli.Get("/data?uri=foo")
			So(err, ShouldBeNil)

			So(string(b), ShouldEqual, `{}`)
		})

		Convey("test error returned if requesting invalid zebedee url", func() {
			b, err := cli.Get("/invalid")
			So(err, ShouldNotBeNil)
			So(err, ShouldHaveSameTypeAs, ErrInvalidZebedeeResponse{})
			So(err.Error(), ShouldEqual, "invalid response from zebedee - should be 2.x.x or 3.x.x, got: 404, path: /invalid")
			So(b, ShouldBeNil)
		})
	})

	Convey("test GetLanding", t, func() {
		Convey("test getLanding sucessfully returns a landing model", func() {
			m, err := cli.GetDatasetLandingPage("/data?uri=labor")
			So(err, ShouldBeNil)
			So(m, ShouldNotBeEmpty)
			So(m.Type, ShouldEqual, "dataset_landing_page")
		})

		Convey("test error returned if requesting invalid zebedee url", func() {
			_, err := cli.GetDatasetLandingPage("/invalid")
			So(err, ShouldNotBeNil)
			So(err, ShouldHaveSameTypeAs, ErrInvalidZebedeeResponse{})
			So(err.Error(), ShouldEqual, "invalid response from zebedee - should be 2.x.x or 3.x.x, got: 404, path: /invalid")
		})
	})

	Convey("test get dataset details", t, func() {
		d, err := cli.GetDataset("12345")
		So(err, ShouldBeNil)
		So(d.URI, ShouldEqual, "www.google.com")
		So(d.SupplementaryFiles[0].Title, ShouldEqual, "helloworld")
	})

	Convey("test getFileSize returns human readable filesize", t, func() {
		fs, err := cli.GetFileSize("filesize")
		So(err, ShouldBeNil)
		So(fs.Size, ShouldEqual, "5242880")
	})

	Convey("test getPageTitle returns a correctly formatted page title", t, func() {
		t, err := cli.GetPageTitle("pageTitle")
		So(err, ShouldBeNil)
		So(t.Title, ShouldEqual, "baby-names")
		So(t.Edition, ShouldEqual, "2017")
	})
}

func mockZebedeeServer(port chan int) {
	r := mux.NewRouter()

	r.Path("/data").HandlerFunc(d)
	r.Path("/parents").HandlerFunc(parents)
	r.Path("/filesize").HandlerFunc(filesize)

	l, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Error(err, nil)
		os.Exit(2)
	}

	port <- l.Addr().(*net.TCPAddr).Port
	close(port)

	if err := http.Serve(l, r); err != nil {
		log.Error(err, nil)
		os.Exit(2)
	}
}

func d(w http.ResponseWriter, req *http.Request) {
	uri := req.URL.Query().Get("uri")

	switch uri {
	case "foo":
		w.Write([]byte(`{}`))
	case "labor":
		w.Write([]byte(`{"downloads":[{"title":"Latest","file":"/employmentandlabourmarket/peopleinwork/workplacedisputesandworkingconditions/datasets/labourdisputesbysectorlabd02/labd02jul2015_tcm77-408195.xls"}],"section":{"markdown":""},"relatedDatasets":[{"uri":"/employmentandlabourmarket/peopleinwork/workplacedisputesandworkingconditions/datasets/labourdisputeslabd01"},{"uri":"/employmentandlabourmarket/peopleinwork/workplacedisputesandworkingconditions/datasets/stoppagesofworklabd03"}],"relatedDocuments":[{"uri":"/employmentandlabourmarket/peopleinwork/employmentandemployeetypes/bulletins/uklabourmarket/2015-07-15"}],"relatedMethodology":[],"type":"dataset_landing_page","uri":"/employmentandlabourmarket/peopleinwork/workplacedisputesandworkingconditions/datasets/labourdisputesbysectorlabd02","description":{"title":"Labour disputes by sector: LABD02","summary":"Labour disputes by sector.","keywords":["strike"],"metaDescription":"Labour disputes by sector.","nationalStatistic":true,"contact":{"email":"richard.clegg@ons.gsi.gov.uk\n","name":"Richard Clegg\n","telephone":"+44 (0)1633 455400 \n"},"releaseDate":"2015-07-14T23:00:00.000Z","nextRelease":"12 August 2015","datasetId":"","unit":"","preUnit":"","source":""}}`))
	case "12345":
		w.Write([]byte(`{"type":"dataset","uri":"www.google.com","downloads":[{"file":"test.txt"}],"supplementaryFiles":[{"title":"helloworld","file":"helloworld.txt"}],"versions":[{"uri":"www.google.com"}]}`))
	case "pageTitle":
		w.Write([]byte(`{"title":"baby-names","edition":"2017"}`))
	}

}

func parents(w http.ResponseWriter, req *http.Request) {
	uri := req.URL.Query().Get("uri")

	switch uri {
	case "/employmentandlabourmarket/peopleinwork/workplacedisputesandworkingconditions/datasets/labourdisputesbysectorlabd02":
		w.Write([]byte(`[{"uri":"/","description":{"title":"Home"},"type":"home_page"},{"uri":"/employmentandlabourmarket","description":{"title":"Employment and labour market"},"type":"taxonomy_landing_page"},{"uri":"/employmentandlabourmarket/peopleinwork","description":{"title":"People in work"},"type":"taxonomy_landing_page"},{"uri":"/employmentandlabourmarket/peopleinwork/workplacedisputesandworkingconditions","description":{"title":"Workplace disputes and working conditions"},"type":"product_page"}]`))
	}
}

func filesize(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte(`{"fileSize":"5242880"}`))
}
