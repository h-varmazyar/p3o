package errors

func Cast(err error) *Error {
	switch e := err.(type) {
	case nil:
		return nil
	case *Error:
		return e
	}
	return &Error{
		Message:       err.Error(),
		HttpCode:      DefaultHttpCode,
		Code:          DefaultCode,
		originalError: err,
	}
}
