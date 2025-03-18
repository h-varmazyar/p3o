package errors

import (
	"net/http"

	"github.com/h-varmazyar/p3o/pkg/errors"
)

// Link service errors :xyab
var (
	ErrUnexpected = errors.NewWithHttp("general.unexpected", 1000, http.StatusBadRequest)
)
