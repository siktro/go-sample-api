package web

type ErrorResponse struct {
	Error string `json:"error"`
}

// Error is used to add web infomation to a request error.
type WebError struct {
	Err    error
	Status int
}

// NewWebError is used when a known error condition is encountered.
func NewWebError(err error, status int) error {
	return &WebError{Err: err, Status: status}
}

func (e *WebError) Error() string {
	return e.Err.Error()
}
