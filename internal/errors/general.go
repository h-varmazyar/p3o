package errors

import (
	"net/http"

	"github.com/h-varmazyar/p3o/pkg/errors"
)

// Link service errors :xyab
var (
	ErrUnexpected    = errors.NewWithHttp("general.unexpected", 1000, http.StatusBadRequest)
	ErrUnimplemented = errors.NewWithHttp("general.unimplemented", 1002, http.StatusBadRequest)
)
