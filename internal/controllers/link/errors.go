package link

import (
	"github.com/h-varmazyar/p3o/pkg/errors"
	"net/http"
)

// error code format: 120ab
var (
	ErrInvalidLink         = errors.NewWithHttp("invalid_link", 12000, http.StatusBadRequest)
	ErrInvalidData         = errors.NewWithHttp("invalid_data", 12001, http.StatusBadRequest)
	ErrKeyGenerationFailed = errors.NewWithHttp("key_generation_failed", 12002, http.StatusInternalServerError)
)
