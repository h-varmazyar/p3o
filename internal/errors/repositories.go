package errors

import (
	"github.com/h-varmazyar/p3o/pkg/errors"
	"net/http"
)

// Link repository errors :200ab
var (
	//cache errors start from 20000 to 20049
	ErrCacheInsertFailed = errors.NewWithCode("repositories.link.cache_insert_failed", 20000)
	ErrCacheFetchFailed  = errors.NewWithCode("repositories.link.cache_insert_failed", 20001)

	//postgres errors start from 20050 to 20099
	ErrFailedToCreateLink    = errors.NewWithCode("repositories.link.create_link_failed", 20050)
	ErrLinkNotFound          = errors.NewWithHttp("repositories.link.link_not_found", 20051, http.StatusNotFound)
	ErrIncreaseVisitFailed   = errors.NewWithCode("repositories.link.increase_visit_failed", 20052)
	ErrLinkCountFetchFailed  = errors.NewWithCode("repositories.link.link_count_fetch_failed", 20053)
	ErrVisitCountFetchFailed = errors.NewWithCode("repositories.link.visit_count_fetch_failed", 20054)
	ErrUserHasNoLinks        = errors.NewWithCode("repositories.link.user_has_no_links", 20055)
)

// user service errors :201ab
var (
	ErrUserNotFound = errors.NewWithHttp("repositories.auth.user_not_found", 20100, http.StatusNotFound)
)
