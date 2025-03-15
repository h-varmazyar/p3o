package errors

import (
	"net/http"

	"github.com/h-varmazyar/p3o/pkg/errors"
)

// Link service errors :300ab
var(
	ErrLinkOwnerMismatch = errors.NewWithHttp("services.link.owner_mismatch", 30000, http.StatusForbidden)
	ErrLinkActivatedBefore = errors.NewWithHttp("services.link.activated_before", 30001, http.StatusBadRequest)
	ErrLinkDeactivatedBefore = errors.NewWithHttp("services.link.deactivated_before", 30002, http.StatusBadRequest)
	ErrLinkActivationBanned = errors.NewWithHttp("services.link.activation_banned", 30003, http.StatusForbidden)
	ErrInvalidLink         = errors.NewWithHttp("services.link.invalid_link", 30004, http.StatusBadRequest)
	ErrKeyGenerationFailed = errors.NewWithHttp("services.link.key_generation_failed", 30005, http.StatusInternalServerError)
)