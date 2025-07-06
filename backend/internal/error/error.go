package error

type HTTPError struct {
	HTTPStatus int
	Msg        string
}

func New(status int, err *error) error {
	return &HTTPError{
		HTTPStatus: status,
		Msg:        (*err).Error(),
	}
}

func (err *HTTPError) Error() string {
	return err.Msg
}
