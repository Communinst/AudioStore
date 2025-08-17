package httpError

type HTTPError struct {
	HTTPStatus int
	Msg        string
}

func New(status int, msg string) error {
	return &HTTPError{
		HTTPStatus: status,
		Msg:        msg,
	}
}

func (err *HTTPError) Error() string {
	return err.Msg
}
