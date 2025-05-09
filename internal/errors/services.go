package errors

import (
	"net/http"

	"github.com/h-varmazyar/p3o/pkg/errors"
)

// Link service errors :300ab
var (
	ErrLinkOwnerMismatch     = errors.NewWithHttp("services.link.owner_mismatch", 30000, http.StatusForbidden)
	ErrLinkActivatedBefore   = errors.NewWithHttp("services.link.activated_before", 30001, http.StatusBadRequest)
	ErrLinkDeactivatedBefore = errors.NewWithHttp("services.link.deactivated_before", 30002, http.StatusBadRequest)
	ErrLinkActivationBanned  = errors.NewWithHttp("services.link.activation_banned", 30003, http.StatusForbidden)
	ErrInvalidLink           = errors.NewWithHttp("services.link.invalid_link", 30004, http.StatusBadRequest)
	ErrKeyGenerationFailed   = errors.NewWithHttp("services.link.key_generation_failed", 30005, http.StatusInternalServerError)
	ErrLinkVisitMismatch     = errors.NewWithHttp("services.link.link_visit_mismatch", 30006, http.StatusBadRequest)
)

// user service errors :301ab
var (
	ErrWrongPassword           = errors.NewWithHttp("services.user.wrong_password", 30100, http.StatusBadRequest)
	ErrLoginFailed             = errors.NewWithHttp("services.user.login_failed", 30101, http.StatusBadRequest)
	ErrInvalidUsernamePassword = errors.NewWithHttp("services.user.invalid_username", 30102, http.StatusBadRequest)
	ErrPasswordHashingFailed   = errors.NewWithHttp("services.user.password_hashing_failed", 30103, http.StatusInternalServerError)
	ErrMobileMismatch          = errors.NewWithHttp("services.user.mobile_mismatch", 30104, http.StatusBadRequest)
	ErrWrongVerificationCode   = errors.NewWithHttp("services.user.wrong_verificationCode", 30105, http.StatusBadRequest)
	ErrUserMobileAvailable     = errors.NewWithHttp("services.user.user_mobile_available", 30106, http.StatusBadRequest)
)
