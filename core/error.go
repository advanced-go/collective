package core

// ErrorHandler - error handler interface
type ErrorHandler interface {
	Handle(s *Status) *Status
}
