package repository

import "github.com/h-varmazyar/p3o/pkg/errors"

var (
	ErrUnimplemented = errors.NewWithCode("unimplemented", 21199)

	ErrCacheInsertFailed = errors.NewWithCode("cache_insert_failed", 21100)
	ErrCacheFetchFailed  = errors.NewWithCode("cache_insert_failed", 21101)

	ErrFailedToCreateLink  = errors.NewWithCode("create_link_failed", 21150)
	ErrLinkNotFound        = errors.NewWithCode("link_not_found", 21151)
	ErrIncreaseVisitFailed = errors.NewWithCode("increase_visit_failed", 21152)
)
