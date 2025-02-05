package core

import (
	"bytes"
	"errors"
	"fmt"
	//"github.com/behavioral-ai/core/core"
	//"github.com/behavioral-ai/core/iox"
	"io"
	"net/http"
	"reflect"
)

const (
	fileExistsError = "The system cannot find the file specified"
)

var (
	healthOK = []byte("{\n \"status\": \"up\"\n}")
	//healthLength = int64(len(healthOK))
)

func NewResponse(statusCode int, h http.Header, content any) (resp *http.Response, status *Status) {
	resp = &http.Response{StatusCode: statusCode, ContentLength: -1, Header: h, Body: io.NopCloser(bytes.NewReader([]byte{}))}
	if h == nil {
		resp.Header = make(http.Header)
	}
	if content == nil {
		return resp, NewStatus(statusCode)
	}
	switch ptr := (content).(type) {
	case []byte:
		if len(ptr) > 0 {
			resp.ContentLength = int64(len(ptr))
			resp.Body = io.NopCloser(bytes.NewReader(ptr))
		}
	case string:
		if ptr != "" {
			resp.ContentLength = int64(len(ptr))
			resp.Body = io.NopCloser(bytes.NewReader([]byte(ptr)))
		}
	case error:
		if ptr.Error() != "" {
			resp.ContentLength = int64(len(ptr.Error()))
			resp.Body = io.NopCloser(bytes.NewReader([]byte(ptr.Error())))
		}
	default:
		status = NewStatusError(StatusInvalidContent, errors.New(fmt.Sprintf("error: content type is invalid: %v", reflect.TypeOf(ptr))))
		return &http.Response{StatusCode: http.StatusInternalServerError, Header: SetHeader(nil, ContentType, ContentTypeText), ContentLength: int64(len(status.Err.Error())), Body: io.NopCloser(bytes.NewReader([]byte(status.Err.Error())))}, status
	}
	return resp, NewStatus(statusCode)
}

// NewResponseFromUri - read a Http response given a URL
func NewResponseFromUri(uri any) (*http.Response, *Status) {
	serverErr := &http.Response{StatusCode: http.StatusInternalServerError, Status: internalError, Header: make(http.Header)}
	if uri == nil {
		return serverErr, NewStatusError(StatusInvalidArgument, errors.New("error: URL is nil"))
	}
	//if u.Scheme != fileScheme {
	//	return serverErr, core.NewStatusError(core.StatusInvalidArgument, errors.New(fmt.Sprintf("error: Invalid URL scheme : %v", u.Scheme)))
	//}
	/*
		buf, status := iox.ReadFile(uri)
		if !status.OK() {
			if strings.Contains(status.Err.Error(), fileExistsError) {
				return &http.Response{StatusCode: http.StatusNotFound, Status: "Not Found", Header: make(http.Header)}, NewStatusError(StatusInvalidArgument, status.Err)
			}
			return serverErr, NewStatusError(StatusIOError, status.Err)
		}
		resp1, err2 := http.ReadResponse(bufio.NewReader(bytes.NewReader(buf)), nil)
		if err2 != nil {
			return serverErr, NewStatusError(StatusIOError, err2)
		}
	*/
	resp1 := &http.Response{}
	return resp1, StatusOK()

}

func NewHealthResponseOK() *http.Response {
	resp, _ := NewResponse(http.StatusOK, SetHeader(nil, ContentType, ContentTypeText), healthOK)
	return resp
}

func NewNotFoundResponse() *http.Response {
	resp, _ := NewResponse(http.StatusNotFound, SetHeader(nil, ContentType, ContentTypeText), StatusNotFound().String())
	return resp
}
