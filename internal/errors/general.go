package errors

import (
	"net/http"

	"github.com/h-varmazyar/p3o/pkg/errors"
)

// Link service errors :100ab
var(
	ErrUnexpected = errors.NewWithHttp("general.unexpected", 10000, http.StatusBadRequest)
)