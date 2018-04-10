// Code generated by moq; DO NOT EDIT
// github.com/matryer/moq

package commontest

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"sync"
	"time"
)

var (
	lockRCHTTPClienterMockDo                      sync.RWMutex
	lockRCHTTPClienterMockGet                     sync.RWMutex
	lockRCHTTPClienterMockGetMaxRetries           sync.RWMutex
	lockRCHTTPClienterMockHead                    sync.RWMutex
	lockRCHTTPClienterMockPost                    sync.RWMutex
	lockRCHTTPClienterMockPostForm                sync.RWMutex
	lockRCHTTPClienterMockPut                     sync.RWMutex
	lockRCHTTPClienterMockSetAuthToken            sync.RWMutex
	lockRCHTTPClienterMockSetDownloadServiceToken sync.RWMutex
	lockRCHTTPClienterMockSetFlorenceToken        sync.RWMutex
	lockRCHTTPClienterMockSetMaxRetries           sync.RWMutex
	lockRCHTTPClienterMockSetTimeout              sync.RWMutex
)

// RCHTTPClienterMock is a mock implementation of RCHTTPClienter.
//
//     func TestSomethingThatUsesRCHTTPClienter(t *testing.T) {
//
//         // make and configure a mocked RCHTTPClienter
//         mockedRCHTTPClienter := &RCHTTPClienterMock{
//             DoFunc: func(ctx context.Context, req *http.Request) (*http.Response, error) {
// 	               panic("TODO: mock out the Do method")
//             },
//             GetFunc: func(ctx context.Context, url string) (*http.Response, error) {
// 	               panic("TODO: mock out the Get method")
//             },
//             GetMaxRetriesFunc: func() int {
// 	               panic("TODO: mock out the GetMaxRetries method")
//             },
//             HeadFunc: func(ctx context.Context, url string) (*http.Response, error) {
// 	               panic("TODO: mock out the Head method")
//             },
//             PostFunc: func(ctx context.Context, url string, contentType string, body io.Reader) (*http.Response, error) {
// 	               panic("TODO: mock out the Post method")
//             },
//             PostFormFunc: func(ctx context.Context, uri string, data url.Values) (*http.Response, error) {
// 	               panic("TODO: mock out the PostForm method")
//             },
//             PutFunc: func(ctx context.Context, url string, contentType string, body io.Reader) (*http.Response, error) {
// 	               panic("TODO: mock out the Put method")
//             },
//             SetAuthTokenFunc: func(token string)  {
// 	               panic("TODO: mock out the SetAuthToken method")
//             },
//             SetDownloadServiceTokenFunc: func(token string)  {
// 	               panic("TODO: mock out the SetDownloadServiceToken method")
//             },
//             SetFlorenceTokenFunc: func(token string)  {
// 	               panic("TODO: mock out the SetFlorenceToken method")
//             },
//             SetMaxRetriesFunc: func(in1 int)  {
// 	               panic("TODO: mock out the SetMaxRetries method")
//             },
//             SetTimeoutFunc: func(timeout time.Duration)  {
// 	               panic("TODO: mock out the SetTimeout method")
//             },
//         }
//
//         // TODO: use mockedRCHTTPClienter in code that requires RCHTTPClienter
//         //       and then make assertions.
//
//     }
type RCHTTPClienterMock struct {
	// DoFunc mocks the Do method.
	DoFunc func(ctx context.Context, req *http.Request) (*http.Response, error)

	// GetFunc mocks the Get method.
	GetFunc func(ctx context.Context, url string) (*http.Response, error)

	// GetMaxRetriesFunc mocks the GetMaxRetries method.
	GetMaxRetriesFunc func() int

	// HeadFunc mocks the Head method.
	HeadFunc func(ctx context.Context, url string) (*http.Response, error)

	// PostFunc mocks the Post method.
	PostFunc func(ctx context.Context, url string, contentType string, body io.Reader) (*http.Response, error)

	// PostFormFunc mocks the PostForm method.
	PostFormFunc func(ctx context.Context, uri string, data url.Values) (*http.Response, error)

	// PutFunc mocks the Put method.
	PutFunc func(ctx context.Context, url string, contentType string, body io.Reader) (*http.Response, error)

	// SetAuthTokenFunc mocks the SetAuthToken method.
	SetAuthTokenFunc func(token string)

	// SetDownloadServiceTokenFunc mocks the SetDownloadServiceToken method.
	SetDownloadServiceTokenFunc func(token string)

	// SetFlorenceTokenFunc mocks the SetFlorenceToken method.
	SetFlorenceTokenFunc func(token string)

	// SetMaxRetriesFunc mocks the SetMaxRetries method.
	SetMaxRetriesFunc func(in1 int)

	// SetTimeoutFunc mocks the SetTimeout method.
	SetTimeoutFunc func(timeout time.Duration)

	// calls tracks calls to the methods.
	calls struct {
		// Do holds details about calls to the Do method.
		Do []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Req is the req argument value.
			Req *http.Request
		}
		// Get holds details about calls to the Get method.
		Get []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// URL is the url argument value.
			URL string
		}
		// GetMaxRetries holds details about calls to the GetMaxRetries method.
		GetMaxRetries []struct {
		}
		// Head holds details about calls to the Head method.
		Head []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// URL is the url argument value.
			URL string
		}
		// Post holds details about calls to the Post method.
		Post []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// URL is the url argument value.
			URL string
			// ContentType is the contentType argument value.
			ContentType string
			// Body is the body argument value.
			Body io.Reader
		}
		// PostForm holds details about calls to the PostForm method.
		PostForm []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// URI is the uri argument value.
			URI string
			// Data is the data argument value.
			Data url.Values
		}
		// Put holds details about calls to the Put method.
		Put []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// URL is the url argument value.
			URL string
			// ContentType is the contentType argument value.
			ContentType string
			// Body is the body argument value.
			Body io.Reader
		}
		// SetAuthToken holds details about calls to the SetAuthToken method.
		SetAuthToken []struct {
			// Token is the token argument value.
			Token string
		}
		// SetDownloadServiceToken holds details about calls to the SetDownloadServiceToken method.
		SetDownloadServiceToken []struct {
			// Token is the token argument value.
			Token string
		}
		// SetFlorenceToken holds details about calls to the SetFlorenceToken method.
		SetFlorenceToken []struct {
			// Token is the token argument value.
			Token string
		}
		// SetMaxRetries holds details about calls to the SetMaxRetries method.
		SetMaxRetries []struct {
			// In1 is the in1 argument value.
			In1 int
		}
		// SetTimeout holds details about calls to the SetTimeout method.
		SetTimeout []struct {
			// Timeout is the timeout argument value.
			Timeout time.Duration
		}
	}
}

// Do calls DoFunc.
func (mock *RCHTTPClienterMock) Do(ctx context.Context, req *http.Request) (*http.Response, error) {
	if mock.DoFunc == nil {
		panic("moq: RCHTTPClienterMock.DoFunc is nil but RCHTTPClienter.Do was just called")
	}
	callInfo := struct {
		Ctx context.Context
		Req *http.Request
	}{
		Ctx: ctx,
		Req: req,
	}
	lockRCHTTPClienterMockDo.Lock()
	mock.calls.Do = append(mock.calls.Do, callInfo)
	lockRCHTTPClienterMockDo.Unlock()
	return mock.DoFunc(ctx, req)
}

// DoCalls gets all the calls that were made to Do.
// Check the length with:
//     len(mockedRCHTTPClienter.DoCalls())
func (mock *RCHTTPClienterMock) DoCalls() []struct {
	Ctx context.Context
	Req *http.Request
} {
	var calls []struct {
		Ctx context.Context
		Req *http.Request
	}
	lockRCHTTPClienterMockDo.RLock()
	calls = mock.calls.Do
	lockRCHTTPClienterMockDo.RUnlock()
	return calls
}

// Get calls GetFunc.
func (mock *RCHTTPClienterMock) Get(ctx context.Context, url string) (*http.Response, error) {
	if mock.GetFunc == nil {
		panic("moq: RCHTTPClienterMock.GetFunc is nil but RCHTTPClienter.Get was just called")
	}
	callInfo := struct {
		Ctx context.Context
		URL string
	}{
		Ctx: ctx,
		URL: url,
	}
	lockRCHTTPClienterMockGet.Lock()
	mock.calls.Get = append(mock.calls.Get, callInfo)
	lockRCHTTPClienterMockGet.Unlock()
	return mock.GetFunc(ctx, url)
}

// GetCalls gets all the calls that were made to Get.
// Check the length with:
//     len(mockedRCHTTPClienter.GetCalls())
func (mock *RCHTTPClienterMock) GetCalls() []struct {
	Ctx context.Context
	URL string
} {
	var calls []struct {
		Ctx context.Context
		URL string
	}
	lockRCHTTPClienterMockGet.RLock()
	calls = mock.calls.Get
	lockRCHTTPClienterMockGet.RUnlock()
	return calls
}

// GetMaxRetries calls GetMaxRetriesFunc.
func (mock *RCHTTPClienterMock) GetMaxRetries() int {
	if mock.GetMaxRetriesFunc == nil {
		panic("moq: RCHTTPClienterMock.GetMaxRetriesFunc is nil but RCHTTPClienter.GetMaxRetries was just called")
	}
	callInfo := struct {
	}{}
	lockRCHTTPClienterMockGetMaxRetries.Lock()
	mock.calls.GetMaxRetries = append(mock.calls.GetMaxRetries, callInfo)
	lockRCHTTPClienterMockGetMaxRetries.Unlock()
	return mock.GetMaxRetriesFunc()
}

// GetMaxRetriesCalls gets all the calls that were made to GetMaxRetries.
// Check the length with:
//     len(mockedRCHTTPClienter.GetMaxRetriesCalls())
func (mock *RCHTTPClienterMock) GetMaxRetriesCalls() []struct {
} {
	var calls []struct {
	}
	lockRCHTTPClienterMockGetMaxRetries.RLock()
	calls = mock.calls.GetMaxRetries
	lockRCHTTPClienterMockGetMaxRetries.RUnlock()
	return calls
}

// Head calls HeadFunc.
func (mock *RCHTTPClienterMock) Head(ctx context.Context, url string) (*http.Response, error) {
	if mock.HeadFunc == nil {
		panic("moq: RCHTTPClienterMock.HeadFunc is nil but RCHTTPClienter.Head was just called")
	}
	callInfo := struct {
		Ctx context.Context
		URL string
	}{
		Ctx: ctx,
		URL: url,
	}
	lockRCHTTPClienterMockHead.Lock()
	mock.calls.Head = append(mock.calls.Head, callInfo)
	lockRCHTTPClienterMockHead.Unlock()
	return mock.HeadFunc(ctx, url)
}

// HeadCalls gets all the calls that were made to Head.
// Check the length with:
//     len(mockedRCHTTPClienter.HeadCalls())
func (mock *RCHTTPClienterMock) HeadCalls() []struct {
	Ctx context.Context
	URL string
} {
	var calls []struct {
		Ctx context.Context
		URL string
	}
	lockRCHTTPClienterMockHead.RLock()
	calls = mock.calls.Head
	lockRCHTTPClienterMockHead.RUnlock()
	return calls
}

// Post calls PostFunc.
func (mock *RCHTTPClienterMock) Post(ctx context.Context, url string, contentType string, body io.Reader) (*http.Response, error) {
	if mock.PostFunc == nil {
		panic("moq: RCHTTPClienterMock.PostFunc is nil but RCHTTPClienter.Post was just called")
	}
	callInfo := struct {
		Ctx         context.Context
		URL         string
		ContentType string
		Body        io.Reader
	}{
		Ctx:         ctx,
		URL:         url,
		ContentType: contentType,
		Body:        body,
	}
	lockRCHTTPClienterMockPost.Lock()
	mock.calls.Post = append(mock.calls.Post, callInfo)
	lockRCHTTPClienterMockPost.Unlock()
	return mock.PostFunc(ctx, url, contentType, body)
}

// PostCalls gets all the calls that were made to Post.
// Check the length with:
//     len(mockedRCHTTPClienter.PostCalls())
func (mock *RCHTTPClienterMock) PostCalls() []struct {
	Ctx         context.Context
	URL         string
	ContentType string
	Body        io.Reader
} {
	var calls []struct {
		Ctx         context.Context
		URL         string
		ContentType string
		Body        io.Reader
	}
	lockRCHTTPClienterMockPost.RLock()
	calls = mock.calls.Post
	lockRCHTTPClienterMockPost.RUnlock()
	return calls
}

// PostForm calls PostFormFunc.
func (mock *RCHTTPClienterMock) PostForm(ctx context.Context, uri string, data url.Values) (*http.Response, error) {
	if mock.PostFormFunc == nil {
		panic("moq: RCHTTPClienterMock.PostFormFunc is nil but RCHTTPClienter.PostForm was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		URI  string
		Data url.Values
	}{
		Ctx:  ctx,
		URI:  uri,
		Data: data,
	}
	lockRCHTTPClienterMockPostForm.Lock()
	mock.calls.PostForm = append(mock.calls.PostForm, callInfo)
	lockRCHTTPClienterMockPostForm.Unlock()
	return mock.PostFormFunc(ctx, uri, data)
}

// PostFormCalls gets all the calls that were made to PostForm.
// Check the length with:
//     len(mockedRCHTTPClienter.PostFormCalls())
func (mock *RCHTTPClienterMock) PostFormCalls() []struct {
	Ctx  context.Context
	URI  string
	Data url.Values
} {
	var calls []struct {
		Ctx  context.Context
		URI  string
		Data url.Values
	}
	lockRCHTTPClienterMockPostForm.RLock()
	calls = mock.calls.PostForm
	lockRCHTTPClienterMockPostForm.RUnlock()
	return calls
}

// Put calls PutFunc.
func (mock *RCHTTPClienterMock) Put(ctx context.Context, url string, contentType string, body io.Reader) (*http.Response, error) {
	if mock.PutFunc == nil {
		panic("moq: RCHTTPClienterMock.PutFunc is nil but RCHTTPClienter.Put was just called")
	}
	callInfo := struct {
		Ctx         context.Context
		URL         string
		ContentType string
		Body        io.Reader
	}{
		Ctx:         ctx,
		URL:         url,
		ContentType: contentType,
		Body:        body,
	}
	lockRCHTTPClienterMockPut.Lock()
	mock.calls.Put = append(mock.calls.Put, callInfo)
	lockRCHTTPClienterMockPut.Unlock()
	return mock.PutFunc(ctx, url, contentType, body)
}

// PutCalls gets all the calls that were made to Put.
// Check the length with:
//     len(mockedRCHTTPClienter.PutCalls())
func (mock *RCHTTPClienterMock) PutCalls() []struct {
	Ctx         context.Context
	URL         string
	ContentType string
	Body        io.Reader
} {
	var calls []struct {
		Ctx         context.Context
		URL         string
		ContentType string
		Body        io.Reader
	}
	lockRCHTTPClienterMockPut.RLock()
	calls = mock.calls.Put
	lockRCHTTPClienterMockPut.RUnlock()
	return calls
}

// SetAuthToken calls SetAuthTokenFunc.
func (mock *RCHTTPClienterMock) SetAuthToken(token string) {
	if mock.SetAuthTokenFunc == nil {
		panic("moq: RCHTTPClienterMock.SetAuthTokenFunc is nil but RCHTTPClienter.SetAuthToken was just called")
	}
	callInfo := struct {
		Token string
	}{
		Token: token,
	}
	lockRCHTTPClienterMockSetAuthToken.Lock()
	mock.calls.SetAuthToken = append(mock.calls.SetAuthToken, callInfo)
	lockRCHTTPClienterMockSetAuthToken.Unlock()
	mock.SetAuthTokenFunc(token)
}

// SetAuthTokenCalls gets all the calls that were made to SetAuthToken.
// Check the length with:
//     len(mockedRCHTTPClienter.SetAuthTokenCalls())
func (mock *RCHTTPClienterMock) SetAuthTokenCalls() []struct {
	Token string
} {
	var calls []struct {
		Token string
	}
	lockRCHTTPClienterMockSetAuthToken.RLock()
	calls = mock.calls.SetAuthToken
	lockRCHTTPClienterMockSetAuthToken.RUnlock()
	return calls
}

// SetDownloadServiceToken calls SetDownloadServiceTokenFunc.
func (mock *RCHTTPClienterMock) SetDownloadServiceToken(token string) {
	if mock.SetDownloadServiceTokenFunc == nil {
		panic("moq: RCHTTPClienterMock.SetDownloadServiceTokenFunc is nil but RCHTTPClienter.SetDownloadServiceToken was just called")
	}
	callInfo := struct {
		Token string
	}{
		Token: token,
	}
	lockRCHTTPClienterMockSetDownloadServiceToken.Lock()
	mock.calls.SetDownloadServiceToken = append(mock.calls.SetDownloadServiceToken, callInfo)
	lockRCHTTPClienterMockSetDownloadServiceToken.Unlock()
	mock.SetDownloadServiceTokenFunc(token)
}

// SetDownloadServiceTokenCalls gets all the calls that were made to SetDownloadServiceToken.
// Check the length with:
//     len(mockedRCHTTPClienter.SetDownloadServiceTokenCalls())
func (mock *RCHTTPClienterMock) SetDownloadServiceTokenCalls() []struct {
	Token string
} {
	var calls []struct {
		Token string
	}
	lockRCHTTPClienterMockSetDownloadServiceToken.RLock()
	calls = mock.calls.SetDownloadServiceToken
	lockRCHTTPClienterMockSetDownloadServiceToken.RUnlock()
	return calls
}

// SetFlorenceToken calls SetFlorenceTokenFunc.
func (mock *RCHTTPClienterMock) SetFlorenceToken(token string) {
	if mock.SetFlorenceTokenFunc == nil {
		panic("moq: RCHTTPClienterMock.SetFlorenceTokenFunc is nil but RCHTTPClienter.SetFlorenceToken was just called")
	}
	callInfo := struct {
		Token string
	}{
		Token: token,
	}
	lockRCHTTPClienterMockSetFlorenceToken.Lock()
	mock.calls.SetFlorenceToken = append(mock.calls.SetFlorenceToken, callInfo)
	lockRCHTTPClienterMockSetFlorenceToken.Unlock()
	mock.SetFlorenceTokenFunc(token)
}

// SetFlorenceTokenCalls gets all the calls that were made to SetFlorenceToken.
// Check the length with:
//     len(mockedRCHTTPClienter.SetFlorenceTokenCalls())
func (mock *RCHTTPClienterMock) SetFlorenceTokenCalls() []struct {
	Token string
} {
	var calls []struct {
		Token string
	}
	lockRCHTTPClienterMockSetFlorenceToken.RLock()
	calls = mock.calls.SetFlorenceToken
	lockRCHTTPClienterMockSetFlorenceToken.RUnlock()
	return calls
}

// SetMaxRetries calls SetMaxRetriesFunc.
func (mock *RCHTTPClienterMock) SetMaxRetries(in1 int) {
	if mock.SetMaxRetriesFunc == nil {
		panic("moq: RCHTTPClienterMock.SetMaxRetriesFunc is nil but RCHTTPClienter.SetMaxRetries was just called")
	}
	callInfo := struct {
		In1 int
	}{
		In1: in1,
	}
	lockRCHTTPClienterMockSetMaxRetries.Lock()
	mock.calls.SetMaxRetries = append(mock.calls.SetMaxRetries, callInfo)
	lockRCHTTPClienterMockSetMaxRetries.Unlock()
	mock.SetMaxRetriesFunc(in1)
}

// SetMaxRetriesCalls gets all the calls that were made to SetMaxRetries.
// Check the length with:
//     len(mockedRCHTTPClienter.SetMaxRetriesCalls())
func (mock *RCHTTPClienterMock) SetMaxRetriesCalls() []struct {
	In1 int
} {
	var calls []struct {
		In1 int
	}
	lockRCHTTPClienterMockSetMaxRetries.RLock()
	calls = mock.calls.SetMaxRetries
	lockRCHTTPClienterMockSetMaxRetries.RUnlock()
	return calls
}

// SetTimeout calls SetTimeoutFunc.
func (mock *RCHTTPClienterMock) SetTimeout(timeout time.Duration) {
	if mock.SetTimeoutFunc == nil {
		panic("moq: RCHTTPClienterMock.SetTimeoutFunc is nil but RCHTTPClienter.SetTimeout was just called")
	}
	callInfo := struct {
		Timeout time.Duration
	}{
		Timeout: timeout,
	}
	lockRCHTTPClienterMockSetTimeout.Lock()
	mock.calls.SetTimeout = append(mock.calls.SetTimeout, callInfo)
	lockRCHTTPClienterMockSetTimeout.Unlock()
	mock.SetTimeoutFunc(timeout)
}

// SetTimeoutCalls gets all the calls that were made to SetTimeout.
// Check the length with:
//     len(mockedRCHTTPClienter.SetTimeoutCalls())
func (mock *RCHTTPClienterMock) SetTimeoutCalls() []struct {
	Timeout time.Duration
} {
	var calls []struct {
		Timeout time.Duration
	}
	lockRCHTTPClienterMockSetTimeout.RLock()
	calls = mock.calls.SetTimeout
	lockRCHTTPClienterMockSetTimeout.RUnlock()
	return calls
}
