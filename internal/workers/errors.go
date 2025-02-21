package workers

import "github.com/h-varmazyar/p3o/pkg/errors"

// error code format: 421ab
var (
	ErrVisitWorkerStartedBefore = errors.NewWithCode("visit_worker_start_before", 21301)
	ErrNilVisitChannel          = errors.NewWithCode("nil_visit_chan", 21302)
)
