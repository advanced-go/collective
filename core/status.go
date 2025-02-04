package core

import "net/http"

type Status struct {
	Code int   `json:"code"`
	Err  error `json:"err"`
}

var okStatus = func() *Status {
	s := new(Status)
	s.Code = http.StatusOK
	return s
}()

var badRequestStatus = func() *Status {
	s := new(Status)
	s.Code = http.StatusBadRequest
	return s
}()

var notFoundStatus = func() *Status {
	s := new(Status)
	s.Code = http.StatusNotFound
	return s
}()

func StatusOK() *Status {
	return okStatus
}

func StatusBadRequest() *Status {
	return badRequestStatus
}

func StatusNotFound() *Status {
	return notFoundStatus
}

func NewStatus(code int) *Status {
	s := new(Status)
	s.Code = code
	return s
}

func NewStatusError(code int, err error) *Status {
	s := NewStatus(code)
	s.Code = code
	s.Err = err
	return s
}
