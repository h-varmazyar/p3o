package errors

import (
	"net/http"

	"github.com/h-varmazyar/p3o/pkg/errors"
)

// Link service errors :100ab
var(
	ErrInvalidData = errors.NewWithHttp("controllers.invalid_data", 10000, http.StatusBadRequest)
)