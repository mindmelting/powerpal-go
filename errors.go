package powerpalgo

import "fmt"

type PowerpalAuthenticationError struct{}
type PowerpalAuthorizationError struct{}
type PowerpalRequestError struct {
	StatusCode int
	ErrMessage string
}

func (r *PowerpalAuthenticationError) Error() string {
	return "authentication error"
}

func (r *PowerpalAuthorizationError) Error() string {
	return "authorization error"
}

func (r *PowerpalRequestError) Error() string {
	return fmt.Sprintf("[%[1]v]: %[2]v", r.StatusCode, r.ErrMessage)
}
