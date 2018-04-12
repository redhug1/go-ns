package identity

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ONSdigital/go-ns/common"
	"github.com/ONSdigital/go-ns/common/commontest"
	"github.com/pkg/errors"
	. "github.com/smartystreets/goconvey/convey"
)

const (
	url                = "/whatever"
	florenceToken      = "roundabout"
	callerAuthToken    = "YourClaimToBeWhoYouAre"
	callerIdentifier   = "externalCaller"
	userIdentifier     = "fred@ons.gov.uk"
	zebedeeURL         = "http://localhost:8082"
	expectedZebedeeURL = zebedeeURL + "/identity"
)

func TestHandler_NoAuth(t *testing.T) {

	Convey("Given a request with no auth info", t, func() {

		req := httptest.NewRequest("GET", url, nil)
		httpClient := &commontest.RCHTTPClienterMock{}
		idClient := NewAPIClient(httpClient, zebedeeURL)

		Convey("When CheckRequest is called", func() {
			ctx, _, err := idClient.CheckRequest(req)
			Convey("Then the downstream HTTP handler should not be called and no auth returned", func() {
				So(len(httpClient.DoCalls()), ShouldEqual, 0)
				So(err, ShouldBeNil)
				So(common.IsUserPresent(ctx), ShouldBeFalse)
				So(common.IsPresent(ctx), ShouldBeFalse)
				So(len(httpClient.DoCalls()), ShouldEqual, 0)
			})
		})
	})
}

func TestHandler_IdentityServiceError(t *testing.T) {

	Convey("Given a request with a florence token, and a mock client that returns an error", t, func() {

		req := httptest.NewRequest("GET", url, nil)
		req.Header = map[string][]string{
			common.FlorenceHeaderKey: {florenceToken},
		}

		expectedError := errors.New("broken")
		httpClient := getClientReturningError(expectedError)
		idClient := NewAPIClient(httpClient, zebedeeURL)

		Convey("When CheckRequest is called", func() {

			ctx, status, err := idClient.CheckRequest(req)

			Convey("Then the identity service was called as expected", func() {
				So(len(httpClient.DoCalls()), ShouldEqual, 1)
				So(httpClient.DoCalls()[0].Req.URL.String(), ShouldEqual, expectedZebedeeURL)
			})

			Convey("Then the error and no context is returned", func() {
				So(err, ShouldEqual, expectedError)
				So(status, ShouldNotEqual, http.StatusOK)
				So(ctx, ShouldBeNil)
			})
		})
	})
}

func TestHandler_IdentityServiceErrorResponseCode(t *testing.T) {

	Convey("Given a request with a florence token, and mock client that returns a non-200 response", t, func() {

		req := httptest.NewRequest("GET", url, nil)
		req.Header = map[string][]string{
			common.FlorenceHeaderKey: {florenceToken},
		}

		httpClient := &commontest.RCHTTPClienterMock{
			SetAuthTokenFunc: func(string) {},
			DoFunc: func(ctx context.Context, req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: http.StatusNotFound,
				}, nil
			},
		}
		idClient := NewAPIClient(httpClient, zebedeeURL)

		Convey("When CheckRequest is called", func() {

			ctx, status, err := idClient.CheckRequest(req)

			Convey("Then the identity service is called as expected", func() {
				So(len(httpClient.DoCalls()), ShouldEqual, 1)
				So(httpClient.DoCalls()[0].Req.URL.String(), ShouldEqual, expectedZebedeeURL)
			})

			Convey("Then there is no error but the response code matches the identity service", func() {
				So(err, ShouldBeNil)
				So(status, ShouldEqual, http.StatusNotFound)
				So(ctx, ShouldBeNil)
			})
		})
	})
}

func TestHandler_florenceToken(t *testing.T) {

	Convey("Given a request with a florence token, and mock client that returns 200", t, func() {

		req := httptest.NewRequest("GET", url, nil)
		req.Header = map[string][]string{
			common.FlorenceHeaderKey: {florenceToken},
		}

		httpClient := getClientReturningIdentifier(userIdentifier)
		idClient := NewAPIClient(httpClient, zebedeeURL)

		Convey("When CheckRequest is called", func() {

			ctx, status, err := idClient.CheckRequest(req)

			Convey("Then the identity service is called as expected", func() {
				So(err, ShouldBeNil)
				So(status, ShouldEqual, http.StatusOK)
				So(len(httpClient.DoCalls()), ShouldEqual, 1)
				zebedeeReq := httpClient.DoCalls()[0].Req
				So(zebedeeReq.URL.String(), ShouldEqual, expectedZebedeeURL)
				So(zebedeeReq.Header[common.FlorenceHeaderKey][0], ShouldEqual, florenceToken)
			})

			Convey("Then the downstream HTTP handler returned no error and expected context", func() {
				So(common.Caller(ctx), ShouldEqual, userIdentifier)
				So(common.User(ctx), ShouldEqual, userIdentifier)
			})
		})
	})
}

func TestHandler_InvalidIdentityResponse(t *testing.T) {

	Convey("Given a request with a florence token, and mock client that returns invalid response JSON", t, func() {

		req := httptest.NewRequest("GET", url, nil)
		req.Header = map[string][]string{
			common.FlorenceHeaderKey: {florenceToken},
		}

		httpClient := &commontest.RCHTTPClienterMock{
			SetAuthTokenFunc: func(string) {},
			DoFunc: func(ctx context.Context, req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       ioutil.NopCloser(bytes.NewBufferString("{ invalid JSON")),
				}, nil
			},
		}
		idClient := NewAPIClient(httpClient, zebedeeURL)

		Convey("When CheckRequest is called", func() {

			ctx, _, err := idClient.CheckRequest(req)

			Convey("Then the identity service is called as expected", func() {
				So(len(httpClient.DoCalls()), ShouldEqual, 1)
				zebedeeReq := httpClient.DoCalls()[0].Req
				So(zebedeeReq.URL.String(), ShouldEqual, expectedZebedeeURL)
				So(zebedeeReq.Header[common.FlorenceHeaderKey][0], ShouldEqual, florenceToken)
			})

			Convey("Then the response is set as expected", func() {
				So(err, ShouldNotBeNil)
				So(ctx, ShouldBeNil)
			})
		})
	})
}

func TestHandler_authToken(t *testing.T) {

	Convey("Given a request with an auth token, and mock client that returns 200", t, func() {

		req := httptest.NewRequest("GET", url, nil)
		req.Header = map[string][]string{
			common.AuthHeaderKey: {callerAuthToken},
			common.UserHeaderKey: {userIdentifier},
		}

		httpClient := getClientReturningIdentifier(callerIdentifier)
		idClient := NewAPIClient(httpClient, zebedeeURL)

		Convey("When CheckRequest is called", func() {

			ctx, status, err := idClient.CheckRequest(req)
			So(err, ShouldBeNil)
			So(status, ShouldEqual, http.StatusOK)

			Convey("Then the identity service is called as expected", func() {
				So(len(httpClient.DoCalls()), ShouldEqual, 1)
				zebedeeReq := httpClient.DoCalls()[0].Req
				So(zebedeeReq.URL.String(), ShouldEqual, expectedZebedeeURL)
				So(len(zebedeeReq.Header[common.UserHeaderKey]), ShouldEqual, 0)
				So(len(zebedeeReq.Header[common.AuthHeaderKey]), ShouldEqual, 1)
				So(zebedeeReq.Header[common.AuthHeaderKey][0], ShouldEqual, callerAuthToken)
			})

			Convey("Then the downstream HTTP handler request has the expected context values", func() {
				So(len(httpClient.DoCalls()), ShouldEqual, 1)
				So(common.IsPresent(ctx), ShouldBeTrue)
				So(common.IsUserPresent(ctx), ShouldBeTrue)
				So(common.Caller(ctx), ShouldEqual, callerIdentifier)
				So(common.User(ctx), ShouldEqual, userIdentifier)
			})
		})
	})
}

func TestHandler_bothTokens(t *testing.T) {

	Convey("Given a request with both a florence token and service token", t, func() {

		req := httptest.NewRequest("GET", url, nil)
		req.Header = map[string][]string{
			common.FlorenceHeaderKey: {florenceToken},
			common.AuthHeaderKey:     {callerAuthToken},
		}

		httpClient := getClientReturningIdentifier(userIdentifier)
		idClient := NewAPIClient(httpClient, zebedeeURL)

		Convey("When CheckRequest is called", func() {

			ctx, status, err := idClient.CheckRequest(req)
			So(err, ShouldBeNil)
			So(status, ShouldEqual, http.StatusOK)

			Convey("Then the identity service is called as expected", func() {
				So(len(httpClient.DoCalls()), ShouldEqual, 1)
				zebedeeReq := httpClient.DoCalls()[0].Req
				So(zebedeeReq.URL.String(), ShouldEqual, expectedZebedeeURL)
				So(zebedeeReq.Header[common.FlorenceHeaderKey][0], ShouldEqual, florenceToken)
			})

			Convey("Then the context returns with expected values", func() {
				So(common.IsUserPresent(ctx), ShouldBeTrue)
				So(common.User(ctx), ShouldEqual, userIdentifier)
				So(common.Caller(ctx), ShouldEqual, userIdentifier)
			})
		})
	})
}

func getClientReturningIdentifier(id string) *commontest.RCHTTPClienterMock {
	return &commontest.RCHTTPClienterMock{
		SetAuthTokenFunc: func(string) {},
		DoFunc: func(ctx context.Context, req *http.Request) (*http.Response, error) {
			response := &common.IdentityResponse{Identifier: id}
			body, _ := json.Marshal(response)
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewBuffer(body)),
			}, nil
		},
	}
}
func getClientReturningError(err error) *commontest.RCHTTPClienterMock {
	return &commontest.RCHTTPClienterMock{
		SetAuthTokenFunc: func(string) {},
		DoFunc: func(ctx context.Context, req *http.Request) (*http.Response, error) {
			return nil, err
		},
	}
}